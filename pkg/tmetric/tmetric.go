package tmetric

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	tmetric "git.missionfocus.com/ours/code/libraries/go/tmetric/client"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/client/accounts"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/client/time_entries"
	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/xanzy/go-gitlab"
)

const AccountID = 105432

type taskPerformanceRecord struct {
	description string
	url         string
	pointsSpent float64
	weight      int
	score       float64
}

func (r taskPerformanceRecord) More(other taskPerformanceRecord) bool {
	more := false
	if r.score > other.score {
		more = true
	} else if r.score == other.score {
		if r.weight > other.weight {
			more = true
		} else if r.weight == other.weight {
			if (r.url != "") && (other.url == "") {
				more = true
			}
		}
	}
	return more
}

func GetReports(glClient *gitlab.Client, vaultClient vault.Vault, progress io.Writer, startDate string, endDate string, format string) error {
	// Set up tabwriter for formatting output
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, ' ', 0)

	// Init regexps to get issue and MR info from paths
	issueRe := regexp.MustCompile(`/(.*)/-/issues/([0-9]+)`)

	secret, err := vaultClient.KVUserGet("tmetric")
	if err != nil {
		return err
	}
	if secret == nil {
		return errors.New("could not retrieve TMetric token. You may need to set it with `mfc tmetric set-token`")
	}

	tok := secret.Data["data"].(map[string]interface{})["token"].(string)
	auth := httptransport.BearerToken(tok)
	params := accounts.NewAccountsGetAccountScopeParams().WithAccountID(AccountID)

	fmt.Fprintln(progress, "Fetching GL Projects...")

	git := gl.New(glClient)

	// Query for all projects
	simpleQuery := true
	opts := &gitlab.ListProjectsOptions{
		Simple: &simpleQuery,
	}
	projects, err := git.ListAllProjectsWithOptions(opts)

	// Create a map from the project path to the struct
	projMap := make(map[string]*gitlab.Project)

	for _, p := range projects {
		projMap[p.PathWithNamespace] = p
	}

	fmt.Fprintln(progress, "Fetching TMetric Members...")

	// Get TMetric account scope seems to be the best way to get the list of members
	resp, err := tmetric.Default.Accounts.AccountsGetAccountScope(params, auth)
	if err != nil {
		return err
	}

	scope := resp.Payload

	for _, m := range scope.Members {
		profileId := m.UserProfileID

		startDt, err := strfmt.ParseDateTime(startDate)
		if err != nil {
			return err
		}

		endDt := strfmt.DateTime(time.Now())
		if endDate != "" {
			endDt, err = strfmt.ParseDateTime(startDate)
			if err != nil {
				return err
			}
		}

		params := time_entries.NewTimeEntriesGetTimeEntriesParams().
			WithAccountID(AccountID).
			WithUserProfileID(profileId).
			WithTimeRangeStartTime(&startDt).
			WithTimeRangeEndTime(&endDt)

		resp, err := tmetric.Default.TimeEntries.TimeEntriesGetTimeEntries(params, auth)
		if err != nil {
			return err
		}

		timeEntries := resp.Payload

		record := make(map[string]taskPerformanceRecord)

		for _, e := range timeEntries {
			start, err := time.Parse(time.RFC3339, e.StartTime.String())
			if err != nil {
				return err
			}

			end, err := time.Parse(time.RFC3339, e.EndTime.String())
			if err != nil {
				return err
			}

			duration := end.Sub(start)

			points := duration.Hours() / 8

			if points < 0 {
				continue
			}

			ref := e.Details.Description
			desc := e.Details.Description
			url := ""
			weight := 0

			if task := e.Details.ProjectTask; task != nil {
				ref = task.RelativeIssueURL

				if _, ok := record[ref]; !ok {
					desc = task.Description
					url = task.IntegrationURL + task.RelativeIssueURL

					if matches := issueRe.FindStringSubmatch(task.RelativeIssueURL); len(matches) == 3 {
						projPath := matches[1]
						if proj := projMap[projPath]; proj != nil {
							projID := proj.ID

							var issue *gitlab.Issue = nil
							idNum, err := strconv.Atoi(matches[2])
							if err != nil {
								return err
							}
							issue, err = git.GetIssue(projID, idNum)
							if err != nil {
								fmt.Printf("Error retrieving issue: %s\n", err)
								continue
							}

							weight = issue.Weight
						}
					}
				}
			}

			if val, ok := record[ref]; ok {
				val.pointsSpent += points
				record[ref] = val
			} else {
				record[ref] = taskPerformanceRecord{
					description: desc,
					url:         url,
					pointsSpent: points,
					weight:      weight,
					score:       0.0,
				}
			}
		}

		totalSpent := 0.0
		totalWeight := 0
		totalScore := 0.0

		switch format {
		case "md":
			fmt.Printf("\n## %s\n\n", m.UserProfile.UserName)
		case "org":
			fmt.Printf("\n** %s\n\n", m.UserProfile.UserName)
		}
		fmt.Fprintln(w, "| Issue\t| Points Spent\t| Weight\t| Score \t|")
		fmt.Fprintln(w, "|-\t|-\t|-\t-|-\t|")

		entries := make([]taskPerformanceRecord, 0, len(record))
		for _, v := range record {
			score := 0.0
			if v.pointsSpent > 0 {
				score = (float64(v.weight) / v.pointsSpent)
			}
			v.score = score

			totalSpent += v.pointsSpent
			totalWeight += v.weight
			totalScore += score

			entries = append(entries, v)
		}

		sort.SliceStable(entries, func(i int, j int) bool { return entries[i].More(entries[j]) })

		for _, v := range entries {
			if v.url == "" {
				fmt.Fprintf(w, "| %s \t| %.2f \t| %d \t| %.2f \t|\n", v.description, v.pointsSpent, v.weight, v.score)
			} else {
				switch format {
				case "md":
					fmt.Fprintf(w, "| [%s](%s) \t| %.2f \t| %d \t| %.2f \t|\n", v.description, v.url, v.pointsSpent, v.weight, v.score)
				case "org":
					fmt.Fprintf(w, "| [[%s][%s]] \t| %.2f \t| %d \t| %.2f \t|\n", v.url, v.description, v.pointsSpent, v.weight, v.score)
				}
			}
		}

		fmt.Fprintf(w, "| Total\t| %.2f\t| %d\t | %.2f \t|\n", totalSpent, totalWeight, totalScore)
		w.Flush()
	}

	return nil
}

func GetPersonHoursSummary(vaultClient vault.Vault, progress io.Writer, person string) error {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, ' ', 0)

	secret, err := vaultClient.KVUserGet("tmetric")
	if err != nil {
		return err
	}
	if secret == nil {
		return errors.New("could not retrieve TMetric token. You may need to set it with `mfc tmetric set-token`")
	}

	tok := secret.Data["data"].(map[string]interface{})["token"].(string)
	auth := httptransport.BearerToken(tok)
	params := accounts.NewAccountsGetAccountScopeParams().WithAccountID(AccountID)

	fmt.Fprintln(progress, "Fetching TMetric Member(s)...")

	resp, err := tmetric.Default.Accounts.AccountsGetAccountScope(params, auth)
	if err != nil {
		return err
	}

	scope := resp.Payload
	foundUser := false

	for _, m := range scope.Members {
		emailPointerValue := *m.UserProfile.Email
		profileId := m.UserProfileID
		if person == "" ||  strings.ToLower(m.UserProfile.UserName) == strings.ToLower(person) || emailPointerValue == person + "@missionfocus.com"  {
			foundUser = true
		} else {foundUser = false}
		if foundUser {
			calRange := []string{"past day", "week", "month", "lifetime"}
			for _, t := range calRange {
				startDt := strfmt.DateTime(time.Now())
				endDt := strfmt.DateTime(time.Now())
				switch t {
				case "past day":
					startDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * -10))
					endDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59))
				case "week":
					startDt = strfmt.DateTime(time.Now().AddDate(0, 0, -5).Add(time.Hour * -10))
					endDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59))
				case "month":
					startDt = strfmt.DateTime(time.Now().AddDate(0, -1, 0).Add(time.Hour * -10))
					endDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59))
				case "lifetime":
					startDt = strfmt.DateTime(time.Now().AddDate(-100, -1, 0).Add(time.Hour * -10))
					endDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59))
				}
				params := time_entries.NewTimeEntriesGetTimeEntriesParams().
					WithAccountID(AccountID).
					WithUserProfileID(profileId).
					WithTimeRangeStartTime(&startDt).
					WithTimeRangeEndTime(&endDt)

				resp, err := tmetric.Default.TimeEntries.TimeEntriesGetTimeEntries(params, auth)
				if err != nil {
					return err
				}
				timeEntries := resp.Payload

				var totalWorkedHours time.Duration = 0
				for _, e := range timeEntries {
					start, err := time.Parse(time.RFC3339, e.StartTime.String())
					if err != nil {
						return err
					}
					end, err := time.Parse(time.RFC3339, e.EndTime.String())
					if err != nil {
						return err
					}
					duration := end.Sub(start)
					totalWorkedHours = totalWorkedHours + duration
				}
				fmt.Println(m.UserProfile.UserName + "'s time this", t, "is a total of ", totalWorkedHours)
			}
		}
	}
	return nil
}

func Scanner (vaultClient vault.Vault, progress io.Writer) error {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, ' ', 0)

	secret, err := vaultClient.KVUserGet("tmetric")
	if err != nil {
		return err
	}
	if secret == nil {
		return errors.New("could not retrieve TMetric token. You may need to set it with `mfc tmetric set-token`")
	}

	tok := secret.Data["data"].(map[string]interface{})["token"].(string)
	auth := httptransport.BearerToken(tok)
	params := accounts.NewAccountsGetAccountScopeParams().WithAccountID(AccountID)

	fmt.Fprintln(progress, "Fetching TMetric Members...")

	resp, err := tmetric.Default.Accounts.AccountsGetAccountScope(params, auth)
	if err != nil {
		return err
	}

	scope := resp.Payload

	for _, m := range scope.Members {
		profileId := m.UserProfileID
		if m.UserProfile.UserName == "Mission Focus" {
			continue
		}
		fmt.Printf("\n## %s\n\n", m.UserProfile.UserName)

		fmt.Print("Checking username... ")
		if !Contains(acceptedUserNames, m.UserProfile.UserName) {
			fmt.Println("Failed. Unable to find " + m.UserProfile.UserName + " in profile database (username was changed or does not exist).")
		} else {
			fmt.Println("Passed")
		}

		// startDt/startDt test expected be at 10:00 AM each day; making start time = 12:00 AM and end time = 11:59 PM.
		startDt := strfmt.DateTime(time.Now().AddDate(0, 0 , -1).Add(time.Hour * -10)) // Start time = 12:00 AM
		endDt := strfmt.DateTime(time.Now().AddDate(0, 0 , -1).Add(time.Hour * 13).Add(time.Minute * 59)) // End Time = 11:59 PM
		params := time_entries.NewTimeEntriesGetTimeEntriesParams().
			WithAccountID(AccountID).
			WithUserProfileID(profileId).
			WithTimeRangeStartTime(&startDt).
			WithTimeRangeEndTime(&endDt)

		resp, err := tmetric.Default.TimeEntries.TimeEntriesGetTimeEntries(params, auth)
		if err != nil {
			return err
		}

		timeEntries := resp.Payload

		fmt.Print("Checking project names... ")
		var totalWorkedHours time.Duration
		var requiredHours = time.Duration(8)*time.Hour
		projsPassed := true

		for _, e := range timeEntries {
			desc := e.Details.Description

			if !Contains(acceptedProjects, e.ProjectName) {
				projsPassed = false
				fmt.Println("Failed. The entry " + desc + " by " + m.UserProfile.UserName + " has the invalid project of " + e.ProjectName)
			}

			start, err := time.Parse(time.RFC3339, e.StartTime.String())
			if err != nil {
				return err
			}
			end, err := time.Parse(time.RFC3339, e.EndTime.String())
			if err != nil {
				return err
			}
			duration := end.Sub(start)
			totalWorkedHours = totalWorkedHours + duration
		}
		if projsPassed {
			fmt.Println("Passed")
		}
		fmt.Print("Checking total worked hours... ")
		if totalWorkedHours < 0 {
			fmt.Println("Critical Error! User is still logging hours.")
		} else if totalWorkedHours < requiredHours {
			fmt.Print(totalWorkedHours)
			if totalWorkedHours < time.Duration(7)*time.Hour + time.Duration(30)*time.Minute { // If an employee has less than 7.5 hrs
				fmt.Println("Failed. Total hours is less than 7 hours and 30 minutes.")
			} else if totalWorkedHours > time.Duration(8)*time.Hour + time.Duration(30)*time.Minute {
				fmt.Println("Failed. Total hours is MORE than 8 hours and 30 minutes.")
			} else {
				fmt.Println("Warning. Total hours are less than 8 hours.")
			}
		} else { fmt.Println("Passed")}
	}
	return nil
}

//acceptedUserNames this should probably moved to an external file.
var acceptedUserNames = []string {
	"Jacob Stover",
	"Collin Day" ,
	"Cam Cook",
	"Eric Capito" ,
	"Wei Zhu" ,
	"John Kroeker",
	"Matthew Smith",
	"Casey Sault",
	"Alexander Gronowski",
	"Matthew Harbour",
	"Arlo Parker",
	"David Busey",
	"Levi Paulk",
	"Abraham Moshekh",
	"Andrew Zaw",
}

var acceptedProjects = []string{"GDAC", "BD-EDM", "PTO", "Overhead"}

// Contains tells whether A [array] contains S [String].
func Contains(a []string, s string) bool {
	for _, n := range a {
		if s == n {
			return true
		}
	}
	return false
}

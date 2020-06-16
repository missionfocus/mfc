package tmetric

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
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
	skill       int
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
			if r.skill > other.skill {
				more = true
			} else if (r.url != "") && (other.url == "") {
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
	skillRe := regexp.MustCompile(`^skill::([0-9]+)`)

	secret, err := vaultClient.KVUserGet("tmetric")
	if secret == nil || err != nil {
		return fmt.Errorf("could not retrieve TMetric token. You may need to set it with `mfc tmetric set-token`: %w", err)
	}

	auth := httptransport.BearerToken(secret.Data["token"].(string))
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
			skill := 0

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

							for _, l := range issue.Labels {
								if matches := skillRe.FindStringSubmatch(l); len(matches) == 2 {
									s, err := strconv.Atoi(matches[1])
									if err != nil {
										return err
									}
									skill = s
								}
							}
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
					skill:       skill,
					score:       0.0,
				}
			}
		}

		totalSpent := 0.0
		totalWeight := 0
		totalSkill := 0
		totalScore := 0.0

		switch format {
		case "md":
			fmt.Printf("\n## %s\n\n", m.UserProfile.UserName)
		case "org":
			fmt.Printf("\n** %s\n\n", m.UserProfile.UserName)
		}
		fmt.Fprintln(w, "| Issue\t| Points Spent\t| Weight\t| Skill\t| Score \t|")
		fmt.Fprintln(w, "|-\t|-\t|-\t|-\t|-\t|")

		entries := make([]taskPerformanceRecord, 0, len(record))
		for _, v := range record {
			score := 0.0
			if v.pointsSpent > 0 {
				score = (float64(v.weight) / v.pointsSpent) * float64(v.skill)
			}
			v.score = score

			totalSpent += v.pointsSpent
			totalWeight += v.weight
			totalSkill += v.skill
			totalScore += score

			entries = append(entries, v)
		}

		sort.SliceStable(entries, func(i int, j int) bool { return entries[i].More(entries[j]) })

		for _, v := range entries {
			if v.url == "" {
				fmt.Fprintf(w, "| %s \t| %.2f \t| %d \t| %d \t| %.2f \t|\n", v.description, v.pointsSpent, v.weight, v.skill, v.score)
			} else {
				switch format {
				case "md":
					fmt.Fprintf(w, "| [%s](%s) \t| %.2f \t| %d \t| %d \t| %.2f \t|\n", v.description, v.url, v.pointsSpent, v.weight, v.skill, v.score)
				case "org":
					fmt.Fprintf(w, "| [[%s][%s]] \t| %.2f \t| %d \t| %d \t| %.2f \t|\n", v.url, v.description, v.pointsSpent, v.weight, v.skill, v.score)
				}
			}
		}

		fmt.Fprintf(w, "| Total\t| %.2f\t| %d\t| %d\t| %.2f \t|\n", totalSpent, totalWeight, totalSkill, totalScore)
		w.Flush()
	}

	return nil
}

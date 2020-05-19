package main

import (
	"fmt"
	"os"
	"sort"
	"time"
	"regexp"
	"strconv"
	"text/tabwriter"
	"github.com/spf13/cobra"
	"github.com/go-openapi/strfmt"
	"github.com/xanzy/go-gitlab"
	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	tmetric "git.missionfocus.com/ours/code/tools/tmetric-api/client"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/accounts"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/time_entries"
	httptransport "github.com/go-openapi/runtime/client"
)

var (
	format string
	startDate string
	endDate string
	tmetricCmd = &cobra.Command{
		Use:     "tmetric",
		Short:   "Interact with GitLab",
		Aliases: []string{"tm"},
		Run: func(cmd *cobra.Command, args []string) {
			getReports()
		},
	}
)

const (
	UTC = "2006-01-02T00:00:00.000Z"
	TMETRIC_ACCOUNT_ID = 105432
)

type taskPerformanceRecord struct {
	description string
	url string
	pointsSpent float64
	weight int
	skill int
	score float64
}

func (r taskPerformanceRecord) More(other taskPerformanceRecord) (bool) {
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

func init() {
	mfcCmd.AddCommand(tmetricCmd)
	tmetricCmd.Flags().StringVarP(&format, "format", "f", "md", "output format to use for performace records")
	tmetricCmd.Flags().StringVarP(&startDate, "start-date", "d", "", "start date from which to query time entries")
	tmetricCmd.Flags().StringVarP(&endDate, "end-date", "e", "", "end date from which to query time entries")
}

func getReports() {
	// Set up tabwriter for formatting output
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, ' ', 0)

	// Init regexps to get issue and MR info from paths
	issueRe := regexp.MustCompile(`/(.*)/-/issues/([0-9]+)`)
	skillRe := regexp.MustCompile(`^skill::([0-9]+)`)

	// TODO: Read TMETRIC_TOKEN from personal vault
	auth := httptransport.BearerToken(os.Getenv("TMETRIC_TOKEN"))
	params := accounts.NewAccountsGetAccountScopeParams().WithAccountID(TMETRIC_ACCOUNT_ID)

	silentPrint("Fetching GL Projects...\n")

	// Setup gitlab client
	glClient, err := getGitLabClient()
	check(err)

	git := gl.New(glClient)

	// Query for all projects
	simpleQuery := true;
	opts := &gitlab.ListProjectsOptions{
		Simple: &simpleQuery,
	}
	projects, err := git.ListAllProjectsWithOptions(opts)

	// Create a map from the project path to the struct
	projMap := make(map[string]*gitlab.Project)

	for _, p := range projects {
		projMap[p.PathWithNamespace] = p
	}

	silentPrint("Fetching TMetric Members...\n")

	// Get TMetric account scope seems to be the best way to get the list of members
	resp, err := tmetric.Default.Accounts.AccountsGetAccountScope(params,  auth)
	check(err)

	scope := resp.Payload

	for _, m := range scope.Members {
		profileId := m.UserProfileID

		startDt, err := strfmt.ParseDateTime(startDate)
		check(err)

		endDt := strfmt.DateTime(time.Now())
		if endDate != "" {
			endDt, err = strfmt.ParseDateTime(startDate)
			check(err)
		}

		params := time_entries.NewTimeEntriesGetTimeEntriesParams().
			WithAccountID(TMETRIC_ACCOUNT_ID).
			WithUserProfileID(profileId).
			WithTimeRangeStartTime(&startDt).
			WithTimeRangeEndTime(&endDt)

		resp, err := tmetric.Default.TimeEntries.TimeEntriesGetTimeEntries(params, auth)
		check(err)

		timeEntries := resp.Payload

		record := make(map[string]taskPerformanceRecord)

		for _, e := range timeEntries {
			start, err := time.Parse(time.RFC3339, e.StartTime.String())
			check(err)

			end, err := time.Parse(time.RFC3339, e.EndTime.String())
			check(err)

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
							check(err)
							issue, err = git.GetIssue(projID, idNum);
							if err  != nil {
								fmt.Printf("Error retrieving issue: %s", err)
								continue
							}

							weight = issue.Weight

							for _, l := range issue.Labels {
								if matches := skillRe.FindStringSubmatch(l); len(matches) == 2 {
									s, err := strconv.Atoi(matches[1])
									check(err)
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
					url: url,
					pointsSpent: points,
					weight: weight,
					skill: skill,
					score: 0.0,
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
}

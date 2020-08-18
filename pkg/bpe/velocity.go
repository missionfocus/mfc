package bpe

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"github.com/xanzy/go-gitlab"

	"github.com/emirpasic/gods/utils"
)

const gitlabBaseURL = "https://git.missionfocus.com"

type TrackVelocity struct {
	issueTitles      []string
	issueURLs        []string
	weights          []int
	parentEpicTitles []string
	parentEpicURLs   []string
}

func (data *TrackVelocity) AppendTrackVelocity(issueTitle, issueURL string, weight int, epicTitle, parentEpicURL string) {
	data.issueTitles = append(data.issueTitles, issueTitle)
	data.issueURLs = append(data.issueURLs, issueURL)
	data.weights = append(data.weights, weight)
	data.parentEpicTitles = append(data.parentEpicTitles, epicTitle)
	data.parentEpicURLs = append(data.parentEpicURLs, parentEpicURL)
}

func VelocityReport(glClient *gitlab.Client, milestone, iteration string) error {
	g := gl.New(glClient)
	m := make(map[string]*TrackVelocity)
	labels := make(map[string]int)
	epics := make(map[string]int)

	state := "closed"
	scope := "all"
	opts := &gitlab.ListIssuesOptions{
		State:     &state,
		Milestone: &milestone,
		Scope:     &scope,
	}
	issues, _ := g.GetIssuesWithOptions(opts)

	for _, issue := range issues {
		if issue.Weight != 0 {
			epicTitle := ""
			epicURL := ""
			if issue.Epic != nil {
				epicTitle = issue.Epic.Title
				epicURL = gitlabBaseURL + issue.Epic.URL

				totalWeightPerEpic := epics[epicTitle]
				epics[epicTitle] = totalWeightPerEpic + issue.Weight
			}
			if _, ok := m[issue.Assignee.Name]; ok {
				foundPersonVelocity := m[issue.Assignee.Name]
				foundPersonVelocity.AppendTrackVelocity(issue.Title, issue.WebURL, issue.Weight, epicTitle, epicURL)
				m[issue.Assignee.Name] = foundPersonVelocity
			} else { //[]string{issue.Epic.Title}, []string{gitlabBaseURL + issue.Epic.URL}}
				newPersonVelocity := &TrackVelocity{[]string{issue.Title}, []string{issue.WebURL}, []int{issue.Weight}, []string{epicTitle}, []string{epicURL}}
				m[issue.Assignee.Name] = newPersonVelocity
			}
			for _, label := range issue.Labels {
				if strings.Contains(label, "epic-") {
					totalWeightPerLabel := labels[label]
					labels[label] = totalWeightPerLabel + issue.Weight
				}
			}
		}
	}

	csvfile, err := os.Create("VelocityReport " + milestone + ".csv")
	if err != nil {
		return err
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	headers := []string{"Person", "Issue Title", "Issue URL", "Issue Weight", "Epic Title" , "Epic URL"}
	writer.Write(headers)

	sumTotalWeight := 0
	for key, value := range m {
		totalSingleMFMWeight := 0

		for count, _ := range value.issueTitles {
			totalSingleMFMWeight = totalSingleMFMWeight + value.weights[count]
			record := []string{key, value.issueTitles[count], value.issueURLs[count], utils.ToString(value.weights[count]), value.parentEpicTitles[count], value.parentEpicURLs[count]}
			writer.Write(record)
		}
		sumTotalWeight = sumTotalWeight + totalSingleMFMWeight
		totalLine := []string{"", "", "MFM Total Weight: ", utils.ToString(totalSingleMFMWeight), ""}
		writer.Write(totalLine)

		blankLine := []string{""}
		writer.Write(blankLine)
	}
	totalWeight := []string{"", "", "All MFM Total Weight: ", utils.ToString(sumTotalWeight)}
	writer.Write(totalWeight)

	blankLine := []string{""}
	writer.Write(blankLine)

	headers = []string{"", "", "Epic Label", "Weight Per Label"}
	writer.Write(headers)

	for key, value := range labels {
		record := []string{"", "", key, utils.ToString(value)}
		writer.Write(record)
	}

	blankLine = []string{""}
	writer.Write(blankLine)

	headers = []string{"", "", "Epic Title", "Weight Per Title"}
	writer.Write(headers)

	for key, value := range epics {
		record := []string{"", "",  key, utils.ToString(value)}
		writer.Write(record)
	}
	fmt.Println("Results printed to file VelocityReport " + milestone + ".csv")

	return nil
}

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

type TrackVelocity struct {
	issueTitles []string
	issueURLs   []string
	weights     []int
}

func (data *TrackVelocity) AppendTrackVelocity(issuetitle string, issueURL string, weight int) {
	data.issueTitles = append(data.issueTitles, issuetitle)
	data.weights = append(data.weights, weight)
	data.issueURLs = append(data.issueURLs, issueURL)
}

func VelocityReport(glClient *gitlab.Client, milestone, iteration string) error {
	g := gl.New(glClient)
	m := make(map[string]*TrackVelocity)
	labels := make(map[string]int)

	state := "opened"
	opts := &gitlab.ListIssuesOptions{
		State: &state,
		Milestone: &milestone,
	}
	issues, _ := g.GetIssuesWithOptions(opts)

	for _, issue := range issues {
		if issue.Weight != 0 {
			if _, ok := m[issue.Assignee.Name]; ok {
				foundPersonVelocity := m[issue.Assignee.Name]
				foundPersonVelocity.AppendTrackVelocity(issue.Title, issue.WebURL, issue.Weight)
				m[issue.Assignee.Name] = foundPersonVelocity
			} else {
				newPersonVelocity := &TrackVelocity{[]string{issue.Title}, []string{issue.WebURL}, []int{issue.Weight}}
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

	csvfile, err := os.Create("VelocityReport.csv")
	if err != nil {
		return err
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	headers := []string{"Person", "Issue Title", "Issue URL", "Weight"}
	writer.Write(headers)

	for key, value := range m {
		totalWeight := 0
		writer.Write([]string{key})

		for count, _ := range value.issueTitles {
			totalWeight = totalWeight + value.weights[count]
			record := []string{"", value.issueTitles[count], value.issueURLs[count], utils.ToString(value.weights[count])}
			writer.Write(record)
		}

		totalLine := []string{"", "", "Total Weight: ", utils.ToString(totalWeight), ""}
		writer.Write(totalLine)

		blankLine := []string{""}
		writer.Write(blankLine)
	}

	headers = []string{"", "", "Epic Label", "Total Weight"}
	writer.Write(headers)

	for key, value := range labels {
		record := []string{"", "", key, utils.ToString(value)}
		writer.Write(record)
	}
	fmt.Println("Results printed to file IssueReport.csv")

	return nil
}

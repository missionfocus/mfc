package bpe

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"github.com/emirpasic/gods/utils"
	"github.com/xanzy/go-gitlab"
)

const (
	gitlabBaseURL = "https://git.missionfocus.com"
)

var blankLine = []string{""}
var promptEpicCheck, promptIssueCheck = true, true
var checkEpicAfterwards, checkIssueAfterwards = false, false

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
	if len(issues) == 0 {
		return fmt.Errorf("no issues were found for the milestone: %s", milestone)
	}
	for _, issue := range issues {
		if issue.Weight != 0 {
			if issue.Labels != nil {
				var labelArray []string
				addLabels := true
				for _, label := range issue.Labels {
					if strings.Contains(label, "dta") {
						addLabels = false
						break
					} else if strings.Contains(label, "epic-") {
						labelArray = append(labelArray, label)
					}
				}
				if addLabels {
					for _, label := range labelArray {
						totalWeightPerLabel := labels[label]
						labels[label] = totalWeightPerLabel + issue.Weight
					}
				} else { continue }
			} else if promptIssueCheck {
				log.Println("[WARNING] No labels for issue: " + issue.Title)
				checkIssueAfterwards, _ = promptForAnswer("Would you like run an issue check request after?")
				promptEpicCheck = false
			}

			epicTitle := ""
			epicURL := ""
			if issue.Epic != nil {
				epicTitle = issue.Epic.Title
				epicURL = gitlabBaseURL + issue.Epic.URL
				totalWeightPerEpic := epics[epicTitle]
				epics[epicTitle] = totalWeightPerEpic + issue.Weight
			} else if promptEpicCheck {
				log.Println("[WARNING] No parent epic for: " + issue.Title)
				checkEpicAfterwards, _ = promptForAnswer("Would you like run an epic check for epics after?")
				promptEpicCheck = false
			}

			if issue.Assignee == nil {
				log.Println("[WARNING] No assignee for: \"" + issue.Title + "\" using Author instead.")
				if _, ok := m[issue.Author.Name]; ok {
					foundPersonVelocity := m[issue.Assignee.Name]
					foundPersonVelocity.AppendTrackVelocity(issue.Title, issue.WebURL, issue.Weight, epicTitle, epicURL)
					m[issue.Assignee.Name] = foundPersonVelocity
				} else {
					newPersonVelocity := &TrackVelocity{[]string{issue.Title}, []string{issue.WebURL}, []int{issue.Weight}, []string{epicTitle}, []string{epicURL}}
					m[issue.Author.Name] = newPersonVelocity
				}
			} else if _, ok := m[issue.Assignee.Name]; ok {
				foundPersonVelocity := m[issue.Assignee.Name]
				foundPersonVelocity.AppendTrackVelocity(issue.Title, issue.WebURL, issue.Weight, epicTitle, epicURL)
				m[issue.Assignee.Name] = foundPersonVelocity
			} else {
				newPersonVelocity := &TrackVelocity{[]string{issue.Title}, []string{issue.WebURL}, []int{issue.Weight}, []string{epicTitle}, []string{epicURL}}
				m[issue.Assignee.Name] = newPersonVelocity
			}
		}
	}
	csvFile, err := os.Create("VelocityReport " + milestone + ".csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	headers := []string{"Person", "Issue Title", "Issue URL", "Issue Weight", "Epic Title", "Epic URL"}
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
		writer.Write(blankLine)
	}
	totalWeight := []string{"", "", "All MFM Total Weight: ", utils.ToString(sumTotalWeight)}
	err = writer.Write(totalWeight)
	if err != nil {
		return fmt.Errorf("ERROR Writing MFM Weight: %v",totalWeight)
	}
	writer.Write(blankLine)

	headers = []string{"", "", "Epic Label", "Weight Per Label"}
	writer.Write(headers)
	for key, value := range labels {
		record := []string{"", "", key, utils.ToString(value)}
		err = writer.Write(record)
		if err != nil {
			return fmt.Errorf("ERROR Writing labels: %v", record)
		}
	}
	writer.Write(blankLine)

	headers = []string{"", "", "Epic Title", "Weight Per Title"}
	writer.Write(headers)
	for key, value := range epics {
		record := []string{"", "", key, utils.ToString(value)}
		err = writer.Write(record)
		if err != nil {
			return fmt.Errorf("ERROR Writing Epic: %v", record)
		}
	}
	fmt.Println("Results printed to file VelocityReport " + milestone + ".csv")

	if checkIssueAfterwards {
		CheckIssuesWithOptions(glClient, opts, 0, nil)
	}
	if checkEpicAfterwards {
		//CheckEpicsWithinGroup(glClient, "", "", "", "")
	}
	return nil
}

func promptForAnswer(question string) (bool, error){
	fmt.Println(question + " [y/n]? ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return false, err
	}
	if strings.ToLower(input) == "y" {
		return true, nil
	}
	return false, nil
}

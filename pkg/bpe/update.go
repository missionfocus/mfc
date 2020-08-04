package bpe

import (
	"fmt"
	"strings"
	"time"

	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"github.com/asaskevich/govalidator"
	"github.com/xanzy/go-gitlab"
)

const (
	glTimeFormat    = "2006-01-02"
)

//GetTimeParameters is used to alter the format [date] | [date] into a comparable format
func GetTimeParameters(str string) []time.Time {
	dates := make([]time.Time, 0)

	if len(str) == 0 {
		date := "1999-12-31"
		t, _ := time.Parse(glTimeFormat, date)
		dates = append(dates, t)

		currentTime := time.Now()
		currentTime.Format(glTimeFormat)
		dates = append(dates, currentTime)
	}

	splitDateStrings := strings.Split(str, "|")
	for _, d := range splitDateStrings {
		strToDate := strings.Replace(d, " ", "", -1)
		t, _ := time.Parse(glTimeFormat, strToDate)
		dates = append(dates, t)
	}
	return dates
}

type EpicIssuesStruct struct {
	epic   *gitlab.Epic
	issues []*gitlab.Issue
}

func GetLabelParameters(str string) []string {
	if len(str) == 0 {
		return nil
	}
	label := make([]string, 0)
	splitDateStrings := strings.Split(str, "|")

	for _, d := range splitDateStrings {
		strToDate := strings.Replace(d, " ", "", -1)
		label = append(label, strToDate)
	}

	return label
}

//UpdateEpicIssuesLabels will update all labels related - includes epic and children issues
func UpdateEpicIssuesLabels(glClient *gitlab.Client, location, label string) error {
	g := gl.New(glClient)
	epicIssues := make([]EpicIssuesStruct, 0)
	fmt.Println("Location", location)
	labels := GetLabelParameters(label)
	if labels[0] == labels[1] {
		fmt.Println("Please try again. Error same label requested")
		return nil
	}

	groups, err := g.ListAllGroups()
	if err != nil {
		return err
	}

	if location == "" {
		return nil
	}

	locationFound := false
	epicHasOldLabel, epicHasNewLabel := false, false

	fmt.Println("Finding location epic and issues. Please wait...")
	for _, group := range groups {
		if !strings.Contains(location, "/code") && strings.Contains(group.WebURL, "/code") { // Included for optimization.
			continue
		}
		groupEpics, _ := g.ListAllGroupEpics(group.ID)
		for _, epic := range groupEpics {
			// The imported API does not use WebURL for epics -- this will detect epic location.
			if strings.Contains(location, "/epics/") {
				if strings.Contains(location, group.WebURL) {
					splitURL := strings.Split(location, "epics/")
					if splitURL[1] == govalidator.ToString(epic.IID) {
						locationFound = true
						fmt.Println("Epic found:", epic.Title)
						epicIssues = append(epicIssues, EpicIssuesStruct{epic, g.GetEpicIssues(group.ID, epic.IID)})
					}
				}
			}

			issues := g.GetEpicIssues(group.ID, epic.IID)
			for _, i := range issues {
				if i.WebURL == location {
					fmt.Println("Issue found:", i.Title)
					epicIssues = append(epicIssues, EpicIssuesStruct{epic, g.GetEpicIssues(group.ID, epic.IID)})
					locationFound = true
				}
				if locationFound == true {
					break
				}
			}

			if locationFound {
				for ct, label := range epic.Labels {
					if label == labels[0] {
						epicHasOldLabel = true
						epic.Labels = append(epic.Labels[:ct], epic.Labels[ct+1:]...)
					}
					if label == labels[1] {
						epicHasNewLabel = true
					}
				}
				if !epicHasNewLabel {
					epic.Labels = append(epic.Labels, labels[1])
				}
				if epicHasOldLabel || !epicHasNewLabel {
					opt := &gitlab.UpdateEpicOptions{
						Labels: epic.Labels,
					}
					g.UpdateEpicWithOpts(group.ID, epic.IID, opt)
				}

				for _, ei := range epicIssues {
					for _, issue := range ei.issues {
						issueHasOldLabel, issueHasNewLabel := false, false

						for ct, label := range issue.Labels {
							if label == labels[0] && labels[0] != "" {
								issueHasOldLabel = true
								issue.Labels = append(issue.Labels[:ct], issue.Labels[ct+1:]...)
							}
							if label == labels[1] && labels[1] != "" {
								issueHasNewLabel = true
							}
						}
						if !issueHasNewLabel {
							issue.Labels = append(issue.Labels, labels[1])
						}
						if issueHasOldLabel || !issueHasNewLabel {
							opt := &gitlab.UpdateIssueOptions{
								Labels: &issue.Labels,
							}
							g.UpdateIssueWithOpts(issue.ProjectID, issue.IID, opt)
						}
					}
				}
				break
			}
		}
		if locationFound {
			break
		}
	}
	return nil
}

//UpdateAllLabels - This will inherit parent epic labels to sub epics and issues.
func UpdateAllLabels(glClient *gitlab.Client) error {
	g := gl.New(glClient)

	groups, err := g.ListAllGroups()
	if err != nil {
		return err
	}

	for _, group := range groups {
		fmt.Println("\nUpdating Epics under the group:", group.Name)
		groupEpics, _ := g.ListAllGroupEpics(group.ID)
		for _, epic := range groupEpics {
			if epic.GroupID != group.ID {
				continue
			}
			if epic.ParentID != 0 {
				continue
			}
			fmt.Println("Updating Epics and Issues for the Parent Epic:", epic.Title)
			UpdateChildEpicsAndIssues(glClient, group, epic)
		}
	}
	return nil
}

func UpdateChildEpicsAndIssues(glClient *gitlab.Client, group *gitlab.Group, epic *gitlab.Epic) error {
	g := gl.New(glClient)
	epicLabels := epic.Labels
	issues := g.GetEpicIssues(group.ID, epic.IID)

	for _, issue := range issues {
		for _, epicLabel := range epicLabels {
			if strings.Contains(epicLabel, "epic-") {
				issue.Labels = append(issue.Labels, epicLabel)
			}
		}
		opt := &gitlab.UpdateIssueOptions{
			Labels: &issue.Labels,
		}
		g.UpdateIssueWithOpts(issue.ProjectID, issue.IID, opt)
	}

	childEpics := g.GetEpicLinks(group.ID, epic.IID)
	if childEpics != nil {
		for _, childEpic := range childEpics {
			if childEpic.ParentID == epic.ID {
				for _, epicLabel := range epicLabels {
					if strings.Contains(epicLabel, "epic-") {
						childEpic.Labels = append(childEpic.Labels, epicLabel)
					}
				}
				opt := &gitlab.UpdateEpicOptions{
					Labels: childEpic.Labels,
				}

				g.UpdateEpicWithOpts(group.ID, childEpic.IID, opt)
				UpdateChildEpicsAndIssues(glClient, group, childEpic)
			}
		}
	}
	return nil
}

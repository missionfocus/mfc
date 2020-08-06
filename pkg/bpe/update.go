package bpe

import (
	"fmt"
	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"github.com/xanzy/go-gitlab"
	"log"
	"strings"
)

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
	var epicIssues []*gitlab.Issue
	oursGroupID := 125

	labels := GetLabelParameters(label)
	fmt.Println(labels[0])
	fmt.Println(labels[1])
	if labels[0] == labels[1] {
		log.Fatal("Please try again. Error same label requested")
	}

	epicHasOldLabel, epicHasNewLabel := false, false

	groupEpics, _ := g.ListAllGroupEpics(oursGroupID) //TODO find a better method for getting epic.
	for _, epic := range groupEpics {
		if epic.WebURL == location {
			fmt.Println("Epic found:", epic.Title)
			epicIssues = g.GetEpicIssues(epic.GroupID, epic.IID)

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
				g.UpdateEpicWithOpts(oursGroupID, epic.IID, opt)
			}
			break
		}
	}

	for _, issue := range epicIssues {
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
	issues := g.GetEpicIssues(group.ID, epic.IID)

	for _, issue := range issues {
		for _, epicLabel := range epic.Labels {
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
				for _, epicLabel := range epic.Labels {
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

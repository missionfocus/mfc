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

const (
	oursGroupID = 125
	codeGroupID = 145
)

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
	if labels[0] == labels[1] {
		log.Fatal("Error same label requested. Please try again.")
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
	groups, _ := g.ListSubGroups(codeGroupID)
	addCodeGroup := g.GetGroup(codeGroupID) // Code includes Epics.
	groups = append(groups, addCodeGroup)

	for _, group := range groups {
		fmt.Println("\n--------------------------\n", "Reviewing Group: ", group.Name, "\n--------------------------")
		groupEpics, _ := g.ListAllGroupEpics(group.ID)
		for _, epic := range groupEpics {
			if epic.GroupID != group.ID { //DO NOT DELETE.
				continue
			}
			if epic.ParentID != 0 {
				continue
			}
			fmt.Println("\nRoot Parent Epic: " + epic.Title)
			UpdateChildEpicsAndIssues(glClient, group.ID, epic)
		}
	}
	return nil
}

func UpdateChildEpicsAndIssues(glClient *gitlab.Client, groupID int, epic *gitlab.Epic) error {
	g := gl.New(glClient)
	issues := g.GetEpicIssues(groupID, epic.IID)

	for _, issue := range issues {
		updateIssue := false
		for _, epicLabel := range epic.Labels {
			if strings.Contains(epicLabel, "epic-") && !contains(issue.Labels, epicLabel) {
				issue.Labels = append(issue.Labels, epicLabel)
				updateIssue = true
			}
		}
		if updateIssue {
			log.Println("[Updated] Issue: " + issue.Title)
			opt := &gitlab.UpdateIssueOptions{
				AddLabels: &issue.Labels,
			}
			g.UpdateIssueWithOpts(issue.ProjectID, issue.IID, opt)
		}
	}

	childEpics := g.GetEpicLinks(groupID, epic.IID)
	if childEpics != nil {
		for _, childEpic := range childEpics {
			updateEpic := false
			if childEpic.ParentID == epic.ID {
				for _, epicLabel := range epic.Labels {
					if strings.Contains(epicLabel, "epic-") && !contains(childEpic.Labels, epicLabel) {
						childEpic.Labels = append(childEpic.Labels, epicLabel)
						updateEpic = true
					}
				}
				if updateEpic {
					log.Println("[Updated] Child-Epic: " + childEpic.Title)
					opt := &gitlab.UpdateEpicOptions{
						Labels: childEpic.Labels,
					}
					g.UpdateEpicWithOpts(childEpic.GroupID, childEpic.IID, opt)
				}
				fmt.Println("  - [Checking] Child-Epic: " + childEpic.Title)
				UpdateChildEpicsAndIssues(glClient, childEpic.GroupID, childEpic) //Recursion does this process for inherited children.
			}
		}
	}
	return nil
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

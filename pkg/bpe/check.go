package bpe

import (
	"encoding/csv"
	"fmt"
	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
	"strings"
)

type EpicReport struct {
	epic   *gitlab.Epic
	reason string
}

type IssueReport struct {
	issue  *gitlab.Issue
	reason string
}

func CheckIssuesWithinProject(glClient *gitlab.Client, location string, cd string, ud string, state string) error {
	g := gl.New(glClient)
	issuesInReport := make([]IssueReport, 0)
	creationDates := GetTimeParameters(cd)
	updatedDates := GetTimeParameters(ud)

	if state == "" {
		state = "all"
	}

	var Issues []*gitlab.Issue
	if location == "" {
		opt := &gitlab.ListIssuesOptions{
			State:         &state,
			CreatedAfter:  &creationDates[0],
			CreatedBefore: &creationDates[1],
			UpdatedAfter:  &updatedDates[0],
			UpdatedBefore: &updatedDates[1],
		}
		Issues, _ = g.GetIssuesWithOptions(opt)
	} else {
		searchNameSpaces := true
		opt := &gitlab.ListProjectsOptions{
			Search: &location,
			SearchNamespaces: &searchNameSpaces,
		}
		projects, _ := g.ListProjectsWithOptions(opt)
		for _, proj := range projects {
			Issues, _ = g.ListAllProjectIssues(proj.ID)
		}
		//projects, _ := g.ListAllProjects()
		//for _, proj := range projects {
		//	if proj.PathWithNamespace == location {
		//		Issues, _ = g.ListAllProjectIssues(proj.ID)//TODO change this to meet opts
		//		break
		//	}
		//}
	}

	for _, issue := range Issues {
		needMilestoneAndLabel, needMilestoneHasLabel, hasLabelState, needLabelStateResolved := false, false, false, false

		if issue.Labels == nil {
			if issue.Milestone == nil {
				needMilestoneAndLabel = true
			}
		}

		for _, label := range issue.Labels {
			if label == "management" || label == "meeting" || label == "standup" {
				break
			}

			if label == "state::in-progress" {
				if issue.Milestone == nil {
					needMilestoneHasLabel = true
				}
			}

			if strings.Contains(strings.ToLower(label), "state::") {
				hasLabelState = true
			}

			if issue.State == "closed" {
				if strings.Contains(strings.ToLower(label), "state::") {
					if label != "state::resolved" && label != "state::abandoned" && label != "state::moved" {
						needLabelStateResolved = true
					}
				}
			}
		}

		if needMilestoneAndLabel {
			issuesInReport = append(issuesInReport, IssueReport{issue, " This issue has no milestones or labels set."})
		}
		if needMilestoneHasLabel {
			issuesInReport = append(issuesInReport, IssueReport{issue, " This issue is in-progress, but has no milestone."})
		}
		if !hasLabelState {
			issuesInReport = append(issuesInReport, IssueReport{issue, " This issue does not contain a `state::` label"})
		}
		if needLabelStateResolved {
			issuesInReport = append(issuesInReport, IssueReport{issue, " This issue is requires the `resolved` or `abandoned` label."})
		}
	}

	csvfile, err :=os.OpenFile("IssueReport.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	headers := []string{"Issue Title", "Issue URL", "Author", "Reason"}
	writer.Write(headers)

	for _, i := range issuesInReport {
		record := []string{i.issue.Title, i.issue.WebURL, i.issue.Author.Name, i.reason}
		writer.Write(record)
	}
	fmt.Println("Results printed to file IssueReport.csv")
	writer.Flush()
	csvfile.Close()

	return nil
}

func CheckEpicsWithinGroup(glClient *gitlab.Client, location string, cd string, ud string, state string) error {
	g := gl.New(glClient)
	epics := make([]EpicReport, 0)
	creationDates := GetTimeParameters(cd)
	updatedDates := GetTimeParameters(ud)
	var groupEpics []*gitlab.Epic

	if state == "" {
		state = "all"
	}

	if location == "" {
		opt := &gitlab.ListGroupEpicsOptions{
			State:         &state,
			CreatedAfter:  &creationDates[0],
			CreatedBefore: &creationDates[1],
			UpdatedAfter:  &updatedDates[0],
			UpdatedBefore: &updatedDates[1],
		}
		groupEpics, _ = g.ListGroupEpicsWithOptions(125, opt)
	} else {
		groups, _ := g.ListAllGroups()
		for _, group := range groups {
			if group.FullPath == location {
				opt := &gitlab.ListGroupEpicsOptions{
					State:         &state,
					CreatedAfter:  &creationDates[0],
					CreatedBefore: &creationDates[1],
					UpdatedAfter:  &updatedDates[0],
					UpdatedBefore: &updatedDates[1],
				}
				groupEpics, _ = g.ListGroupEpicsWithOptions(group.ID, opt)
				break
			}
		}
	}

	for _, epic := range groupEpics {
			if epic.Description == "" {
				epics = append(epics, EpicReport{epic, " This epic has no description"})
			}

			requiresEpicLabel, needLabelStateResolved := false, false

			for _, label := range epic.Labels {
				if label == "management" || label == "meeting" || label == "standup" {
					break
				}

				if strings.Contains(strings.ToLower(label), "epic-") || strings.Contains(strings.ToLower(label), "epic::") {
					requiresEpicLabel = true
				}

				if epic.State == "closed" {
					if strings.Contains(strings.ToLower(label), "state::") {
						if label != "state::resolved" && label != "state::abandoned" {
							needLabelStateResolved = true
						}
					}

				}
			}

			if !requiresEpicLabel {
				epics = append(epics, EpicReport{epic, " This epic does not contain a epic label"})
			}
			if needLabelStateResolved {
				epics = append(epics, EpicReport{epic, " This epic is requires the `resolved` or `abandoned` label."})
			}
	}

	csvfile, err := os.OpenFile("EpicReport.csv", os.O_CREATE & os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	headers := []string{"Epic Name", "Epic URL", "Author", "Reason"}
	writer.Write(headers)

	for _, e := range epics {
		record := []string{e.epic.Title, e.epic.WebURL, e.epic.Author.Name, e.reason}
		writer.Write(record)
	}
	fmt.Println("Results printed to file EpicReport.csv")
	writer.Flush()
	csvfile.Close()

	return nil
}

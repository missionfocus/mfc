package bpe

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/tmetric"
	"github.com/xanzy/go-gitlab"
)

type EpicReport struct {
	epic   *gitlab.Epic
	reason string
}

type IssueReport struct {
	issue  *gitlab.Issue
	reason string
}

func IssueOptsByCheckCommand(glClient *gitlab.Client, location string, cd string, ud string, state string) error {
	g := gl.New(glClient)
	found := false
	creationDates := tmetric.GetTimeParameters(cd)
	updatedDates := tmetric.GetTimeParameters(ud)
	if state == "" {
		state = "all"
	}
	scope := "all"
	if location == "" {
		opts := &gitlab.ListIssuesOptions{
			State:         &state,
			CreatedAfter:  &creationDates[0],
			CreatedBefore: &creationDates[1],
			UpdatedAfter:  &updatedDates[0],
			UpdatedBefore: &updatedDates[1],
			Scope:         &scope,
		}
		CheckIssuesWithOptions(glClient, opts, 0, nil)
	} else {
		searchNameSpaces := true
		opt := &gitlab.ListProjectsOptions{
			Search:           &location,
			SearchNamespaces: &searchNameSpaces,
		}
		projects, _ := g.ListProjectsWithOptions(opt)
		opts := &gitlab.ListProjectIssuesOptions{
			State:         &state,
			CreatedAfter:  &creationDates[0],
			CreatedBefore: &creationDates[1],
			UpdatedAfter:  &updatedDates[0],
			UpdatedBefore: &updatedDates[1],
			Scope:         &scope,
		}
		if projects == nil {
			log.Fatal("error, cannot find a project with the location: " + location)
		}
		if projects[0].PathWithNamespace == location {
			CheckIssuesWithOptions(glClient, nil, projects[0].ID, opts)
		} else {
			//Precautionary: this should not be called on.
			log.Println("Attempting to find project location...")
			for _, project := range projects {
				if project.PathWithNamespace == location {
					CheckIssuesWithOptions(glClient, nil, project.ID, opts)
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("location and path err %s %s" , location)
			}
		}
	}
	return nil
}

func CheckIssuesWithOptions(glClient *gitlab.Client, issueOpts *gitlab.ListIssuesOptions, projID int, projectOpts *gitlab.ListProjectIssuesOptions) error {
	g := gl.New(glClient)
	issuesInReport := make([]IssueReport, 0)
	var Issues []*gitlab.Issue

	if projectOpts == nil && issueOpts != nil {
		Issues, _ = g.GetIssuesWithOptions(issueOpts)
	} else if projectOpts != nil && issueOpts == nil {
		Issues, _ = g.ListAllProjectIssuesWithOpts(projID, projectOpts)
		if Issues == nil {
			return fmt.Errorf("no issues found %v", Issues)
		}
	} else {
		return fmt.Errorf("invalid options, please try again %v %v", issueOpts, projectOpts)
	}

	ignoreIssue := false
	for _, issue := range Issues {
		needMilestoneAndLabel, needMilestoneHasLabel, hasLabelState := false, false, false
		needLabelStateResolved := true

		if issue.Labels == nil {
			if issue.Milestone == nil {
				needMilestoneAndLabel = true
			}
		}
		for _, label := range issue.Labels {
			if label == "state::in-progress" {
				if issue.Milestone == nil {
					needMilestoneHasLabel = true
				}
			}
			if strings.Contains(strings.ToLower(label), "state::") {
				hasLabelState = true
			}
			if issue.State == "closed" {
				if label == "state::resolved" || label == "state::abandoned" || label == "state::moved" {
					needLabelStateResolved = false
				}
			}
		}
		if ignoreIssue {
			break
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
	csvfile, err := os.Create("IssueReport.csv")
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

//TODO make this accept options;
func CheckEpicsWithinGroup(glClient *gitlab.Client, location string, cd string, ud string, state string) error {
	g := gl.New(glClient)
	epics := make([]EpicReport, 0)
	creationDates := tmetric.GetTimeParameters(cd)
	updatedDates := tmetric.GetTimeParameters(ud)
	var groupEpics []*gitlab.Epic

	if state == "" {
		state = "all"
	}
	opt := &gitlab.ListGroupEpicsOptions{
		State:         &state,
		CreatedAfter:  &creationDates[0],
		CreatedBefore: &creationDates[1],
		UpdatedAfter:  &updatedDates[0],
		UpdatedBefore: &updatedDates[1],
	}
	if location == "" {
		groupEpics, _ = g.ListGroupEpicsWithOptions(145, opt)
	} else {
		groups, _ := g.ListAllGroups()
		for _, group := range groups {
			if group.FullPath == location {
				groupEpics, _ = g.ListGroupEpicsWithOptions(group.ID, opt)
				break
			}
		}
	}

	for _, epic := range groupEpics {
		if epic.Description == "" {
			epics = append(epics, EpicReport{epic, " This epic has no description"})
		}
		requiresEpicLabel := false
		ignoreEpic := false
		needLabelStateResolved := true

		for _, label := range epic.Labels {
			if strings.Contains(strings.ToLower(label), "epic-") || strings.Contains(strings.ToLower(label), "epic::") {
				requiresEpicLabel = true
			}
			if epic.State == "closed" {
				if label == "state::resolved" || label == "state::abandoned" || label == "state::moved" {
					needLabelStateResolved = false
				}
			}
		}
		if ignoreEpic {
			break
		}
		if !requiresEpicLabel {
			epics = append(epics, EpicReport{epic, " This epic does not contain a epic label"})
		}
		if needLabelStateResolved {
			epics = append(epics, EpicReport{epic, " This epic is requires the `resolved` or `abandoned` label."})
		}
	}

	csvfile, err := os.Create("EpicReport.csv")
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

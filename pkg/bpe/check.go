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
	var Issues []*gitlab.Issue

	if state == "" {
		state = "all"
	}
	scope := "all"
	if location == "" {
		issueOpts := &gitlab.ListIssuesOptions {
			State:         &state,
			CreatedAfter:  &creationDates[0],
			CreatedBefore: &creationDates[1],
			UpdatedAfter:  &updatedDates[0],
			UpdatedBefore: &updatedDates[1],
			Scope:         &scope,
		}
		Issues, _ = g.GetIssuesWithOptions(issueOpts)
	} else {
		searchNameSpaces := true
		opt := &gitlab.ListProjectsOptions{
			Search:           &location,
			SearchNamespaces: &searchNameSpaces,
		}
		projects, _ := g.ListProjectsWithOptions(opt)
		projectOpts := &gitlab.ListProjectIssuesOptions {
			State:         &state,
			CreatedAfter:  &creationDates[0],
			CreatedBefore: &creationDates[1],
			UpdatedAfter:  &updatedDates[0],
			UpdatedBefore: &updatedDates[1],
			Scope:         &scope,
		}
		if projects == nil {
			log.Fatal("error, cannot find a project with the location: " + location)
		} else if projects[0].PathWithNamespace == location {
			Issues, _ = g.ListAllProjectIssuesWithOpts(projects[0].ID, projectOpts)
		} else {
			//Precautionary: this should not be called on.
			log.Println("Attempting to find project location" + location)
			for _, project := range projects {
				if project.PathWithNamespace == location {
					log.Println("Found project location"  + location)
					Issues, _ = g.ListAllProjectIssuesWithOpts(project.ID, projectOpts)
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("location and path err %s" , location)
			}
		}
	}
	if Issues != nil {
		CheckIssues(Issues)
	} else {
		return fmt.Errorf("no issues found %s" , location)
	}
	return nil
}

func CheckIssues(Issues []*gitlab.Issue) error {
	issuesInReport := make([]IssueReport, 0)
	for _, issue := range Issues {
		missingLabels := false
		missingMilestoneHasLabel := false
		doneWithIssue := false
		needStateLabel := false

		if issue.State == "closed" {
			needStateLabel = true
		}
		if issue.Labels == nil {
			missingLabels = true
		} else {
			for _, label := range issue.Labels {
				if strings.Contains(label, "dta::") || strings.Contains(label, "x-epic-") || strings.Contains(label, "dta-resource"){
					doneWithIssue = true
					break
				}
				if label == "state::in-progress" && issue.Milestone == nil {
					missingMilestoneHasLabel = true
				}
				if issue.State == "closed" {
					if strings.Contains(label, "resolved") || strings.Contains(label, "abandoned")  ||strings.Contains(label, "moved") {
						needStateLabel = false
					}
				}
			}
		}
		if doneWithIssue {
			continue
		} else if issue.Description == "" {
			issuesInReport = append(issuesInReport, IssueReport{issue, " This issue has no description"})
		} else if missingLabels {
			issuesInReport = append(issuesInReport, IssueReport{issue, " This issue has no labels."})
		} else if missingMilestoneHasLabel {
			issuesInReport = append(issuesInReport, IssueReport{issue, " This issue is in-progress, but has no milestone."})
		} else if needStateLabel {
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
	return nil
}

func EpicOptsByCheckCommand(glClient *gitlab.Client, location string, cd string, ud string, state string) error {
	g := gl.New(glClient)
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
		groupEpics, _ = g.ListGroupEpicsWithOptions(codeGroupID, opt)
	} else {
		groups, _ := g.ListAllGroups()
		for _, group := range groups {
			if group.FullPath == location {
				groupEpics, _ = g.ListGroupEpicsWithOptions(group.ID, opt)
				break
			}
		}
	}
	if groupEpics == nil {
		return fmt.Errorf("no epics found %s" , location)
	} else {
		CheckEpics(groupEpics)
	}

	return nil
}

func CheckEpics(groupEpics []*gitlab.Epic) error {
	epics := make([]EpicReport, 0)
	for _, epic := range groupEpics {
		doneWithEpic := false
		missingEpicLabel := true
		for _, label := range epic.Labels {
			if strings.Contains(label, "value-stream") || strings.Contains(label, "mgmt") || strings.Contains(label, "+"){
				doneWithEpic = true
				break
			}
			if strings.Contains(label, "epic-")  {
				missingEpicLabel = false
			}
		}
		if doneWithEpic {
			continue
		} else if epic.Description == "" {
			epics = append(epics, EpicReport{epic, " This epic has no description"})
		} else if missingEpicLabel {
			epics = append(epics, EpicReport{epic, " This epic does not contain a epic label"})
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
	return nil
}

package bpe

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	gl "git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"github.com/asaskevich/govalidator"
	"github.com/xanzy/go-gitlab"
)

type EpicReport struct {
	epic   *gitlab.Epic
	group  *gitlab.Group
	reason string
}

type IssueReport struct {
	issue  *gitlab.Issue
	reason string
}

func CheckIssuesWithinProject(glClient *gitlab.Client, location string, cd string, ud string, state string) error {
	g := gl.New(glClient)
	issues := make([]IssueReport, 0)
	creationDates := GetTimeParameters(cd)
	updatedDates := GetTimeParameters(ud)

	opts := &gitlab.ListIssuesOptions{
		State: &state,
		CreatedAfter: &creationDates[0],
		CreatedBefore: &creationDates[1],
		UpdatedAfter: &updatedDates[0],
		UpdatedBefore: &updatedDates[1],
	}
	queryIssues, _ := g.GetIssuesWithOptions(opts)

	for _, issue := range queryIssues {

		issueIsMeeting := strings.Contains(strings.ToLower(issue.Title), "meeting")
		isStandUp := strings.Contains(strings.ToLower(issue.Title), "stand")
		issueIsManagement := false

		if strings.Contains(strings.ToLower(issue.Title), "management") || strings.Contains(strings.ToLower(issue.Title), "managing") || strings.Contains(strings.ToLower(issue.Title), "manage") || strings.Contains(strings.ToLower(issue.Title), "mgmt") {
			issueIsManagement = true
		}
		if  issueIsManagement || issueIsMeeting || isStandUp {
			continue
		}

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
					if label != "state::resolved" && label != "state::abandoned" && label != "state::moved"{
						needLabelStateResolved = true
					}
				}
			}
		}

		if needMilestoneAndLabel {
			issues = append(issues, IssueReport{issue, " This issue has no milestones or labels set."})
		}
		if needMilestoneHasLabel {
			issues = append(issues, IssueReport{issue, " This issue is in-progress, but has no milestone."})
		}
		if !hasLabelState {
			issues = append(issues, IssueReport{issue, " This issue does not contain a `state::` label"})
		}
		if needLabelStateResolved {
			issues = append(issues, IssueReport{issue, " This issue is requires the `resolved` or `abandoned` label."})
		}
	}

	csvfile, err := os.Create("IssueReport.csv")
	if err != nil {
		return err
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	headers := []string{"Issue Title", "Issue URL", "Author", "Reason"}
	writer.Write(headers)
	writer.Flush()

	csvfile, err = os.OpenFile("IssueReport.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	for _, i := range issues {
		record := []string{i.issue.Title, i.issue.WebURL, i.issue.Author.Name, i.reason}
		writer.Write(record)
	}
	fmt.Println("Results printed to file IssueReport.csv")

	csvfile.Close()

	return nil
}

func CheckEpicsWithinGroup(glClient *gitlab.Client, location string, creationDates string, updatedDates string, status string) error {
	g := gl.New(glClient)
	epics := make([]EpicReport, 0)
	state := SetState(status)
	cd := GetTimeParameters(creationDates)
	ud := GetTimeParameters(updatedDates)

	checkForCD := false
	if cd != nil {
		checkForCD = true
	}
	checkForUD := false
	if ud != nil {
		checkForUD = true
	}

	groups, err := g.ListAllGroups()
	if err != nil {
		return nil
	}

	for _, group := range groups {
		var groupEpics []*gitlab.Epic

		if location == "" && group.Name == "ours" {
			groupEpics, _ = g.ListAllGroupEpics(group.ID) //group.ID for ours is 125
		} else {
			locatonSplit := strings.SplitAfter(location, "https://git.missionfocus.com/")
			if strings.Replace(locatonSplit[0]+"groups/"+locatonSplit[1], " ", "", -1) == group.WebURL {
				groupEpics, _ = g.ListAllGroupEpics(group.ID)
			}
		}

		for _, epic := range groupEpics {
			if epic.State == state || state == "" {
				epicIsMeeting := strings.Contains(strings.ToLower(epic.Title), "meeting")
				isTeamEpic, epicIsManagement := false, false

				if strings.Contains(strings.ToLower(epic.Title), "team") || strings.Contains(strings.ToLower(epic.Title), "stand") || strings.Contains(strings.ToLower(epic.Title), "sustainment") || strings.Contains(strings.ToLower(epic.Title), "planning") {
					isTeamEpic = true
				}

				if strings.Contains(strings.ToLower(epic.Title), "management") || strings.Contains(strings.ToLower(epic.Title), "managing") || strings.Contains(strings.ToLower(epic.Title), "manage") || strings.Contains(strings.ToLower(epic.Title), "mgmt") {
					epicIsManagement = true
				}
				if epicIsManagement || epicIsMeeting || isTeamEpic {
					continue
				}

				if checkForCD {
					if !cd[0].IsZero() && epic.CreatedAt.Before(cd[0]) || !cd[1].IsZero() && epic.CreatedAt.After(cd[1]) {
						continue
					}
				}
				if checkForUD {
					if epic.UpdatedAt.Before(ud[0]) || epic.UpdatedAt.After(ud[1]) {
						continue
					}
				}

				if epic.State == state || state == "" {
					if epic.Description == "" {
						epics = append(epics, EpicReport{epic, group, " This epic has no description"})
					}

					if epic.Description == "## Increment Objectives\n- [ ]  " || epic.Description == "## To Do\n- [ ]   " || epic.Description == "\n**Initial State**\n- SayWhatItIs\n" {
						epics = append(epics, EpicReport{epic, group, " This epic has blank parts of the required template."})
					}

					requiresEpicLabel, needLabelStateResolved := false, false

					for _, label := range epic.Labels {
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
						epics = append(epics, EpicReport{epic, group, " This epic does not contain a epic label"})
					}
					if needLabelStateResolved {
						epics = append(epics, EpicReport{epic, group, " This epic is requires the `resolved` or `abandoned` label."})
					}
				}
			}
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
	writer.Flush()

	csvfile, err = os.OpenFile("EpicReport.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	for _, e := range epics {
		var createEpicURL string
		if location != "" {
			createEpicURL = strings.Replace(e.group.WebURL+"/-/epics/"+govalidator.ToString(e.epic.IID), " ", "", -1)
		} else {
			createEpicURL = strings.Replace(e.group.WebURL+"/-/epics/", " ", "", -1)
		}
		record := []string{e.epic.Title, createEpicURL, e.epic.Author.Name, e.reason}
		writer.Write(record)
	}
	fmt.Println("Results printed to file EpicReport.csv")
	csvfile.Close()

	return nil
}

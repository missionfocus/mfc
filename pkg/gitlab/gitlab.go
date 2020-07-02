package gitlab

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/xanzy/go-gitlab"
)

type GitLab struct {
	client *gitlab.Client
}

//type checkEpicReport struct {
//	title    string
//	groupURL string
//	error    string
//}
//
//type checkIssueReport struct {
//	title string
//	url   string
//	error string
//}

type PostComment struct {
	comment *string
}

func New(client *gitlab.Client) *GitLab {
	return &GitLab{client}
}

// CloneAll clones all `projects` into `directory`, using the project namespace as the directory structure.
func (g *GitLab) CloneAll(projects []*gitlab.Project, directory string, progress io.Writer) error {
	if err := os.MkdirAll(directory, 0777); err != nil {
		return fmt.Errorf("creating base directory: %w", err)
	}

	for _, proj := range projects {
		cloneDir := filepath.Join(directory, proj.PathWithNamespace)
		fmt.Fprintf(progress, "\n==> Cloning %s into %s\n", proj.PathWithNamespace, cloneDir)

		if err := os.MkdirAll(cloneDir, 0777); err != nil {
			return fmt.Errorf("creating repository directory: %w", err)
		}

		_, err := git.PlainClone(cloneDir, false, &git.CloneOptions{
			URL:      proj.SSHURLToRepo,
			Progress: progress,
		})
		if err != nil {
			if err.Error() == "repository already exists" {
				fmt.Fprint(progress, "--> Repository exists, skipping\n")
				continue
			}

			if err.Error() == "remote repository is empty" {
				fmt.Fprint(progress, "--> Repository is empty, skipping\n")
				continue
			}

			return fmt.Errorf("cloning repo: %w", err)
		}
	}

	return nil
}

// ListAllGroups lists all of the groups the caller has access to.
func (g *GitLab) ListAllGroups() ([]*gitlab.Group, error) {
	groups := make([]*gitlab.Group, 0)

	opt := &gitlab.ListGroupsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
		},
	}

	for {
		gs, res, err := g.client.Groups.ListGroups(opt)
		if err != nil {
			return nil, fmt.Errorf("listing groups: %w", err)
		}

		groups = append(groups, gs...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}

	return groups, nil
}

// ListAllProjects lists all of the projects the caller has access to.
func (g *GitLab) ListAllProjects() ([]*gitlab.Project, error) {
	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
		},
	}

	return g.ListAllProjectsWithOptions(opt)
}

// ListAllProjectIssues retrieves all the issues within a project
func (g *GitLab) ListAllProjectIssues(projID interface{}) ([]*gitlab.Issue, error) {
	issues := make([]*gitlab.Issue, 0)

	opt := &gitlab.ListProjectIssuesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
		},
	}

	for {
		is, res, err := g.client.Issues.ListProjectIssues(projID, opt)
		if err != nil {
			return nil, fmt.Errorf("listing project project issues: %w", err)
		}

		issues = append(issues, is...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}

	return issues, nil
}

// ListAllGroupIssues retrieves all the issues within a group
func (g *GitLab) ListAllGroupIssues(projID interface{}) ([]*gitlab.Issue, error) {
	issues := make([]*gitlab.Issue, 0)

	opt := &gitlab.ListGroupIssuesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
		},
	}

	for {
		is, res, err := g.client.Issues.ListGroupIssues(projID, opt)
		if err != nil {
			return nil, fmt.Errorf("listing project group issues: %w", err)
		}

		issues = append(issues, is...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}
	return issues, nil

}

// ListAllGroupEpics returns all the epics related to a group
func (g *GitLab) ListAllGroupEpics(gid interface{}) ([]*gitlab.Epic, error) {
	Epic := make([]*gitlab.Epic, 0)

	opt := &gitlab.ListGroupEpicsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
		},
	}
	for {
		es, res, err := g.client.Epics.ListGroupEpics(gid, opt)
		if err != nil {
			return nil, fmt.Errorf("listing group epics: %w", err)
		}

		Epic = append(Epic, es...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}

	return Epic, nil
}

// ListAllProjects lists all of the projects the caller has access to.
func (g *GitLab) ListAllProjectsWithOptions(opt *gitlab.ListProjectsOptions) ([]*gitlab.Project, error) {
	projects := make([]*gitlab.Project, 0)

	for {
		ps, res, err := g.client.Projects.ListProjects(opt)
		if err != nil {
			return nil, fmt.Errorf("listing projects: %w", err)
		}

		projects = append(projects, ps...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}

	return projects, nil
}

// ListAllProjectsWithRe lists all of the projects the caller has access to and filters by `re`.
func (g *GitLab) ListAllProjectsWithRe(re *regexp.Regexp) ([]*gitlab.Project, error) {
	projects, err := g.ListAllProjects()
	if err != nil {
		return nil, err
	}

	matches := make([]*gitlab.Project, 0)
	for _, proj := range projects {
		if re.MatchString(proj.PathWithNamespace) {
			matches = append(matches, proj)
		}
	}

	return matches, nil
}

func (g *GitLab) GetIssue(projID interface{}, issueID int) (*gitlab.Issue, error) {
	issue, _, err := g.client.Issues.GetIssue(projID, issueID)
	if err != nil {
		return nil, fmt.Errorf("Retrieving issue: %w", err)
	}

	return issue, nil
}

// GetEpic retrieves a specific epic.
func (g *GitLab) GetEpic(gid interface{}, epic int) (*gitlab.Epic, error) {
	Epic, _, err := g.client.Epics.GetEpic(gid, epic)
	if err != nil {
		return nil, fmt.Errorf("Retrieving epic: %w", err)
	}

	return Epic, nil
}

func (g *GitLab) GetMergeRequest(projID interface{}, mergeRequestID int) (*gitlab.MergeRequest, error) {
	mr, _, err := g.client.MergeRequests.GetMergeRequest(projID, mergeRequestID, nil)
	if err != nil {
		return nil, fmt.Errorf("Retrieving issue: %w", err)
	}

	return mr, nil
}

// PostNoteOnIssue posts a "note" (comment) onto an Issue
func (g *GitLab) PostNoteOnIssue(pid interface{}, issue int, message *string) (*gitlab.Note, error) {
	opt := &gitlab.CreateIssueNoteOptions {
		message,
		nil,
	}
	n, _, err := g.client.Notes.CreateIssueNote(pid, issue, opt)
	if err != nil {
		return nil, fmt.Errorf("posting Issue comment: %w", err)
	}

	return n, nil
}

// ListAllIssueNotes returns all the comments within an Issues.
func (g *GitLab) ListAllIssueNotes(pid interface{}, issue int) ([]*gitlab.Note, error) {
	notes := make([]*gitlab.Note, 0)
	opt := &gitlab.ListIssueNotesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
		},
	}

	for {
		nts, res, err := g.client.Notes.ListIssueNotes(pid, issue, opt)
		if err != nil {
			return nil, fmt.Errorf("listing Issue notes: %w", err)
		}

		notes = append(notes, nts...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}

	return notes, nil
}

// PostNoteOnEpic posts a "note" (comment) onto an epic
func (g *GitLab) PostNoteOnEpic(gid interface{}, epic int, message *string) (*gitlab.Note, error) {
	opt := &gitlab.CreateEpicNoteOptions {
		message,
	}
	n, _, err := g.client.Notes.CreateEpicNote(gid, epic, opt)
	if err != nil {
		return nil, fmt.Errorf("posting epic comment: %w", err)
	}

	return n, nil
}

// ListAllEpicNotes returns all the comments within an epic.
func (g *GitLab) ListAllEpicNotes(gid interface{}, epic int) ([]*gitlab.Note, error) {
	notes := make([]*gitlab.Note, 0)
	opt := &gitlab.ListEpicNotesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
		},
	}

	for {
		nts, res, err := g.client.Notes.ListEpicNotes(gid, epic, opt)
		if err != nil {
			return nil, fmt.Errorf("listing epic notes: %w", err)
		}

		notes = append(notes, nts...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}

	return notes, nil
}

func (g *GitLab) GetEpicIssues(gid interface{}, epic int) ([]*gitlab.Issue, error) {
	issues := make([]*gitlab.Issue, 0)

	opt := &gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
	}

	for {
		i, res, err := g.client.EpicIssues.ListEpicIssues(gid, epic, opt)
		if err != nil {
			return nil, fmt.Errorf("listing Issue notes: %w", err)
		}

		issues = append(issues, i...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}

	return issues, nil
}

//TODO GetAllEpicIssues was a test for finding all epics related to a given issue. If this will be used, add a return function
func (g *GitLab) GetAllEpicIssues() error {
	groups, err := g.ListAllGroups()
	if err != nil {
		return err
	}

	for _, group := range groups {
		groupEpics, _ := g.ListAllGroupEpics(group.ID)
		for _, epic := range groupEpics {
			fmt.Println(g.GetEpicIssues(group.ID, epic.ID)) //returns issues that relate to epic, under a group
		}
	}

return nil
}

// SetState ensures that the state can be queried
func SetState(status string) string {
	state := strings.ToLower(status)
	if state == "open" {
		state = "opened"
	}
	if state == "close" {
		state = "closed"
	} else {
		state = ""
	}
	return state
}

const (
	glTimeFormat    = "2006-01-02"
	inputTimeFormat = "01/02/2006"
)

//GetTimeParameters is used to alter the format [date] | [date] into a comparable format
func GetTimeParameters(str string) []time.Time {
	if len(str) == 0 {
		return nil
	}
	date := make([]time.Time, 0)
	splitDateStrings := strings.Split(str, "|")

	for _, d := range splitDateStrings {
		strToDate := strings.Replace(d, " ", "", -1)
		t, _ := time.Parse(inputTimeFormat, strToDate)
		t.Format(glTimeFormat)
		date = append(date, t)
	}

	return date
}

// CheckIssues checks all issues for errors
func (g *GitLab) CheckIssues(location string, creationDates string, updatedDates string, status string) error {
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

	fmt.Println("Gathering all projects...")
	projects, err := g.ListAllProjects()
	if err != nil {
		return nil
	}
	groups, err := g.ListAllGroups()
	if err != nil {
		return nil
	}

	fmt.Println("Finding Issue within project(s)...")
	for _, proj := range projects {
		if proj.WebURL == location || location == "" {
			g.CheckIssuesWithinProject(proj.ID, state, cd, ud, checkForCD, checkForUD)
		} else {
			fmt.Println("Location requested not in projects, attempting to find issue in groups...")
			for _, group := range groups {
				if group.WebURL == location {
					fmt.Println("Found Location. Finding Issues within group.")
					fmt.Println("Warning, if this is a parent group, duplicated result can occur") //TODO fix duplicates...
					g.CheckIssuesWithinGroup(group.ID, state, cd, ud, checkForCD, checkForUD)
					return nil
				} else {
					fmt.Println("Error - Unable to find location given. Please try again.")
					continue
				}
			}
		}
	}

	return nil
}

// CheckIssuesWithinGroup will produce duplicate reports due to subgroups. Also issues are under projects anyways. TODO - Fix error with duplications occuring.
func (g *GitLab) CheckIssuesWithinGroup(groupID int, status string, creationDates []time.Time, updatedDates []time.Time, checkForCD bool, checkForUD bool) error {
	state := SetState(status)
	groupIssues, _ := g.ListAllGroupIssues(groupID)

	for _, issue := range groupIssues {
		if checkForCD {
			if !creationDates[0].IsZero() && issue.CreatedAt.Before(creationDates[0]) || !creationDates[1].IsZero() && issue.CreatedAt.After(creationDates[1]) {
				continue
			}
		}
		if checkForUD {
			fmt.Println(issue.CreatedAt.After(updatedDates[0]))
			if issue.UpdatedAt.Before(updatedDates[0]) || issue.UpdatedAt.After(updatedDates[1]) {
				continue
			}
		}
		if issue.State == state || state == "" {
			if issue.Description == "Acceptance Criteria\n- [ ]   \n- [ ] Automated test: FILEPATHNAME\n- [ ] Pipeline passes with no critical / high vulnerabilities\n" {
				fmt.Println(issue.WebURL, "has not filled out the acceptance criteria.")
			}
			if issue.Description == "" {
				fmt.Println(issue.WebURL, "has no description")
			}

			needMilestoneAndLabel, needMilestoneHasLabel, needLabelHasMilestone, needLabelStateChange := false, false, false, false

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
				//if label != "backlog" || label != "waiting" {
				//
				//}
				if issue.State == "closed" {
					//	if label != "state::resolved" ||  label != "scrum::abandoned" { needLabelHasMilestone = true }
					if label == "state::in-progress" {
						needLabelStateChange = true
					}
				}
			}

			if needMilestoneAndLabel {
				fmt.Println(issue.WebURL, "This issue has no milestones or labels set")
			}
			if needMilestoneHasLabel {
				fmt.Println(issue.WebURL, "This Issue is in-progress, but has no milestone")
			}
			if needLabelHasMilestone {
				fmt.Println(issue.WebURL, "This Issue is missing a `state::abandoned` or `state::resolved` label")
			}
			if needLabelStateChange {
				fmt.Println(issue.WebURL, "This Issue has a `state::in-progress` label, but is closed.")
			}
		}
	}
	return nil
}

func (g *GitLab) CheckIssuesWithinProject(projID int, status string, creationDates []time.Time, updatedDates []time.Time, checkForCD bool, checkForUD bool) error {
	state := SetState(status)

	projIssues, _ := g.ListAllProjectIssues(projID)

	for _, issue := range projIssues {
		if checkForCD {
			if !creationDates[0].IsZero() && issue.CreatedAt.Before(creationDates[0]) || !creationDates[1].IsZero() && issue.CreatedAt.After(creationDates[1]) {
				continue
			}
		}
		if checkForUD {
			fmt.Println(issue.CreatedAt.After(updatedDates[0]))
			if issue.UpdatedAt.Before(updatedDates[0]) || issue.UpdatedAt.After(updatedDates[1]) {
				continue
			}
		}
		if issue.State == state || state == "" {
			if issue.Description == "Acceptance Criteria\n- [ ]   \n- [ ] Automated test: FILEPATHNAME\n- [ ] Pipeline passes with no critical / high vulnerabilities\n" {
				fmt.Println(issue.WebURL, "has not filled out the acceptance criteria.")
			}
			if issue.Description == "" {
				fmt.Println(issue.WebURL, "has no description")
			}

			needMilestoneAndLabel, needMilestoneHasLabel, needLabelHasMilestone, needLabelStateChange := false, false, false, false

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
				//if label != "backlog" || label != "waiting" {
				//
				//}
				if issue.State == "closed" {
					//	if label != "state::resolved" ||  label != "scrum::abandoned" { needLabelHasMilestone = true }
					if label == "state::in-progress" {
						needLabelStateChange = true
					}
				}
			}

			if needMilestoneAndLabel {
				fmt.Println(issue.WebURL, "This issue has no milestones or labels set")
			}
			if needMilestoneHasLabel {
				fmt.Println(issue.WebURL, "This Issue is in-progress, but has no milestone")
			}
			if needLabelHasMilestone {
				fmt.Println(issue.WebURL, "This Issue is missing a `state::abandoned` or `state::resolved` label")
			}
			if needLabelStateChange {
				fmt.Println(issue.WebURL, "This Issue has a `state::in-progress` label, but is closed.")
			}
		}
	}
	return nil
}

func (g *GitLab) CheckEpics(location string, creationDates string, updatedDates string, status string) error {
	state := SetState(status)

	groups, err := g.ListAllGroups()
	if err != nil {
		return nil
	}

	for _, group := range groups {
		groupEpics, _ := g.ListAllGroupEpics(group.ID)
		for _, epic := range groupEpics {
			if epic.State == state || state == "" {

				if epic.Description == "\n**Initial State**\n- SayWhatItIs" {

				}
				if epic.Description == "" {
					fmt.Println("has no description")
				}
				if epic.Labels == nil {
					fmt.Println("Within group ", group.Description, "Epic", epic.Title, "has no labels")
				}
			}
		}
	}
	return nil
}

//TODO UpdateEpicIssues
// *** Thoughts from params: Location (url), label (old label | new label) format, epic (old label | new label), milestone (old, new), status (open/close)
func (g *GitLab) UpdateEpicIssues(location, label, epic, milestone, status string) error {
	state := SetState(status)
	fmt.Println(state) // I did this just so it compiles for now
	return nil
}


func (g *GitLab) PostTest() error {
	groups, err := g.ListAllGroups()
	if err != nil {
		return nil
	}
	projects, err := g.ListAllProjects()
	if err != nil {
		return nil
	}
	//comment :=  "Epic Test Successful!"
	for _, group := range groups {
		groupEpics, _ := g.ListAllGroupEpics(group.ID)
		for _, epic := range groupEpics {
			if group.ID == 161 {
				if epic.ID == 105 {
					if epic.Title == "DMBusey" {
						//g.PostNoteOnEpic(group.ID, epic.IID, &comment)
						fmt.Println("...Posted Epic Comment") // https://git.missionfocus.com/groups/ours/mfm/-/epics/1#note_59800
					}
				}
			}
		}
	}
	//comment =  "Issue Test Successful!"
	for _, proj := range projects {
		projIssues, _ := g.ListAllProjectIssues(proj.ID)
		for _, issue := range projIssues {
			if issue.Title == "MFC GitLab Business Verifications" {
				if issue.ProjectID == 394 {
					//g.PostNoteOnIssue(issue.ProjectID, issue.IID, &comment)
					fmt.Println("...Posted Issue Comment") // https://git.missionfocus.com/ours/code/tools/mfc/-/issues/33
				}
			}
		}
	}
	return nil
}

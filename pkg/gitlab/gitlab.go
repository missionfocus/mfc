package gitlab

import (
	"fmt"
	"github.com/asaskevich/govalidator"
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

type EpicReport struct {
	title    string
	groupURL string
	error    string
}

type IssueReport struct {
	issue  *gitlab.Issue
	reason string
}

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

func (g *GitLab) UpdateIssueWithOpts(pid interface{}, issue int, opt *gitlab.UpdateIssueOptions) (*gitlab.Issue, error) {
	i, _, err := g.client.Issues.UpdateIssue(pid, issue, opt)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (g *GitLab) UpdateEpicWithOpts(gid interface{}, epic int, opt *gitlab.UpdateEpicOptions) (*gitlab.Epic, error) {
	e, _, err := g.client.Epics.UpdateEpic(gid, epic, opt)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// PostNoteOnIssue posts a "note" (comment) onto an Issue
func (g *GitLab) PostNoteOnIssue(pid interface{}, issue int, message *string) (*gitlab.Note, error) {
	opt := &gitlab.CreateIssueNoteOptions{
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
	opt := &gitlab.CreateEpicNoteOptions{
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

//
func (g *GitLab) GetEpicIssues(gid interface{}, epic int) []*gitlab.Issue {
	issues := make([]*gitlab.Issue, 0)

	opt := &gitlab.ListOptions{
		PerPage: 20,
		Page:    1,
	}

	for {
		i, res, err := g.client.EpicIssues.ListEpicIssues(gid, epic, opt)
		if err != nil {
			return nil
		}

		issues = append(issues, i...)

		if res.CurrentPage >= res.TotalPages {
			break
		}
		opt.Page = res.NextPage
	}

	return issues
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
		fmt.Errorf("Retrieving issue: %w", err)
	}

	fmt.Println("Finding Issue within project(s)...")
	for _, proj := range projects {
		if proj.WebURL == location || location == "" {
			g.CheckIssuesWithinProject(proj.ID, state, cd, ud, checkForCD, checkForUD)
		}
	}

	return nil
}

func (g *GitLab) CheckIssuesWithinProject(projID int, status string, creationDates, updatedDates []time.Time, checkForCD, checkForUD bool) error {
	issues := make([]IssueReport, 0)
	state := SetState(status)
	projIssues, _ := g.ListAllProjectIssues(projID)

	for _, issue := range projIssues {
		onlyOurs := strings.Contains(issue.WebURL, "/code")

		issueIsMeeting := strings.Contains(strings.ToLower(issue.Title), "meeting")
		isStandUp := strings.Contains(strings.ToLower(issue.Title), "stand")
		issueIsManagement := false

		if strings.Contains(strings.ToLower(issue.Title), "management") || strings.Contains(strings.ToLower(issue.Title), "managing") || strings.Contains(strings.ToLower(issue.Title), "manage") || strings.Contains(strings.ToLower(issue.Title), "mgmt") {
			issueIsManagement = true
		}

		if !onlyOurs || issueIsManagement || issueIsMeeting || isStandUp {
			continue
		}

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
				issues = append(issues, IssueReport{issue, " This issue has not filled out the acceptance criteria."})
			}
			if issue.Description == "" {
				issues = append(issues, IssueReport{issue, " This issue has no description"})
			}

			needMilestoneAndLabel, needMilestoneHasLabel, hasLabelState, needLabelStateResolved := false, false, false, false

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
					if strings.Contains(strings.ToLower(label), "state::") {
						if label != "state::resolved" && label != "state::abandoned" {
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
	}

	for _, i := range issues {

		fmt.Printf("[%10s](%10s)%10s\n", i.issue.Title, i.issue.WebURL, i.reason)

		//comment := strings.Replace("@" + govalidator.ToString(i.issue.Author), " " , "", -1) + i.reason
		//g.PostNoteOnIssue(i.issue.ProjectID, i.issue.IID, &comment
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
func (g *GitLab) UpdateEpicIssuesLabels(location, label string) error {
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
					if splitURL[1] ==  govalidator.ToString(epic.IID) {
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
			if locationFound {break}
		}
		if locationFound {break}
	}
	return nil
}

func (g *GitLab) PostTest() error {
	groups, err := g.ListAllGroups()
	if err != nil {
		return err
	}
	projects, err := g.ListAllProjects()
	if err != nil {
		return err
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

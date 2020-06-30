package gitlab

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/xanzy/go-gitlab"
)

type GitLab struct {
	client *gitlab.Client
}

type checkEpicReport struct {
	title    string
	groupURL string
	error    string
}

type checkIssueReport struct {
	title string
	url   string
	error string
}

type postComment struct {
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

// GetProjectIssues retrieves all the issues within a project TODO refactor to ListAllProjectIssues
func (g *GitLab) GetProjectIssues(projID interface{}) ([]*gitlab.Issue, error) {
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

// GetGroupIssues retrieves all the issues within a group TODO refactor to ListAllGroupIssues
func (g *GitLab) GetGroupIssues(projID interface{}) ([]*gitlab.Issue, error) {
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

// GetGroupEpics returns all the epics related to a group TODO refactor to ListAllGroupEpics
func (g *GitLab) GetGroupEpics(gid interface{}) ([]*gitlab.Epic, error) {
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

// TODO ListAllEpicIssues - https://github.com/xanzy/go-gitlab/blob/master/epic_issues.go (not in current release)
func (g *GitLab) ListAllEpicIssues(gid int, epic int) ([]*gitlab.Issue, error) {
	issues := make([]*gitlab.Issue, 0)

	return issues, nil
}

// TODO ListAllSubEpics

// CheckIssues checks all issues for errors ---- TODO needs to be refactored.
func (g *GitLab) CheckIssues(location string, dates string, open bool) error {
	projects, err := g.ListAllProjects()
	if err != nil {
		return nil
	}
	groups, err := g.ListAllGroups()
	if err != nil {
		return nil
	}
	for _, proj := range projects {
		projIssues, _ := g.GetProjectIssues(proj.ID)
		for _, issue := range projIssues {
			if issue.Labels == nil {
				//fmt.Println("Issue", issue.Title, issue.WebURL, "has no labels")
			}
			if issue.Description == "" {
				//fmt.Println("Issue", issue.Title, issue.WebURL, "has no description")
			}
			if issue.Milestone == nil {
				missingMilestoneAndLabel := true
				labels := issue.Labels
				for _, label := range labels {
					if label == "backlog" || label == "waiting" {
						missingMilestoneAndLabel = false
					}
				}
				if missingMilestoneAndLabel {
					//fmt.Println("Issue", issue.Title, issue.WebURL, "has no milestones or labels set")
				}
			} else {
				if issue.Weight == 0 {
					//fmt.Println("Issue", issue.Title, issue.WebURL, "has no weight set")
				}
			}
		}
	}
	for _, group := range groups {
		groupIssues, _ := g.GetGroupIssues(group.ID)
		for _, issue := range groupIssues {
			if issue.Labels == nil {
				//fmt.Println("Issue", issue.Title, issue.WebURL, "has no labels")
			}
			if issue.Description == "" {
				//fmt.Println("Issue", issue.Title, issue.WebURL, "has no description")
			}
			if issue.Milestone == nil {
				missingMilestoneAndLabel := true
				labels := issue.Labels
				for _, label := range labels {
					if label == "backlog" || label == "waiting" {
						missingMilestoneAndLabel = false
					}
				}
				if missingMilestoneAndLabel {
					//fmt.Println("Issue", issue.Title, issue.WebURL, "has no milestones or labels set")
				}
			} else {
				if issue.Weight == 0 {
					//fmt.Println("Issue", issue.Title, issue.WebURL, "has no weight set")
				}
			}
		}
	}
	return nil
}

func (g *GitLab) CheckEpics(location string, dates string, open bool) error {
	groups, err := g.ListAllGroups()
	if err != nil {
		return nil
	}

	for _, group := range groups {
		groupEpics, _ := g.GetGroupEpics(group.ID)
		for _, epic := range groupEpics {
			fmt.Println(epic.State)
			fmt.Println(group.WebURL)
			if epic.Labels == nil {
				fmt.Println("Within group ", group.Description, "Epic", epic.Title, "has no labels")
			}
			if epic.Description == "" {
				fmt.Println("Within group ", group.Description, "Epic,", epic.Title, "has no description")
			}
		}
	}
	return nil
}

// TODO CheckAllIssuesAndEpics
func (g *GitLab) CheckAllIssuesAndEpics(location string, dates string, open bool) error {

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
		groupEpics, _ := g.GetGroupEpics(group.ID)
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
		projIssues, _ := g.GetProjectIssues(proj.ID)
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

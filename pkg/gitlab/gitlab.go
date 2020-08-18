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

func (g *GitLab) GetIssue(projID interface{}, issueID int) (*gitlab.Issue, error) {
	issue, _, err := g.client.Issues.GetIssue(projID, issueID)
	if err != nil {
		return nil, fmt.Errorf("Retrieving issue: %w", err)
	}

	return issue, nil
}

func (g *GitLab) GetIssuesWithOptions(opt *gitlab.ListIssuesOptions) ([]*gitlab.Issue, error) {
	issues := make([]*gitlab.Issue, 0)

	for {
		is, res, err := g.client.Issues.ListIssues(opt)
		if err != nil {
			return nil, fmt.Errorf("Retrieving issues with options: %w", err)
		}
		if res.CurrentPage >= res.TotalPages {
			break
		}
		issues = append(issues, is...)
		opt.Page = res.NextPage
	}
	return issues, nil
}

// ListAllGroupEpicsWithOptions returns epics within a specificed group and meets specified options.
func (g *GitLab) ListGroupEpicsWithOptions(gid interface{}, opt *gitlab.ListGroupEpicsOptions) ([]*gitlab.Epic, error) {
	epics := make([]*gitlab.Epic, 0)

	for {
		e, res, err := g.client.Epics.ListGroupEpics(gid, opt)
		if err != nil {
			return nil, fmt.Errorf("listing group epics with options: %w", err)
		}
		if res.CurrentPage >= res.TotalPages {
			break
		}
		epics = append(epics, e...)
		opt.Page = res.NextPage
	}

	return epics, nil
}

// GetEpic retrieves a specific epic.
func (g *GitLab) GetEpic(gid interface{}, epic int) (*gitlab.Epic, error) {
	Epic, _, err := g.client.Epics.GetEpic(gid, epic)
	if err != nil {
		return nil, fmt.Errorf("Retrieving epic: %w", err)
	}

	return Epic, nil
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

func (g *GitLab) ListSubGroups(groupID int) ([]*gitlab.Group, error) {
	groups := make([]*gitlab.Group, 0)
	opt := &gitlab.ListSubgroupsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 20,
			Page:    1,
		},
	}
	for {
		gs, res, err := g.client.Groups.ListSubgroups(groupID, opt)
		if err != nil {
			return nil, fmt.Errorf("listing subgroups: %w", err)
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

	return g.ListProjectsWithOptions(opt)
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

func (g *GitLab) ListAllProjectIssuesWithOpts(projID interface{}, opt *gitlab.ListProjectIssuesOptions) ([]*gitlab.Issue, error) {
	issues := make([]*gitlab.Issue, 0)
	for {
		is, res, err := g.client.Issues.ListProjectIssues(projID, opt)
		if err != nil {
			return nil, fmt.Errorf("listing project project issues with opts: %w", err)
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
func (g *GitLab) ListProjectsWithOptions(opt *gitlab.ListProjectsOptions) ([]*gitlab.Project, error) {
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

func (g *GitLab) GetEpicLinks(gid interface{}, epic int) []*gitlab.Epic {
	epics, _, err := g.client.Epics.GetEpicLinks(gid, epic)
	if err != nil {
		return nil
	}

	return epics
}

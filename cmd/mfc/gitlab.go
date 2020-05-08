package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

func init() {
	mfcCmd.AddCommand(gitlabCmd)
}

var gitlabCmd = &cobra.Command{
	Use:     "gitlab",
	Short:   "Interact with GitLab",
	Aliases: []string{"gl"},
}

/** GitLab utility functions */

const gitlabBaseURL = "https://git.missionfocus.com"

func getGitLabClient() (*gitlab.Client, error) {
	baseURLOpt := gitlab.WithBaseURL(gitlabBaseURL)
	if token := os.Getenv("GITLAB_PAT"); token != "" {
		return gitlab.NewClient(token, baseURLOpt)
	}

	tokenFile := filepath.Join(homeDir(), ".gitlab-pat")
	_, err := os.Stat(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("failed to get GitLab client: %w", err)
	}

	txt, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("could not read GitLab token file: %w", err)
	}
	return gitlab.NewClient(strings.TrimSpace(string(txt)), baseURLOpt)
}

package main

import (
	"errors"
	"os"

	"github.com/missionfocus/mfc/pkg/vault"
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

func getGitLabClient(vault vault.Vault) (*gitlab.Client, error) {
	baseURLOpt := gitlab.WithBaseURL(gitlabBaseURL)
	if token := os.Getenv("GITLAB_PAT"); token != "" {
		return gitlab.NewClient(token, baseURLOpt)
	}

	secret, err := vault.KVUserGet("gitlab")
	if err != nil {
		return nil, err
	}
	if secret == nil {
		return nil, errors.New("missing GitLab PAT, you may need to set it with `mfc gitlab set-token`")
	}

	tok := secret.Data["data"].(map[string]interface{})["token"].(string)
	return gitlab.NewClient(tok, baseURLOpt)
}

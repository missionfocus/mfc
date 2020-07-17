package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	gitlabCmd.AddCommand(gitlabUpdateCmd)
	gitlabUpdateCmd.AddCommand(gitlabUpdateEpicIssuesLabelCmd)
	gitlabUpdateCmd.AddCommand(gitlabUpdateAllLabelsCmd)
}


const gitlabUpdateEpicIssueLabels = `
  mfc gitlab update eil "https://git.missionfocus.com/groups/ours/mfm/-/epics/1" "dev::coding|"      					# Deletes dev::coding label and adds no label in place.
  mfc gitlab update eil "https://git.missionfocus.com/ours/mfm/mfm-records/-/issues/5" "check-this|dev::coding"         # Removes check-this label and adds dev:coding label
`

var gitlabUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Gitlab check <cmd>",
	Aliases: []string{"u"},
}

var gitlabUpdateEpicIssuesLabelCmd = &cobra.Command {
	Use:     "EpicIssuesLabel <location> <OldLabel|NewLabel>",
	Short:   "Update a specific epic and issues labels",
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"eil"},
	Example: gitlabUpdateEpicIssueLabels,
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		if args[0] == "" || args[1] == "" {
			return
		}

		check(gl.UpdateEpicIssuesLabels(args[0], args[1]))
	},
}

var gitlabUpdateAllLabelsCmd = &cobra.Command {
	Use:     "all-labels",
	Short:   "Update all Epics and Children labels",
	Args:    cobra.ExactArgs(0),
	Aliases: []string{"eal"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		check(gl.UpdateAllLabels())
	},
}


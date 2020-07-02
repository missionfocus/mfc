package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	gitlabCmd.AddCommand(gitlabUpdateCmd)
	gitlabUpdateCmd.AddCommand(gitlabUpdateEpicIssuesCmd)

	gitlabUpdateEpicIssuesCmd.PersistentFlags().StringVarP(&gitlabLocation, "Location", "l", "", "Define a location")
	gitlabUpdateEpicIssuesCmd.PersistentFlags().StringVarP(&gitlabLabel, "Label", "", "", "(Removes) Old Label|New Label (Adds)")
	gitlabUpdateEpicIssuesCmd.PersistentFlags().StringVarP(&gitlabStatus, "Status", "s", "", "Retrieve only closed/open issues and/or epics")
}

// If variables are used within the gitlab_check file, should they be moved into a parent class?  ---> gitlab.go
var (
	gitlabLabel			string
	gitlabStatus			string
)


var gitlabUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Gitlab check <cmd>",
	Aliases: []string{"u"},
}

var gitlabUpdateEpicIssuesCmd = &cobra.Command {
	Use:     "epicissues",
	Short:   "Check all issues and epics.",
	Aliases: []string{"ei"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		check(gl.GetAllEpicIssues())
	},
}
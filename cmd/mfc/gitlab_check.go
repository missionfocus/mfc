package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	gitlabCmd.AddCommand(gitlabCheckCmd)
	gitlabCheckCmd.AddCommand(gitlabCheckIssuesAndEpicsCmd)
	gitlabCheckCmd.AddCommand(gitlabCheckIssuesCmd)
	gitlabCheckCmd.AddCommand(gitlabCheckEpicsCmd)
	gitlabCheckCmd.AddCommand(gitlabPostTestCmd)


	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabLocation, "Location", "l", "", "Define a location")
	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabCreationDate, "CreationDate", "c", "", "AfterDate|BeforeDate")
	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabUpdatedDate, "UpdateDate", "u", "", "AfterDate|BeforeDate")
	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabStatus, "Status", "s", "", "Retrieve only closed/open issues and/or epics")
}

var (
	gitlabLocation			string
	gitlabCreationDate 		string
	gitlabUpdatedDate		string
	gitlabStatus			string
)

const gitlabCheckExample = `
  mfc gitlab check 
`

var gitlabCheckCmd = &cobra.Command{
	Use:     "check",
	Short:   "Gitlab check <cmd>",
	Aliases: []string{"ck"},
}

var gitlabCheckIssuesAndEpicsCmd = &cobra.Command {
	Use:     "all",
	Short:   "Check all issues and epics.",
	Aliases: []string{"all"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		check(gl.CheckEpics(gitlabLocation, gitlabCreationDate, gitlabUpdatedDate, gitlabStatus))
		check(gl.CheckIssues(gitlabLocation, gitlabCreationDate, gitlabUpdatedDate, gitlabStatus))
	},
}

var gitlabCheckEpicsCmd = &cobra.Command {
	Use:     "epics",
	Short:   "Check epics for errors",
	Aliases: []string{"e"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		check(gl.CheckEpics(gitlabLocation, gitlabCreationDate, gitlabUpdatedDate, gitlabStatus))
	},
}

var gitlabCheckIssuesCmd = &cobra.Command {
	Use:     "issues",
	Short:   "Check issues for errors",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		check(gl.CheckIssues(gitlabLocation, gitlabCreationDate, gitlabUpdatedDate, gitlabStatus))
	},
}

var gitlabPostTestCmd = &cobra.Command {
	Use:     "post",
	Short:   "Post Test",
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		check(gl.PostTest())
	},
}
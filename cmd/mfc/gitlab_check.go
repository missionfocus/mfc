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
	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabDate, "Date", "d", "", "Define a date Date | Date")
	gitlabCheckCmd.PersistentFlags().BoolVarP(&gitlabOpenOnly, "Open", "o", false,"")
}

var (
	gitlabLocation	string
	gitlabDate 		string
	gitlabOpenOnly	bool
)

const gitlabCheckExample = `
  mfc gitlab check 
`

var gitlabCheckCmd = &cobra.Command{
	Use:     "check",
	Short:   "Gitlab check <cmd>",
	Aliases: []string{"ck"},
}

var gitlabCheckIssuesCmd = &cobra.Command {
	Use:     "issuesepics",
	Short:   "Check all issues and epics.",
	Aliases: []string{"ie"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		check(gl.CheckAllIssuesAndEpics(gitlabLocation, gitlabDate, gitlabOpenOnly))
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

		check(gl.CheckEpics(gitlabLocation, gitlabDate, gitlabOpenOnly))
	},
}

var gitlabCheckIssuesAndEpicsCmd = &cobra.Command {
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

		check(gl.CheckIssues(gitlabLocation, gitlabDate, gitlabOpenOnly))
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
package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/bpe"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	gitlabCmd.AddCommand(gitlabCheckCmd)
	gitlabCheckCmd.AddCommand(gitlabCheckIssuesAndEpicsCmd)
	gitlabCheckCmd.AddCommand(gitlabCheckIssuesCmd)
	gitlabCheckCmd.AddCommand(gitlabCheckEpicsCmd)
	gitlabCheckCmd.AddCommand(gitlabVelocityReportCmd)


	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabCheckLocation, "Location", "l", "", "Define a location")
	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabCheckCreationDate, "CreationDate", "c", "", "AfterDate|BeforeDate")
	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabCheckUpdatedDate, "UpdateDate", "u", "", "AfterDate|BeforeDate")
	gitlabCheckCmd.PersistentFlags().StringVarP(&gitlabCheckStatus, "Status", "s", "", "Retrieve only closed/open issues and/or epics")
}

var (
	gitlabCheckLocation			string
	gitlabCheckCreationDate 	string
	gitlabCheckUpdatedDate		string
	gitlabCheckStatus			string
)

const gitlabCheckExample = `
  mfc gitlab check 
`
const gitlabVelocityExample = `
    mfc gitlab check velocity  ""		# Generate documentation for mfc in markdown format
 	mfc gitlab check velocity "GDAC 18" # Gets the velocity for the GDAC 18 milestone`


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

		check(gl.CheckEpicsWithinGroup(gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
		check(gl.CheckIssuesWithinProject(gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
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

		check(gl.CheckEpicsWithinGroup(gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
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

		check(gl.CheckIssuesWithinProject(gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
	},
}

var gitlabVelocityReportCmd = &cobra.Command {
	Use:     "velocity",
	Short:   "Check each MFM velocity for a given milestone and iteration",
	Example: gitlabVelocityExample,
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)

		check(bpe.VelocityReport(client, args[0], ""))
	},
}
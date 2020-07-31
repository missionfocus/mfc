package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/bpe"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	gitlabCmd.AddCommand(bpeCheckCmd)
	bpeCheckCmd.AddCommand(bpeCheckIssuesAndEpicsCmd)
	bpeCheckCmd.AddCommand(bpeCheckIssuesCmd)
	bpeCheckCmd.AddCommand(bpeCheckEpicsCmd)

	bpeCheckCmd.PersistentFlags().StringVarP(&gitlabCheckLocation, "Location", "l", "", "Define a location")
	bpeCheckCmd.PersistentFlags().StringVarP(&gitlabCheckCreationDate, "CreationDate", "c", "", "AfterDate|BeforeDate")
	bpeCheckCmd.PersistentFlags().StringVarP(&gitlabCheckUpdatedDate, "UpdateDate", "u", "", "AfterDate|BeforeDate")
	bpeCheckCmd.PersistentFlags().StringVarP(&gitlabCheckStatus, "Status", "s", "", "Retrieve only closed/open issues and/or epics")
}

var (
	gitlabCheckLocation     string
	gitlabCheckCreationDate string
	gitlabCheckUpdatedDate  string
	gitlabCheckStatus       string
)

const gitlabCheckExample = `
  mfc gitlab check 
`

var bpeCheckCmd = &cobra.Command{
	Use:     "check",
	Short:   "Gitlab check <cmd>",
	Aliases: []string{"ck"},
}

var bpeCheckIssuesAndEpicsCmd = &cobra.Command{
	Use:     "all",
	Short:   "Check all issues and epics.",
	Aliases: []string{"all"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)

		check(bpe.CheckEpicsWithinGroup(client, gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
		check(bpe.CheckIssuesWithinProject(client, gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
	},
}

var bpeCheckEpicsCmd = &cobra.Command{
	Use:     "epics",
	Short:   "Check epics for errors",
	Aliases: []string{"e"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)

		check(bpe.CheckEpicsWithinGroup(client, gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
	},
}

var bpeCheckIssuesCmd = &cobra.Command{
	Use:     "issues",
	Short:   "Check issues for errors",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)

		check(bpe.CheckIssuesWithinProject(client, gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
	},
}

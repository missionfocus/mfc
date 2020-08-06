package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/bpe"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	bpeCmd.AddCommand(bpeCheckCmd)
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

const bpeCheckEpicsExample = `
  mfc bpe check epics									# By default checks all epics within /ours.
  mfc bpe check epics -l "ours/code"						# Checks epics within the group /code.
  mfc bpe check epics -c "1999-12-31|2020-1-1"			# Checks epics in-between the dates of December 31st, 1999 and January 1st, 2020.
`

const bpeCheckIssuesExample = `
  mfc bpe check issues									# By default checks all issues
  mfc bpe check issues -l "ours/code/tools/mfc"			# Check issues that match the project path of ours/code/tools/mfc
  mfc bpe check issues -c "1999-12-31|2020-1-1"			# Checks issues in-between the dates of December 31st, 1999 and January 1st, 2020.
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
	Example: bpeCheckEpicsExample,
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
	Example: bpeCheckIssuesExample,
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)

		check(bpe.CheckIssuesWithinProject(client, gitlabCheckLocation, gitlabCheckCreationDate, gitlabCheckUpdatedDate, gitlabCheckStatus))
	},
}

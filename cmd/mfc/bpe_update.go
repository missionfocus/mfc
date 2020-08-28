package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/bpe"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	bpeCmd.AddCommand(bpeUpdateCmd)
	bpeUpdateCmd.AddCommand(bpeUpdateEpicIssuesLabelCmd)
	bpeUpdateCmd.AddCommand(bpeUpdateAllLabelsCmd)

	bpeUpdateEpicIssuesLabelCmd.PersistentFlags().BoolVar(&bpeIncludeChildren, "children", false, "Include child Epic and Issues in label update.")
}

var (
	bpeIncludeChildren bool
)

const bpeUpdateEpicIssueLabels = `
  mfc bpe update eil "https://git.missionfocus.com/groups/ours/mfm/-/epics/1" "dev::coding|"      			# Deletes dev::coding label and adds no label in place.
  mfc bpe update eil "https://git.missionfocus.com/groups/ours/mfm/-/epics/1" "check-this|dev::coding"         # Removes check-this label and adds dev:coding label
  mfc bpe update eil "https://git.missionfocus.com/groups/ours/mfm/-/epics/1" "|dev::coding" --children         # Removes dev::coding from the epic and all the sub-epics and issues.
`

var bpeUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Gitlab check <cmd>",
	Aliases: []string{"u"},
}

var bpeUpdateEpicIssuesLabelCmd = &cobra.Command{
	Use:     "EpicIssuesLabel <location> <OldLabel|NewLabel>",
	Short:   "Update a specific epic and issues labels",
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"eil"},
	Example: bpeUpdateEpicIssueLabels,
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)

		check(bpe.UpdateEpicIssuesWith(client, args[0], args[1], bpeIncludeChildren))
	},
}

var bpeUpdateAllLabelsCmd = &cobra.Command{
	Use:   "all-labels",
	Short: "Update all Epics and Children labels",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)

		check(bpe.UpdateAllEpicLabels(client))
	},
}

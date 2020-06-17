package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/gitlab"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	gitlabCmd.AddCommand(gitlabCloneCmd)

	gitlabCloneCmd.PersistentFlags().StringVar(&gitlabCloneRoot, "root", "/opt", "Path to the base directory to clone into")
	gitlabCloneCmd.PersistentFlags().BoolVar(&gitlabCloneNoPrompt, "no-prompt", false, "Disable the confirmation prompt")
}

var (
	gitlabCloneRoot     string
	gitlabCloneNoPrompt bool
)

const gitlabCloneExample = `
  mfc gitlab clone ^ours/              # Clones the entire ours tree
  mfc gitlab clone ^ours/code/tools/   # Clones just the projects under the ours/code/tools group
  mfc gitlab clone ours/code/tools/mfc # Clones only the ours/code/tools/mfc project
`

var gitlabCloneCmd = &cobra.Command{
	Use:     "clone <regexp>",
	Short:   "Clone all repositories matching the specified regexp",
	Args:    cobra.ExactArgs(1),
	Example: gitlabCloneExample,
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)
		gl := gitlab.New(client)

		fmt.Printf("==> Listing projects that match %s\n", args[0])
		projects, err := gl.ListAllProjectsWithRe(regexp.MustCompile(args[0]))
		check(err)
		for _, proj := range projects {
			fmt.Println(proj.PathWithNamespace)
		}

		if !gitlabCloneNoPrompt {
			prompt := promptui.Prompt{Label: "Clone repositories", IsConfirm: true}
			result, err := prompt.Run()
			check(err)
			if strings.ToLower(result) != "y" {
				return
			}
		}

		check(gl.CloneAll(projects, gitlabCloneRoot, os.Stdout))
	},
}

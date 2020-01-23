package mf_vault

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	// Cobra's bash generated zsh script is not usable in place.
	// https://github.com/spf13/cobra/pull/646#issuecomment-498728573
	zsh "github.com/rsteube/cobra-zsh-gen"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configDocsCmd)
	configCmd.AddCommand(completionCmd)

	wd, err := os.Getwd()
	check(err)
	configDocsCmd.PersistentFlags().StringVarP(&configDocsPath, "path", "p", filepath.Join(wd, "docs"), "path to where generated documentation will be written")
	configDocsCmd.PersistentFlags().StringVarP(&configDocsFormat, "format", "f", "md", `format for generated docs, valid values are "md", "rst", or "man"`)
	completionCmd.PersistentFlags().StringVarP(&completionShell, "shell", "s", "bash", "Provide shell to generate completions for: zsh, oh-my-zsh, bash. Defaults to bash.")
	completionCmd.PersistentFlags().BoolVarP(&completionWriteToFile, "write", "w", false, "Write autocompletion files")
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the mf-vault CLI",
}

var (
	configDocsPath        string
	configDocsFormat      string
	completionShell       string
	completionWriteToFile bool
)

var configDocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for mf-vault",
	Run: func(cmd *cobra.Command, args []string) {
		check(os.MkdirAll(configDocsPath, 0700))
		switch configDocsFormat {
		case "md":
			check(doc.GenMarkdownTree(rootCmd, configDocsPath))
		case "rst":
			check(doc.GenReSTTree(rootCmd, configDocsPath))
		case "man":
			header := &doc.GenManHeader{
				Title:   "MF-VAULT",
				Section: "1",
			}
			check(doc.GenManTree(rootCmd, header, configDocsPath))
		}
	},
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates completion script for mf-vault.",
	Run: func(cmd *cobra.Command, args []string) {
		switch completionShell {
		case "bash", "zsh":
			completionsDir := path.Join(homeDir(), ".completions.d", completionShell)
			completionsFile := path.Join(completionsDir, "_mf-vault")
			if completionWriteToFile {
				check(os.MkdirAll(completionsDir, 0700))
				switch completionShell {
				case "bash":
					check(rootCmd.GenBashCompletionFile(completionsFile))
				case "zsh":
					zsh.Wrap(rootCmd).GenZshCompletionFile(completionsFile)
				}
				rcFile := path.Join(homeDir(), "."+completionShell+"rc")
				fmt.Printf("For first run:\n`echo . %s >> %s`.\n\n", completionsFile, rcFile)
			} else {
				switch completionShell {
				case "bash":
					check(rootCmd.GenBashCompletion(os.Stdout))
				case "zsh":
					zsh.Wrap(rootCmd).GenZshCompletion(os.Stdout)
				}
			}
		case "oh-my-zsh":
			completionsDir := path.Join(homeDir(), ".oh-my-zsh", "completions")
			if completionWriteToFile {
				check(os.MkdirAll(completionsDir, 0700))
				check(rootCmd.GenZshCompletionFile(path.Join(completionsDir, "_mf-vault")))
			} else {
				rootCmd.GenZshCompletion(os.Stdout)
			}
		default:
			fmt.Println("Shell not supported. Review help for options.")
			cmd.Help()
			return
		}
		if completionWriteToFile {
			fmt.Printf("Updated %s shell completions available in new terminal.\n", completionShell)
		}
	},
}

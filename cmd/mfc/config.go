package main

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
	mfcCmd.AddCommand(configCmd)
	configCmd.AddCommand(configDocsCmd)
	configCmd.AddCommand(configCompletionCmd)

	wd, err := os.Getwd()
	check(err)

	configDocsCmd.PersistentFlags().StringVarP(&configDocsPath, "path", "p", filepath.Join(wd, "docs"), "path to where generated documentation will be written")
	configDocsCmd.PersistentFlags().StringVarP(&configDocsFormat, "format", "f", "md", `format for generated docs, valid values are "md", "rst", or "man"`)

	configCompletionCmd.PersistentFlags().StringVarP(&configCompletionShell, "shell", "s", "bash", "Provide shell to generate completions for: zsh, oh-my-zsh, bash. Defaults to bash.")
	configCompletionCmd.PersistentFlags().BoolVarP(&configCompletionWriteToFile, "write", "w", false, "Write autocompletion files")
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure mfc",
}

var (
	configDocsPath              string
	configDocsFormat            string
	configCompletionShell       string
	configCompletionWriteToFile bool
)

var configDocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for mfc",
	Run: func(cmd *cobra.Command, args []string) {
		check(os.MkdirAll(configDocsPath, 0700))
		switch configDocsFormat {
		case "md":
			check(doc.GenMarkdownTree(mfcCmd, configDocsPath))
		case "rst":
			check(doc.GenReSTTree(mfcCmd, configDocsPath))
		case "man":
			header := &doc.GenManHeader{
				Title:   "MFC",
				Section: "1",
			}
			check(doc.GenManTree(mfcCmd, header, configDocsPath))
		}
	},
}

var configCompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion scripts for mfc",
	Run: func(cmd *cobra.Command, args []string) {
		switch configCompletionShell {
		case "bash", "zsh":
			completionsDir := path.Join(homeDir(), ".completions.d", configCompletionShell)
			completionsFile := path.Join(completionsDir, "_mfc")
			if configCompletionWriteToFile {
				check(os.MkdirAll(completionsDir, 0700))
				switch configCompletionShell {
				case "bash":
					check(mfcCmd.GenBashCompletionFile(completionsFile))
				case "zsh":
					check(zsh.Wrap(mfcCmd).GenZshCompletionFile(completionsFile))
				}
				rcFile := path.Join(homeDir(), "."+configCompletionShell+"rc")
				fmt.Printf("For first run:\n`echo . %s >> %s`.\n\n", completionsFile, rcFile)
			} else {
				switch configCompletionShell {
				case "bash":
					check(mfcCmd.GenBashCompletion(os.Stdout))
				case "zsh":
					check(zsh.Wrap(mfcCmd).GenZshCompletion(os.Stdout))
				}
			}
		case "oh-my-zsh":
			completionsDir := path.Join(homeDir(), ".oh-my-zsh", "completions")
			if configCompletionWriteToFile {
				check(os.MkdirAll(completionsDir, 0700))
				check(mfcCmd.GenZshCompletionFile(path.Join(completionsDir, "_mfc")))
			} else {
				check(mfcCmd.GenZshCompletion(os.Stdout))
			}
		default:
			fmt.Println("Shell not supported. Review help for options.")
			_ = cmd.Help()
			return
		}
		if configCompletionWriteToFile {
			fmt.Printf("Updated %s shell completions available in new terminal.\n", configCompletionShell)
		}
	},
}

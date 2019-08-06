package mf_vault

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configDocsCmd)

	wd, err := os.Getwd()
	check(err)
	configDocsCmd.PersistentFlags().StringVarP(&configDocsPath, "path", "p", filepath.Join(wd, "docs"), "path to where generated documentation will be written")
	configDocsCmd.PersistentFlags().StringVarP(&configDocsFormat, "format", "f", "md", `format for generated docs, valid values are "md", "rst", or "man"`)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the mf-vault CLI",
}

var (
	configDocsPath   string
	configDocsFormat string
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

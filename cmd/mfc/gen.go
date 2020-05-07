package mfc

import (
	"fmt"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/gen"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.AddCommand(genPasswordCmd)

	genPasswordCmd.PersistentFlags().IntVarP(&genPasswordWordCount, "word-count", "w", 5, "count of words to generate")
	genPasswordCmd.PersistentFlags().StringVarP(&genPasswordSeparator, "separator", "s", "-", "separator between words")
	genPasswordCmd.PersistentFlags().IntVarP(&genPasswordDigitCount, "digit-count", "d", 4, "number of digits in random number appended to password, 0 disables")
	genPasswordCmd.PersistentFlags().BoolVarP(&genPasswordShouldCapitalize, "capitalize", "c", true, "capitalizes the first letter of each word")
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate secrets",
}

var (
	genPasswordWordCount        int
	genPasswordSeparator        string
	genPasswordDigitCount       int
	genPasswordShouldCapitalize bool
)

var genPasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate a password using the Diceware algorithm",
	Run: func(cmd *cobra.Command, args []string) {
		pwGen, err := gen.NewPassword()
		check(err)

		pwGen.WordCount = genPasswordWordCount
		pwGen.Separator = genPasswordSeparator
		pwGen.DigitCount = genPasswordDigitCount
		pwGen.ShouldCapitalize = genPasswordShouldCapitalize

		pw, err := pwGen.Generate()
		check(err)
		fmt.Print(pw)
	},
}

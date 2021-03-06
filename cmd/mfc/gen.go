package main

import (
	"fmt"

	"github.com/missionfocus/mfc/pkg/gen"
	"github.com/spf13/cobra"
)

func init() {
	mfcCmd.AddCommand(genCmd)
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

const genPasswordExample = `
  mfc gen password                  # Generates string with randomized dictionary words in the format of Aa-Bb-Cc-Dd-Ee-1234
  mfc gen password -w 4 -s "_" -d 6 # Generates string with randomized dictionary words in the format of Aa_Bb_Cc_Dd-12345
`

var genPasswordCmd = &cobra.Command{
	Use:     "password",
	Short:   "Generate a password using the Diceware algorithm",
	Example: genPasswordExample,
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

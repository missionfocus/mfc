package main

import (
	"fmt"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/autoupdate"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/cobra"
)

func init() {
	mfcCmd.AddCommand(updateCmd)
}

const (
	updateBucket  = "public.missionfocus.com"
	updatePrefix  = "mfc/"
	updateProfile = "missionfocus"
	updateRegion  = "us-east-1"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the mfc binary to the latest release",
	Run: func(cmd *cobra.Command, args []string) {
		sess := session.Must(session.NewSession(&aws.Config{
			Credentials: credentials.NewSharedCredentials(mfcAWSCredentialsPath, updateProfile),
			Region:      aws.String(updateRegion),
		}))
		updater := autoupdate.NewS3Updater(sess, updateBucket, updatePrefix)

		fmt.Println("Checking for update...")
		nextVer, err := updater.Check(version)
		check(err)
		if nextVer == "" {
			fmt.Println("Up to date.")
			return
		}

		fmt.Printf("Updating to v%s...", nextVer)
		check(updater.Update())
		fmt.Print(" done.\n")
	},
}

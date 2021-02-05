package main

import (
	"context"
	"fmt"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func init() {
	mfcCmd.AddCommand(jumpcloudCmd)
	jumpcloudCmd.AddCommand(manageUserSystemSSHCmd)

	manageUserSystemSSHCmd.Flags().StringVarP(&operation, "operation", "o", "", "add, update, or remove SSH properties")
	manageUserSystemSSHCmd.Flags().BoolVarP(&withRoot, "withRoot", "r", false, "Allow/Restrict sudo for user group")
	manageUserSystemSSHCmd.Flags().BoolVarP(&withoutPassSudo, "withoutPassSudo", "p", false, "Allow/Restrict passwordless sudo for user group")
}

var (
	operation       string
	withoutPassSudo bool
	withRoot        bool
)

var jumpcloudCmd = &cobra.Command{
	Use:   "jumpcloud",
	Short: "Interact with Jumpcloud",
	Long: "export JUMPCLOUD_API_KEY=`vault kv get -format=json secret/jumpcloud | jq -r '.data.data[\"api_key\"]'`" +
		"mfc jc link-users <userGroupID> <systemGroupID> -o update -withRoot" +
		"mfc jc link-users <userGroupID> <systemGroupID> -o update -withRoot",
	Aliases: []string{"jc"},
}

var manageUserSystemSSHCmd = &cobra.Command{
	Use:   "link-users user_group system_group",
	Short: "This command configures additional ssh properties not available from the Jumpcloud UI",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("JUMPCLOUD_API_KEY")
		if apiKey == "" {
			fmt.Print("JUMPCLOUD_API_KEY should be set in the environment. Exiting")
			os.Exit(1)
		}
		userGroupID := args[0]
		systemGroupID := args[1]

		contentType := "application/json"
		accept := "application/json"

		// Instantiate the API client
		client := jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())

		// Set up the API key via context
		auth := context.WithValue(context.TODO(), jcapiv2.ContextAPIKey, jcapiv2.APIKey{
			Key: apiKey,
		})

		groupResp, res, err := client.SystemGroupMembersMembershipApi.GraphSystemGroupMembership(auth, systemGroupID, contentType, accept, nil)
		if err != nil {
			fmt.Printf("Error getting system group %s membership - response = %+v\n", systemGroupID, err, res)
		} else {
			fmt.Printf("Got system group memebership %s: \n", res.Body)
		}

		userResp, res, err := client.UserGroupMembersMembershipApi.GraphUserGroupMembersList(auth, userGroupID, contentType, accept, nil)
		if err != nil {
			fmt.Printf("Error getting user group %user: %user - response = %+v\n", systemGroupID, err, res)
		} else {
			fmt.Printf("Details for adding user to user group %user: \n", res.Body)
		}
		for _, user := range userResp {
			for _, system := range groupResp {
				//setUserSSHPerm(client, auth, user.To.Id, system.Id, contentType, accept)
				//fmt.Printf("%s %v\n", system.Id, system)
				setUserSSHPerm(client, auth, user.To.Id, system.Id, contentType, accept)
			}
			fmt.Print(user)
			//os.Exit(0)
		}
	},
}

func setUserSSHPerm(client *jcapiv2.APIClient, auth context.Context, userID string, systemID string, contentType string, accept string) {
	time.Sleep(time.Millisecond * time.Duration(250))
	optParams := make(map[string]interface{})

	switch operation {
	case "add", "update":
		attr := jcapiv2.SystemGraphManagementReqAttributesSudo{Enabled: withRoot, WithoutPassword: withoutPassSudo}
		attrs := jcapiv2.SystemGraphManagementReqAttributes{Sudo: &attr}
		optParams["body"] = jcapiv2.UserGraphManagementReq{Op: operation, Type_: "system", Id: systemID, Attributes: &attrs}
	case "remove":
		optParams["body"] = jcapiv2.UserGraphManagementReq{Op: operation, Type_: "system", Id: systemID}
	default:
		fmt.Print("requested operation is invalid")
		os.Exit(1)
	}
	res, err := client.UsersApi.GraphUserAssociationsPost(auth, userID, contentType, accept, optParams)
	if err != nil {
		if res.StatusCode == 404 {
			fmt.Printf("User %s not added to System %s. System %s is likely down.\n", userID, systemID, systemID)
			return
		} else if res.StatusCode == 409 {
			fmt.Printf("User %s already added to System %s. Skipping\n", userID, systemID)
			return
		} else {
			fmt.Printf("Requested %s ssh permissions for %s completed for system %s: %s - response = %+v\n", operation, userID, systemID, err, res)
			os.Exit(1)
		}
	}
	fmt.Printf("Requested %s ssh permissions for %s completed for system %s: \n", operation, userID, systemID)
}

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DataSchema is the format of the JSON subsection of ResponseSchema
type DataSchema struct {
	SerialNumber string `json:"serial_number"`
	SignedKey    string `json:"signed_key"`
}

// ResponseSchema is the format of the JSON response we get after a successful request
type ResponseSchema struct {
	RequestID     string     `json:"request_id"`
	LeaseID       string     `json:"lease_id"`
	Renewable     bool       `json:"renewable"`
	LeaseDuration int        `json:"lease_duration"`
	Data          DataSchema `json:"data"`
	WrapInfo      string     `json:"wrap_info"`
	Warnings      string     `json:"warnings"`
	Auth          string     `json:"auth"`
}

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.PersistentFlags().StringVarP(&sshPublicKeyPath, "public-key", "a", filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa.pub"), "Path used to read SSH public key")
	sshCmd.PersistentFlags().StringVarP(&signedPublicKeyPath, "signed-public-key", "b", filepath.Join(os.Getenv("HOME"), ".ssh", "signed-cert.pub"), "Path to write signed certificate")
	sshCmd.PersistentFlags().StringVarP(&userKeyPath, "user-key-path", "u", "ssh-signer/sign/user-key", "Vault endpoint for user key signing")
}

var (
	signedPublicKeyPath string
	sshPublicKeyPath    string
	userKeyPath         string
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Sign client SSH key",
	Run: func(cmd *cobra.Command, args []string) {

		sshPublicKeyBytes, sshPublicKeyReadError := ioutil.ReadFile(sshPublicKeyPath)
		check(sshPublicKeyReadError)
		sshPublicKey := string(sshPublicKeyBytes)
		sshPublicKeyTrimmed := strings.TrimRight(sshPublicKey, "\r\n")

		var urlBuffer bytes.Buffer
		urlBuffer.WriteString("https://vault.missionfocus.com/v1/")
		urlBuffer.WriteString(userKeyPath)
		url := urlBuffer.String()

		client := &http.Client{}
		bodyContents := fmt.Sprintf("{\"public_key\": \"%s\"}", sshPublicKeyTrimmed)

		body := bytes.NewBufferString(bodyContents)

		request, requestError := http.NewRequest(http.MethodPut, url, body)
		check(requestError)

		request.Header.Set("X-Vault-Token", "3TencBRGMsrc48VFIxmonLfS")
		request.Header.Set("Content-Type", "application/json")

		response, responseError := client.Do(request)
		check(responseError)

		responseBody, responseReadError := ioutil.ReadAll(response.Body)
		check(responseReadError)

		var parsedResponse ResponseSchema
		unmarshalError := json.Unmarshal(responseBody, &parsedResponse)
		check(unmarshalError)

		responseData := parsedResponse.Data
		signedKey := responseData.SignedKey

		signedKeyBytes := []byte(signedKey)
		writeSignedPublicKeyError := ioutil.WriteFile(signedPublicKeyPath, signedKeyBytes, 0644)
		check(writeSignedPublicKeyError)

		responseCloseError := response.Body.Close()
		check(responseCloseError)

		fmt.Printf("Signed public key written to: %s\n", signedPublicKeyPath)

	},
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(vaultVPNCmd)
	vaultVPNCmd.AddCommand(vaultVPNInitCmd)
	vaultVPNCmd.AddCommand(vaultVPNRenewCmd)
	vaultVPNCmd.AddCommand(vaultVPNUpCmd)
	vaultVPNCmd.AddCommand(vaultVPNDownCmd)

	vaultVPNCmd.PersistentFlags().StringVar(&vaultVPNGroup, "vault-key", "engineers", "default vpnclient user group")
	vaultVPNCmd.PersistentFlags().StringVarP(&vaultVPNPath, "path", "p", filepath.Join(homeDir(), "/.config/mf/vpn"), "path to client vpn config")
}

const vpnBasePath = "secret/data/vpn"

var (
	vaultVPNGroup string
	vaultVPNPath  string
)

const vaultVPNExample = `
  mfc vault vpn init                   # Setup vpn configuration
  mfc vault vpn renew <ldap username>  # Renew vpn lease for ldap user
  mfc vault vpn up                     # Connect to vpn
  mfc vault vpn down                   # Disconnect from vpn`

var vaultVPNCmd = &cobra.Command{
	Use:     "vpn",
	Short:   "Setup client vpn config",
	Example: vaultVPNExample,
}

var vaultVPNInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Refreshes base vpn conf for client vpn",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		confFile := path.Join(vaultVPNPath, vaultVPNGroup+".ovpn")
		vaultConfFile := path.Join(vpnBasePath, "ovpn")
		check(err)

		secret, err := client.Logical().Read(path.Join(vaultConfFile))
		check(err)
		data := secret.Data["data"].(map[string]interface{})
		confData := data[vaultVPNGroup].(string)

		f, err := openOrCreate(confFile, 0755, 0755)
		check(err)
		defer f.Close()
		vpnConfBytes := []byte(confData)
		f.Write(vpnConfBytes)

		switch platform := runtime.GOOS; platform {
		case "darwin":
			script := fmt.Sprintf("tell application \"/Applications/Tunnelblick.app\"\nconnect \"%s\"\nend tell\n", vaultVPNGroup)
			err := ioutil.WriteFile(path.Join(vaultVPNPath, vaultVPNGroup+"_up.scpt"), []byte(script), 0755)
			check(err)
			script = fmt.Sprintf("tell application \"/Applications/Tunnelblick.app\"\ndisconnect \"%s\"\nend tell\n", vaultVPNGroup)
			err = ioutil.WriteFile(path.Join(vaultVPNPath, vaultVPNGroup+"_down.scpt"), []byte(script), 0755)
			check(err)

			fmt.Println("Make sure tunnelblick is present with `brew cask install tunnelblick`")
			fmt.Println("Run `mfc vault vpn renew <ldap username>`")
		case "linux":
			cmd := exec.Command("/usr/bin/nmcli", "connection", "import", "--temporary", "type", "openvpn", "file", confFile)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s %v\n", string(out), err)
			}
		default:
			fmt.Printf("unsupported platform %s\n", platform)
		}
	},
}

var (
	vaultVPNRenewTTL string
)

const vaultVPNRenewExample = `
  mfc vault vpn renew <ldap username>  # Renew vpn lease for ldap user`

var vaultVPNRenewCmd = &cobra.Command{
	Use:     "renew <ldap_username>",
	Short:   "Renew client vpn",
	Example: vaultVPNRenewExample,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		secret, err := v.PKIIssue(&vault.PKIIssueOptions{
			RoleName:   vault.DefaultPKIEngineRole,
			CommonName: args[0] + ".missionfocus.com",
			TTL:        vaultVPNRenewTTL,
			Format:     vaultPKIIssueFormat,
		})
		check(err)

		chain, err := os.OpenFile(filepath.Join(vaultVPNPath, "certificate.pem"), os.O_CREATE|os.O_WRONLY, 0700)
		check(err)
		defer chain.Close()
		check(secret.WriteCertificate(chain))
		priv, err := os.OpenFile(filepath.Join(vaultVPNPath, "privkey.pem"), os.O_CREATE|os.O_WRONLY, 0700)
		check(err)
		defer priv.Close()
		check(secret.WritePrivateKey(priv))

		confFile := path.Join(vaultVPNPath, vaultVPNGroup+".ovpn")
		oCmd := exec.Command("/usr/bin/open", confFile)
		out, err := oCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("%s %v\n", string(out), err)
		}

		oCmd = exec.Command("/usr/bin/defaults", "write", "net.tunnelblick.tunnelblick", vaultVPNGroup+"-resetPrimaryInterfaceAfterDisconnect", "-bool", "true")
		out, err = oCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("%s %v\n", string(out), err)
		}
	},
}

var vaultVPNUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Connect to the Client VPN",
	Run: func(cmd *cobra.Command, args []string) {
		switch platform := runtime.GOOS; platform {
		case "darwin":
			cmd := exec.Command("/usr/bin/osascript", path.Join(vaultVPNPath, vaultVPNGroup+"_up.scpt"))
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s %v\n", string(out), err)
			}
		case "linux":
			cmd := exec.Command("/usr/bin/nmcli", "con", "up", "id", vaultVPNGroup)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s %v\n", string(out), err)
			}
		default:
			fmt.Printf("unsupported platform %s\n", platform)
		}
	},
}

var vaultVPNDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Disconnect from the Client VPN",
	Run: func(cmd *cobra.Command, args []string) {
		switch platform := runtime.GOOS; platform {
		case "darwin":
			cmd := exec.Command("/usr/bin/osascript", path.Join(vaultVPNPath, vaultVPNGroup+"_down.scpt"))
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s %v\n", string(out), err)
			}
		case "linux":
			cmd := exec.Command("/usr/bin/nmcli", "con", "down", "id", vaultVPNGroup)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s %v\n", string(out), err)
			}
		default:
			fmt.Printf("unsupported platform %s\n", platform)
		}
	},
}

func openOrCreate(path string, dirPerm, filePerm os.FileMode) (*os.File, error) {
	if exists := fileExists(path); exists {
		return os.OpenFile(path, os.O_RDWR, filePerm)
	}

	dir, _ := filepath.Split(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return nil, err
	}

	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, filePerm)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else if err != nil {
		fmt.Print(err)
		return false
	} else {
		return true
	}
}

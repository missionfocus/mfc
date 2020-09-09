# mfc

Mission Focus Control CLI. View detailed documentation [here](docs/generated/mfc.md).

# About
mfc is the definative tool for encapsulating interactions for authentication, and secrets management for Mission Focus infrastructure. Written in Golang with the Cobra CLI package, mfc can be compiled, and run on all major operating systems. The secrets engine that mfc interacts with is Hashicorp Vault. Vault provides a unified secrets management system with API access.
 
 The typical workflow involves the user obtaining a vault token through thier LDAP credentials. Frequently used actions, and credentialing requirements are available within thoughtfully nested sub commands. Some example usage senarios include assuming roles for the Mission Focus AWS accounts, writing/getting secrets, and encrypting/decrypting messages.
 
 # Usage
 mfc has built in help, and guidance for commands, and subcommands that are printed with `-h`, or `--help` passed in to any command and subcommand. 
 
 It is recommended to install mfc specific Alt-Tab completions for your $SHELL. See available guide by running `mfc config completions -h`

# Infrastructure
### A high level overview of how mfc, Vault ties together
* Secrets engines, and identity plugins for services in use (AWS Assume role, KV storage) are enabled on the Vault server. 

* The Vault client, in our case mfc, authenticates through a LDAP secrets engine to obtain a vault-token that is valid for one work day.
 
* With the vault-token, the client is able to run requests based on policies, and roles loaded in the Vault server.   

## Updates
```
mfc update
```

## Adhoc Installs
### mfc is provded by dev-staging. For adhoc installation, use the commands below. 
### OSX
```
URL=$(curl http://public.missionfocus.com/mfc/manifest.yaml | grep darwin | grep -o 'http://.*')
echo $URL 
sudo mkdir /usr/local/bin
sudo curl $URL -o /usr/local/bin/mfc
sudo chmod +x /usr/local/bin/mfc
```
### Linux
```
URL=$(curl http://public.missionfocus.com/mfc/manifest.yaml | grep linux | grep -o 'http://.*')
echo $URL
sudo mkdir /usr/local/bin 
sudo curl $URL -o /usr/local/bin/mfc
sudo chmod +x /usr/local/bin/mfc
```



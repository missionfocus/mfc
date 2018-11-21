package vault

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func getBaseAddr(account string) string {
	switch account {
	case "govcloud":
		return "amazonaws-us-gov.com"
	default:
		return "aws.amazon.com"
	}
}

type STSSecret struct {
	AccessKeyID     string `json:"sessionId"`
	SecretAccessKey string `json:"sessionKey"`
	SecurityToken   string `json:"sessionToken"`
}

// Matches [<profile name>]
var re = regexp.MustCompile(`^\[(.+)]`)

// Adds or updates the STS secret as an AWS profile to the specified credentials file.
func (s *STSSecret) ToProfile(path string, name string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	nLines := len(lines)
	if lines[nLines-1] != "" {
		lines = append(lines, "")
		nLines++
	}

	blocks := make([]string, 0)
	for i := 0; i < nLines; i++ {
		match := re.FindStringSubmatch(lines[i])
		if len(match) == 2 && match[1] != name {
			for j := i + 1; j < nLines; j++ {
				if lines[j] == "" || re.MatchString(lines[j]) {
					blocks = append(blocks, strings.Join(lines[i:j], "\n"))
					i = j - 1
					break
				}
			}
		}
	}

	blocks = append(blocks, "["+name+"]\n"+s.toCredentials())
	return ioutil.WriteFile(path, []byte(strings.Join(blocks, "\n\n")), 0600)
}

func (s *STSSecret) toCredentials() string {
	var sb strings.Builder

	sb.WriteString("aws_access_key_id = ")
	sb.WriteString(s.AccessKeyID)
	sb.WriteRune('\n')

	sb.WriteString("aws_secret_access_key = ")
	sb.WriteString(s.SecretAccessKey)
	sb.WriteRune('\n')

	sb.WriteString("aws_security_token = ")
	sb.WriteString(s.SecurityToken)
	sb.WriteRune('\n')

	return sb.String()
}

type GetSigninTokenResponse struct {
	SigninToken string
}

func (s *STSSecret) GenerateLoginUrl(account string) (*url.URL, error) {
	marshaled, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	baseUrl := getBaseAddr(account)

	tokenQ := url.Values{}
	tokenQ.Set("Action", "getSigninToken")
	tokenQ.Set("Session", string(marshaled))
	client := &http.Client{Timeout: time.Second * 10}
	tokenUrl := url.URL{
		Scheme:   "https",
		Host:     "signin." + baseUrl,
		Path:     "federation",
		RawQuery: tokenQ.Encode(),
	}
	req, _ := http.NewRequest("GET", tokenUrl.String(), nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var decoded GetSigninTokenResponse
	if err := json.Unmarshal(body, &decoded); err != nil {
		return nil, err
	}

	loginQ := url.Values{}
	loginQ.Set("Action", "login")
	loginQ.Set("Issuer", "https://www.missionfocus.com/")
	loginQ.Set("Destination", "https://console."+baseUrl)
	loginQ.Set("SigninToken", decoded.SigninToken)

	return &url.URL{
		Scheme:   "https",
		Host:     "signin." + baseUrl,
		Path:     "federation",
		RawQuery: loginQ.Encode(),
	}, nil
}

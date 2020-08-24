package tmetric

import (
	"errors"
	"strings"
	"time"

	tmetric "git.missionfocus.com/ours/code/libraries/go/tmetric/client"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/client/accounts"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/client/time_entries"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
)

const AccountID = 105432

type taskPerformanceRecord struct {
	description string
	url         string
	pointsSpent float64
	weight      int
	score       float64
}

func (r taskPerformanceRecord) More(other taskPerformanceRecord) bool {
	more := false
	if r.score > other.score {
		more = true
	} else if r.score == other.score {
		if r.weight > other.weight {
			more = true
		} else if r.weight == other.weight {
			if (r.url != "") && (other.url == "") {
				more = true
			}
		}
	}
	return more
}

func GetTmetricAuth(vaultClient vault.Vault) (error, runtime.ClientAuthInfoWriter) {
	secret, err := vaultClient.KVUserGet("tmetric")
	if err != nil {
		return err, nil
	}
	if secret == nil {
		return errors.New("could not retrieve TMetric token. You may need to set it with `mfc tmetric set-token`"), nil
	}

	tok := secret.Data["data"].(map[string]interface{})["token"].(string)
	auth := httptransport.BearerToken(tok)

	return nil, auth
}

func GetAllTMetricMembers(vaultClient vault.Vault) (*models.AccountScope, error) {
	_, auth := GetTmetricAuth(vaultClient)
	params := accounts.NewAccountsGetAccountScopeParams().WithAccountID(AccountID)

	resp, err := tmetric.Default.Accounts.AccountsGetAccountScope(params, auth)
	if err != nil {
		return nil, err
	}

	scope := resp.Payload
	return scope, nil
}

func GetTimeEntriesWithParams(vaultClient vault.Vault, params *time_entries.TimeEntriesGetTimeEntriesParams) ([]*models.TimeEntry, error) {
	_, auth := GetTmetricAuth(vaultClient)
	resp, err := tmetric.Default.TimeEntries.TimeEntriesGetTimeEntries(params, auth)
	if err != nil {
		panic(err)
	}
	timeEntries := resp.Payload
	return timeEntries, nil
}

const (
	glTimeFormat = "2006-01-02"
)

//GetTimeParameters is used to alter the format [date] | [date] into a comparable format
func GetTimeParameters(str string) []time.Time {
	dates := make([]time.Time, 0)

	if len(str) == 0 {
		date := "1999-12-31"
		t, _ := time.Parse(glTimeFormat, date)
		dates = append(dates, t)

		currentTime := time.Now()
		currentTime.Format(glTimeFormat)
		dates = append(dates, currentTime)
	}

	splitDateStrings := strings.Split(str, "|")
	for _, d := range splitDateStrings {
		strToDate := strings.Replace(d, " ", "", -1)
		t, _ := time.Parse(glTimeFormat, strToDate)
		dates = append(dates, t)
	}
	return dates
}

package tmetric

import (
	"errors"
	tmetric "git.missionfocus.com/ours/code/libraries/go/tmetric/client"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/client/accounts"
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

func GetTmetricAuth(vaultClient vault.Vault) (error, runtime.ClientAuthInfoWriter ) {
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



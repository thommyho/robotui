package service

import (
	"net/url"

	"github.com/robotuimyhorobotuiotui/vehicle/vag"
	"github.com/robotuimyhorobotuiotui/vehicle/vag/mbb"
	"github.com/robotuimyhorobotuiotui/vehicle/vag/vwidentity"
	"github.com/thommyho/robotui/util"
)

// MbbTokenSource creates a refreshing token source for use with the MBB api.
// Once the MBB token expires, it is recreated from the token exchanger (either TokenRefreshService or IDK)
func MbbTokenSource(log *util.Logger, tox vag.TokenExchanger, clientID string, q url.Values, user, password string) (vag.TokenSource, error) {
	q, err := vwidentity.Login(log, q, user, password)
	if err != nil {
		return nil, err
	}

	token, err := tox.Exchange(q)
	if err != nil {
		return nil, err
	}

	trs := tox.TokenSource(token)
	mbb := mbb.New(log, clientID)

	mts := vag.MetaTokenSource(func() (*vag.Token, error) {
		// get TRS token from refreshing TRS token source
		itoken, err := trs.TokenEx()
		if err != nil {
			return nil, err
		}

		// exchange TRS id_token for MBB token
		mtoken, err := mbb.Exchange(url.Values{"id_token": {itoken.IDToken}})
		if err != nil {
			return nil, err
		}

		return mtoken, err

		// produce tokens from refresh MBB token source
	}, mbb.TokenSource)

	return mts, nil
}

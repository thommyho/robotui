package jlr

import (
	"fmt"
	"net/http"
	"time"

	"github.com/thommyho/robotui/util"
	"github.com/thommyho/robotui/util/oauth"
	"github.com/thommyho/robotui/util/request"
	"golang.org/x/oauth2"
)

// https://github.com/ardevd/jlrpy

const IFAS_BASE_URL = "https://ifas.prod-row.jlrmotor.com/ifas/jlr"

type Identity struct {
	*request.Helper
	user, password, device string
	oauth2.TokenSource
}

// NewIdentity creates Fiat identity
func NewIdentity(log *util.Logger, user, password, device string) *Identity {
	return &Identity{
		Helper:   request.NewHelper(log),
		user:     user,
		password: password,
		device:   device,
	}
}

// Login authenticates with given payload
func (v *Identity) login(data map[string]string) (Token, error) {
	uri := fmt.Sprintf("%s/tokens", IFAS_BASE_URL)
	req, err := request.New(http.MethodPost, uri, request.MarshalJSON(data), map[string]string{
		"Authorization": "Basic YXM6YXNwYXNz",
		"Content-type":  request.JSONContent,
		"X-Device-Id":   v.device,
	})

	var token Token
	if err == nil {
		err = v.DoJSON(req, &token)
		token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}

	return token, err
}

// Login authenticates with username/password
func (v *Identity) Login() (Token, error) {
	data := map[string]string{
		"grant_type": "password",
		"username":   v.user,
		"password":   v.password,
	}

	token, err := v.login(data)
	if err == nil {
		v.TokenSource = oauth.RefreshTokenSource(&token.Token, v)
	}

	return token, err
}

// RefreshToken implements oauth.TokenRefresher
func (v *Identity) RefreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	data := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token.RefreshToken,
	}

	res, err := v.login(data)

	return &res.Token, err
}

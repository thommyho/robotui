package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/skratchdot/open-golang/open"
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/util"
	"github.com/thommyho/robotui/vehicle"
	"github.com/thommyho/robotui/vehicle/tronity"
	"golang.org/x/oauth2"
)

func tokenExchangeHandler(oc *oauth2.Config, state string, resC chan *oauth2.Token) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if remote := r.URL.Query().Get("state"); state != remote {
			w.WriteHeader(http.StatusBadRequest)
			resC <- nil
			return
		}

		code := r.URL.Query().Get("code")

		ctx := context.Background()
		token, err := oc.Exchange(ctx, code,
			oauth2.SetAuthURLParam("grant_type", "code"), // app
		)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			resC <- nil
			return
		}

		fmt.Fprintln(w, "Token received, see console")
		resC <- token
	}
}

func tronityAuthorize(addr string, oc *oauth2.Config) (*oauth2.Token, error) {
	state := lo.RandomString(16, lo.AlphanumericCharset)

	uri := oc.AuthCodeURL(state, oauth2.AccessTypeOffline)
	uri = strings.ReplaceAll(uri, "scope=", "scopes=")

	if err := open.Start(uri); err != nil {
		return nil, err
	}

	resC := make(chan *oauth2.Token)
	handler := tokenExchangeHandler(oc, state, resC)
	defer close(resC)

	// handle request
	mux := &http.ServeMux{}
	mux.HandleFunc("/auth/tronity", handler)

	wg := new(sync.WaitGroup)
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// start server
	wg.Add(1)
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.FATAL.Fatal(err)
		}
		wg.Done()
	}()

	// close on exit
	defer func() {
		_ = s.Close()
		wg.Wait()
		close(resC)
	}()

	t := time.NewTimer(time.Minute)

	select {
	case <-t.C:
		return nil, api.ErrTimeout

	case token := <-resC:
		if token == nil {
			return nil, errors.New("token not received")
		}

		return token, nil
	}
}

func tronityToken(conf config, vehicleConf qualifiedConfig) (*oauth2.Token, error) {
	var cc struct {
		Credentials vehicle.ClientCredentials
		RedirectURI string
		Other       map[string]interface{} `mapstructure:",remain"`
	}

	if err := util.DecodeOther(vehicleConf.Other, &cc); err != nil {
		return nil, err
	}

	if err := cc.Credentials.Error(); err != nil {
		return nil, err
	}

	oc, err := tronity.OAuth2Config(cc.Credentials.ID, cc.Credentials.Secret)
	if err != nil {
		return nil, err
	}

	if oc.RedirectURL = cc.RedirectURI; oc.RedirectURL == "" {
		oc.RedirectURL = fmt.Sprintf("%s/auth/tronity", conf.Network.URI())
	}

	return tronityAuthorize(conf.Network.HostPort(), oc)
}

package auth

import (
	"context"
	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"log"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://dev-c-559zpw.auth0.com/api/v2/")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     "u2ZbAXZKz4kM0MM27o6R7mmYQ8pteoFw",
		ClientSecret: "MSCR94_grYm42kems9ng7jPvAVOlvpHLV-c7xg4UXm190LBbgUtgwkgBujgPHzHq",
		Endpoint:     oauth2.Endpoint{},
		RedirectURL:  "http://localhost:3000/callback",
		Scopes:       []string{oidc.ScopeOpenID, "home"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		ctx:      ctx,
	}, nil
}

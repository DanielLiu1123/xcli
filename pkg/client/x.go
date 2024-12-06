package client

import (
	"github.com/DanielLiu1123/xcli/pkg/model"
	"github.com/dghubble/oauth1"
	"github.com/g8rswimmer/go-twitter/v2"
	"net/http"
)

func NewXClient(auth *model.Auth) *twitter.Client {
	// https://github.com/g8rswimmer/go-twitter/issues/149#issuecomment-1634979382
	config := oauth1.NewConfig(auth.APIKey, auth.APISecret)
	httpClient := config.Client(oauth1.NoContext, &oauth1.Token{
		Token:       auth.AccessToken,
		TokenSecret: auth.AccessSecret,
	})
	return &twitter.Client{
		Authorizer: &authorize{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}
}

// dummy
type authorize struct {
}

func (a authorize) Add(_ *http.Request) {
}

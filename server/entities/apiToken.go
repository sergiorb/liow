package entities

import (
  "net/http"
  "github.com/sergiorb/liow/server/config"
)

var conf = config.Load()

type ApiToken struct {
  Token string `json: "token"`
}

func NewApiToken(r *http.Request) *ApiToken {

  token := r.Header.Get(conf.Api.ApiTokenName)

  return &ApiToken{token}
}

func (t *ApiToken) IsValid() (bool) {

  var isValid bool

  if t.Token == "" {

    isValid = false

  } else {

    isValid = true
  }

  return isValid
}

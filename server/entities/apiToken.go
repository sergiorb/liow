package entities

import (
  "net/http"
  "github.com/sergiorb/liow/server/config"
  "github.com/op/go-logging"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")

type ApiToken struct {
  Token string `json: "token"`
}

func NewApiToken(r *http.Request) *ApiToken {

  token := r.Header.Get(conf.Api.ApiTokenName)

  log.Warning(token)

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

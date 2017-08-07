package models

import (
  "github.com/sergiorb/liow/server/config"
  "github.com/op/go-logging"
)

const (
  TOKEN_COLLECTION_NAME = "token"
  USER_COLLECTION_NAME = "user"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")

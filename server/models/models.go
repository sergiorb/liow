package models

import (
  "github.com/sergiorb/liow/server/config"
  "github.com/op/go-logging"
)

const (
  TOKEN_COLLECTION_NAME = "token"
  REGISTER_COLLECTION_NAME = "register"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")

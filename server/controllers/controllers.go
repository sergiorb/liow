package controllers

import (
  "github.com/op/go-logging"
  "github.com/sergiorb/liow/server/config"
)

const (
	SESSION = "session"
  SCREEN  = "screen"
	LOGIN   = "login"
	LOGOUT  = "logout"
  LOCK    = "lock"
  UNLOCK  = "unlock"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")

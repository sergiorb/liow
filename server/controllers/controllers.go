package controllers

import (
  "github.com/op/go-logging"
  "github.com/sergiorb/liow/server/config"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")

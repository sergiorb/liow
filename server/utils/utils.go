package utils

import (
  "github.com/sergiorb/liow/server/config"
  "github.com/op/go-logging"
)

var log = logging.MustGetLogger("log-in-out-watcher server")
var conf = config.Load()

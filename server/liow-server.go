package main

import (
  "fmt"
  "github.com/op/go-logging"
  "net/http"
  "github.com/sergiorb/liow/server/config"
  "github.com/sergiorb/liow/server/router"
  _ "github.com/sergiorb/liow/server/logger"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")

func main()  {

  log.Info(fmt.Sprintf("Starting log-in-out-watcher server v%v on %v:%v",
    conf.Api.Version,
    conf.Server.Host,
    conf.Server.Port))

  log.Info("Listening...")

  http.ListenAndServe(fmt.Sprintf("%v:%v",
    conf.Server.Host,
    conf.Server.Port),
    router.Router)
}

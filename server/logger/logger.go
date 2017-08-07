package logger

import (
  "os"
  "github.com/op/go-logging"
  "github.com/sergiorb/liow/server/config"
)

var conf = config.Load()
var format = logging.MustStringFormatter(
	`%{color}%{time:2006/01/02 - 15:04:05.000} %{longfunc} - %{level:.4s} ▶ %{color:reset} %{message}`,
)

func init()  {

  logOutput := logging.NewLogBackend(os.Stderr, "", 0)
  backendFormatter := logging.NewBackendFormatter(logOutput, format)
	stderrorLog := logging.AddModuleLevel(backendFormatter)

  if conf.Logger.Level == "debug" {

    stderrorLog.SetLevel(logging.DEBUG, "")

  } else if conf.Logger.Level == "info" {

    stderrorLog.SetLevel(logging.INFO, "")

  } else if conf.Logger.Level == "notice" {

    stderrorLog.SetLevel(logging.NOTICE, "")

  } else if conf.Logger.Level == "warning" {

    stderrorLog.SetLevel(logging.WARNING, "")

  } else if conf.Logger.Level == "error" {

    stderrorLog.SetLevel(logging.ERROR, "")

  } else if conf.Logger.Level == "critical" {

    stderrorLog.SetLevel(logging.CRITICAL, "")

  } else {

	   stderrorLog.SetLevel(logging.INFO, "")
   }

	logging.SetBackend(stderrorLog)
}

package models

import (
  "github.com/sergiorb/liow/server/config"
  "github.com/op/go-logging"
  "gopkg.in/mgo.v2"
)

const (
  TOKEN_COLLECTION_NAME = "token"
  REGISTER_COLLECTION_NAME = "register"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")


func GetMongoSession() *mgo.Session {

	session, err := mgo.Dial(conf.Database.GetUrl())

	if err != nil { panic("Can't dial database") }

	return session
}

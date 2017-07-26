package router

import (
  "github.com/gorilla/mux"
  "github.com/sergiorb/liow/server/controllers"
  "github.com/sergiorb/liow/server/config"
  "gopkg.in/mgo.v2"
  "net/http"
  "fmt"
  "github.com/op/go-logging"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")
var Router = mux.NewRouter()

func getMongoSession() *mgo.Session {

	session, err := mgo.Dial(conf.Database.GetUrl())

	if err != nil { panic("Can't load config file.") }

	return session
}

func init() {

  sessionController := controllers.NewSessionController(getMongoSession())
  screenController := controllers.NewScreenController(getMongoSession())

  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/session/login"), http.HandlerFunc(sessionController.Login)).Methods("POST")
  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/session/logout"), http.HandlerFunc(sessionController.Logout)).Methods("POST")
  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/screen/lock"), http.HandlerFunc(screenController.Lock)).Methods("POST")
  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/screen/unlock"), http.HandlerFunc(screenController.Unlock)).Methods("POST")
}

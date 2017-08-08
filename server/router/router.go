package router

import (
  "github.com/gorilla/mux"
  "github.com/sergiorb/liow/server/controllers"
  "github.com/sergiorb/liow/server/config"
  "gopkg.in/mgo.v2"
  "net/http"
  "fmt"
  "github.com/op/go-logging"
  //"strings"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")
var Router = mux.NewRouter()

func getMongoSession() *mgo.Session {

	session, err := mgo.Dial(conf.Database.GetUrl())

	if err != nil { panic("Can't dial database") }

	return session
}

func init() {

  //userController := controllers.NewUserController(getMongoSession())
  //tokenController := controllers.NewTokenController(getMongoSession())
  registerController := controllers.NewRegisterController(getMongoSession())

  /*Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/user/{id}"),
    http.HandlerFunc(userController.Read)).Methods("GET")*/

  /*Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/token/{id}"),
    http.HandlerFunc(tokenController.Read)).Methods("GET")*/

  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/register"),
    checkAPIToken(http.HandlerFunc(registerController.Create))).Methods("POST")
}

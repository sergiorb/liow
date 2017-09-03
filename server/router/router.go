package router

import (
  "github.com/gorilla/mux"
  "github.com/sergiorb/liow/server/controllers"
  "github.com/sergiorb/liow/server/entities/authorization"
  "github.com/sergiorb/liow/server/config"
  "github.com/sergiorb/liow/server/models"
  "net/http"
  "fmt"
  "github.com/op/go-logging"
)

var conf = config.Load()
var log = logging.MustGetLogger("log-in-out-watcher server")
var Router = mux.NewRouter()

func init() {

  liowController := controllers.NewLiowController()

  screenController := controllers.NewScreenController(models.GetMongoSession())
  sessionController := controllers.NewSessionController(models.GetMongoSession())

  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/app/ping"),
    http.HandlerFunc(liowController.Ping)).Methods("GET")

  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/screen/lock"),
    withAuth(withRole(authorization.ROLE_API_CLIENT, http.HandlerFunc(screenController.Lock)))).Methods("POST")

  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/screen/unlock"),
    withAuth(withRole(authorization.ROLE_API_CLIENT, http.HandlerFunc(screenController.Unlock)))).Methods("POST")

  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/session/login"),
    withAuth(withRole(authorization.ROLE_API_CLIENT, http.HandlerFunc(sessionController.Login)))).Methods("POST")

  Router.Handle(fmt.Sprintf("%v%v", conf.GetFullApiPrefix(), "/session/logout"),
    withAuth(withRole(authorization.ROLE_API_CLIENT, http.HandlerFunc(sessionController.Logout)))).Methods("POST")
}

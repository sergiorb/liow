package controllers

import (
	// "github.com/sergiorb/liow/server/models"
  "github.com/sergiorb/liow/server/entities"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/op/go-logging"
	"fmt"
)

var log = logging.MustGetLogger("log-in-out-watcher server")

type ScreenController struct {
	session *mgo.Session
}

func NewScreenController(s *mgo.Session) *ScreenController {
	return &ScreenController{s}
}

func (sc ScreenController) Lock(w http.ResponseWriter, r *http.Request) {

  var ap *entities.ApiResponse

	apiToken := entities.NewApiToken(r)

	if !apiToken.IsValid() {

    ap = &entities.ApiResponse{Status:"no valid token"}
    w.WriteHeader(http.StatusBadRequest)

		log.Warning(fmt.Sprintf("No valid token: %v", apiToken.Token))

	} else {

    ap = &entities.ApiResponse{Status:"OK"}
    w.WriteHeader(http.StatusOK)
  }

  payload, _ := json.Marshal(&ap)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

func (sc ScreenController) Unlock(w http.ResponseWriter, r *http.Request) {

  var ap *entities.ApiResponse

	apiToken := entities.NewApiToken(r)

	if !apiToken.IsValid() {

    ap = &entities.ApiResponse{Status:"no valid token"}
    w.WriteHeader(http.StatusBadRequest)

		log.Warning(fmt.Sprintf("No valid token: %v", apiToken.Token))

	} else {

    ap = &entities.ApiResponse{Status:"OK"}
    w.WriteHeader(http.StatusOK)
  }

  payload, _ := json.Marshal(&ap)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

package controllers

import (
	"github.com/sergiorb/liow/server/models"
	"github.com/sergiorb/liow/server/entities/api"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"net/http"
)

type ScreenController struct {
	session *mgo.Session
}

func NewScreenController(s *mgo.Session) *ScreenController {
	return &ScreenController{s}
}

func (sc ScreenController) Lock(w http.ResponseWriter, r *http.Request) {

  var createResponse *api.CreationResponse

	register :=  models.Register{
		Token: r.Header.Get(conf.Api.ApiTokenName),
		Data: map[string]string{"event": SCREEN, "action": LOCK},
	}

  registerDao := models.NewRegisterDao(sc.session)
	defer registerDao.CloseSession()

	err := registerDao.Create(&register)

	if err != nil {

		createResponse = &api.CreationResponse {
			Objects:  []interface{}{register},
			Message:  "Error saving register",
			Errors:   err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

	} else {

		createResponse = &api.CreationResponse{
			Objects:  []interface{}{register},
		}
	}

  payload, _ := json.Marshal(&createResponse)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

func (sc ScreenController) Unlock(w http.ResponseWriter, r *http.Request) {

  var createResponse *api.CreationResponse

	register :=  models.Register{
		Token: r.Header.Get(conf.Api.ApiTokenName),
		Data: map[string]string{"event": SCREEN, "action": UNLOCK},
	}

  registerDao := models.NewRegisterDao(sc.session)
	defer registerDao.CloseSession()

	err := registerDao.Create(&register)

	if err != nil {

		createResponse = &api.CreationResponse {
			Objects:  []interface{}{register},
			Message:  "Error saving register",
			Errors:   err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

	} else {

		createResponse = &api.CreationResponse{
			Objects:  []interface{}{register},
		}
	}

  payload, _ := json.Marshal(&createResponse)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

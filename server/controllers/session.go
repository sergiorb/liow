package controllers

import (
	"github.com/sergiorb/liow/server/models"
	"github.com/sergiorb/liow/server/entities/api"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"net/http"
)

type SessionController struct {
	session *mgo.Session
}

func NewSessionController(s *mgo.Session) *SessionController {
	return &SessionController{s}
}

func (sc SessionController) Login(w http.ResponseWriter, r *http.Request) {

	var createResponse *api.CreationResponse

	register :=  models.Register{
		Token: r.Header.Get(conf.Api.ApiTokenName),
		Data: map[string]string{"event": SESSION, "action": LOGIN},
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

func (sc SessionController) Logout(w http.ResponseWriter, r *http.Request) {

	var createResponse *api.CreationResponse

	register :=  models.Register{
		Token: r.Header.Get(conf.Api.ApiTokenName),
		Data: map[string]string{"event": SESSION, "action": LOGOUT},
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

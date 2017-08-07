package controllers

import (
	"github.com/gorilla/mux"
	"github.com/sergiorb/liow/server/models"
  "github.com/sergiorb/liow/server/entities/api"
	"encoding/json"
	"gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
	"net/http"
)

type RegisterController struct {
	session *mgo.Session
}

func NewRegisterController(s *mgo.Session) *RegisterController {
	return &RegisterController{s}
}

func (rc RegisterController) Read(w http.ResponseWriter, r *http.Request) {

	var readResponse *api.ReadResponse
  vars := mux.Vars(r)

	id := vars["id"]

	if !bson.IsObjectIdHex(id) {

		readResponse = &api.ReadResponse{Message:"No valid id"}
		w.WriteHeader(http.StatusBadRequest)

	} else {

		registerDao := models.NewRegisterDao(rc.session)
	  defer registerDao.CloseSession()

		register, err := registerDao.Read(id)

		if err != nil {

			readResponse = &api.ReadResponse{Message:"Register not found"}
			w.WriteHeader(http.StatusNotFound)

		} else {

			readResponse = &api.ReadResponse{
        Objects: []interface{}{register},
      }

			w.WriteHeader(http.StatusOK)
		}
	}

	payload, _ := json.Marshal(&readResponse)

	w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

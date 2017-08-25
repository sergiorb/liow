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

type TokenController struct {
	session *mgo.Session
}

func NewTokenController(s *mgo.Session) *TokenController {
	return &TokenController{s}
}

func (tc TokenController) Read(w http.ResponseWriter, r *http.Request) {

	var readResponse *api.ReadResponse
  vars := mux.Vars(r)

	id := vars["id"]

	if !bson.IsObjectIdHex(id) {

		readResponse = &api.ReadResponse{Message:"No valid id"}
		w.WriteHeader(http.StatusBadRequest)

	} else {

		tokenDao := models.NewTokenDao(tc.session)
	  defer tokenDao.CloseSession()

		token, err := tokenDao.Read(id)

		if err != nil {

			readResponse = &api.ReadResponse{Message:"Token not found"}
			w.WriteHeader(http.StatusNotFound)

		} else {

			readResponse = &api.ReadResponse{
        Objects: []interface{}{token},
      }

			w.WriteHeader(http.StatusOK)
		}
	}

	payload, _ := json.Marshal(&readResponse)

	w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

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

func (rc RegisterController) Create(w http.ResponseWriter, r *http.Request) {

	var createResponse *api.CreationResponse

  var register models.Register

  err := json.NewDecoder(r.Body).Decode(&register)
  customErr := register.GetErrors()

	if err != nil {

		createResponse = &api.CreationResponse{
			Objects:  []interface{}{register},
      Message:  "No valid register",
      Errors:   err.Error(),
    }

		w.WriteHeader(http.StatusBadRequest)

  } else if customErr != nil {

    createResponse = &api.CreationResponse{
      Message:  "No valid register",
      Errors:   customErr,
    }

    w.WriteHeader(http.StatusBadRequest)

	} else {

		tokenDao := models.NewTokenDao(rc.session)
		defer tokenDao.CloseSession()

		token, err := tokenDao.GetByToken(r.Header.Get(conf.Api.ApiTokenName))

		if err != nil {

			createResponse = &api.CreationResponse{
        Objects:  []interface{}{register},
        Message:  "No valid token",
        Errors:   err.Error(),
      }

      w.WriteHeader(http.StatusUnauthorized)

		} else {

			registerDao := models.NewRegisterDao(rc.session)
		  defer registerDao.CloseSession()

			register.Token = token.Token

			err := registerDao.Create(&register)

	    if err != nil {

	      createResponse = &api.CreationResponse{
	        Objects:  []interface{}{register},
	        Message:  "No valid register",
	        Errors:   err.Error(),
	      }

	      w.WriteHeader(http.StatusBadRequest)

	    } else {

	      createResponse = &api.CreationResponse{
	        Objects:  []interface{}{register},
	      }

	      w.WriteHeader(http.StatusOK)
	    }
		}
	}

	payload, _ := json.Marshal(&createResponse)

	w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
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

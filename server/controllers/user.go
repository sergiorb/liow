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

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) Read(w http.ResponseWriter, r *http.Request) {

	var readResponse *api.ReadResponse
  vars := mux.Vars(r)

	id := vars["id"]

	if !bson.IsObjectIdHex(id) {

		readResponse = &api.ReadResponse{Message:"No valid id"}
		w.WriteHeader(http.StatusBadRequest)

	} else {

		userDao := models.NewUserDao(uc.session)
	  defer userDao.CloseSession()

		user, err := userDao.Read(id)

		if err != nil {

			readResponse = &api.ReadResponse{Message:"User not found"}
			w.WriteHeader(http.StatusNotFound)

		} else {

			readResponse = &api.ReadResponse{
				Objects: []interface{}{user},
			}

			w.WriteHeader(http.StatusOK)
		}
	}

	payload, _ := json.Marshal(&readResponse)

	w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

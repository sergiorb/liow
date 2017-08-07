package controllers

import (
	"github.com/gorilla/mux"
	"github.com/sergiorb/liow/server/models"
  "github.com/sergiorb/liow/server/entities"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"net/http"
	"strings"
	 "fmt"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) Read(w http.ResponseWriter, r *http.Request) {

	var ap *entities.ApiResponse
  vars := mux.Vars(r)

	id := vars["id"]

	if strings.Trim(id, " ") == "" {

		ap = &entities.ApiResponse{Status:"No valid id"}
		w.WriteHeader(http.StatusBadRequest)

	} else {

		userDao := models.NewUserDao(uc.session)
	  defer userDao.CloseSession()

		user, err := userDao.Read(id)

		if err != nil {

			ap = &entities.ApiResponse{Status:"User not found"}
			w.WriteHeader(http.StatusNotFound)

		} else {

			ap = &entities.ApiResponse{Status:fmt.Sprintf("User %v exist", user.Name)}
			w.WriteHeader(http.StatusOK)
		}
	}

	payload, _ := json.Marshal(&ap)

	w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

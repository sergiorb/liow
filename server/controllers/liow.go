package controllers

import (
	"github.com/sergiorb/liow/server/entities/api"
	"encoding/json"
	"net/http"
)

type LiowController struct {
}

func NewLiowController() *LiowController {
	return &LiowController{}
}

func (lc LiowController) Ping(w http.ResponseWriter, r *http.Request) {

	readResponse := &api.ReadResponse {
    Message:  "pong",
  }

  payload, _ := json.Marshal(&readResponse)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

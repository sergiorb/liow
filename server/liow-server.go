package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/op/go-logging"
  "os"
  "encoding/json"
  "net/http"
)

const (
  HOST              = ""
  PORT              = "8081"
  URL_PREFIX        = "/liow"
  API_VERSION       = "0.1"
  // API_PRFIX   = fmt.Sprintf("%v/api/v%v", URL_PREFIX, API_VERSION)
  API_PRFIX         = "/liow/v0.1"
  TOKEN_HEADER_NAME = "liow-token"
)

func init() {

	stderrorLog := logging.NewLogBackend(os.Stderr, "", 0)

	stderrorLogLeveled := logging.AddModuleLevel(stderrorLog)
	stderrorLogLeveled.SetLevel(logging.INFO, "")

	logging.SetBackend(stderrorLogLeveled)
}

var log = logging.MustGetLogger("log-in-out-watcher server")


func main()  {

  log.Info(fmt.Sprintf("Starting log-in-out-watcher server v%v on %v:%v", API_VERSION, HOST, PORT))

  r := mux.NewRouter()

  r.HandleFunc(fmt.Sprintf("%v%v", API_PRFIX, "/log/in"), LoginHandler).Methods("POST")
  r.HandleFunc(fmt.Sprintf("%v%v", API_PRFIX, "/log/out"), LogoutHandler).Methods("POST")

  log.Info("Listening...")
  http.ListenAndServe(fmt.Sprintf("%v:%v", HOST, PORT), r)
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {

  var ap *ApiResponse

  token := r.Header.Get(TOKEN_HEADER_NAME)

  if token == "" {

    ap = &ApiResponse{Status:"no token provided"}

    w.WriteHeader(http.StatusBadRequest)

  } else {

    ap = &ApiResponse{Status:"OK"}

    w.WriteHeader(http.StatusOK)

    log.Info(fmt.Sprintf("Log in for key: %v", token))
  }

  response, _ := json.Marshal(&ap)

  w.Header().Set("Content-Type", "application/json")

  fmt.Fprint(w, string(response))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

  var ap *ApiResponse

  token := r.Header.Get(TOKEN_HEADER_NAME)

  if token == "" {

    ap = &ApiResponse{Status:"no token provided"}

    w.WriteHeader(http.StatusBadRequest)

  } else {

    ap = &ApiResponse{Status:"OK"}

    w.WriteHeader(http.StatusOK)

    log.Info(fmt.Sprintf("Log out for key: %v", token))
  }

  response, _ := json.Marshal(&ap)

  w.Header().Set("Content-Type", "application/json")

  fmt.Fprint(w, string(response))
}


type ApiResponse struct {

  Status string `json:"status"`
}

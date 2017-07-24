package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/op/go-logging"
  "os"
  "encoding/json"
  "net/http"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "time"
)

const (
  HOST              = ""
  PORT              = "8081"
  URL_PREFIX        = "/liow"
  API_VERSION       = "0.1"
  // API_PRFIX   = fmt.Sprintf("%v/api/v%v", URL_PREFIX, API_VERSION)
  API_PRFIX         = "/liow/v0.1"
  TOKEN_HEADER_NAME = "liow-token"
  MONGO_HOST        = "mongodb://localhost:27017"
)

func init() {

	stderrorLog := logging.NewLogBackend(os.Stderr, "", 0)

	stderrorLogLeveled := logging.AddModuleLevel(stderrorLog)
	stderrorLogLeveled.SetLevel(logging.INFO, "")

	logging.SetBackend(stderrorLogLeveled)
}

var log = logging.MustGetLogger("log-in-out-watcher server")

type ApiResponse struct {

  Status string `json:"status"`
}

type RegisterType struct {
  Id    bson.ObjectId "_id, omitempty"
  Name string `json: "name"`
}

type Register struct {
    Id    bson.ObjectId "_id, omitempty"
    Date  time.Time     `json: "date"`
    Type  RegisterType   `json: "type"`
}

type ApiToken struct {
  Token string `json: "token"`
}

type LogInController struct {
	session *mgo.Session
}

type LogOutController struct {
	session *mgo.Session
}

func NewLogInController(s *mgo.Session) *LogInController {
	return &LogInController{s}
}

func NewLogOutController(s *mgo.Session) *LogOutController {
	return &LogOutController{s}
}

func NewApiToken(r *http.Request) *ApiToken {

  token := r.Header.Get(TOKEN_HEADER_NAME)

  return &ApiToken{token}
}

func (t *ApiToken) IsValid() (bool){

  var isValid bool

  if t.Token == "" {

    isValid = false

  } else {

    isValid = true
  }

  return isValid
}

func (lc LogInController) CreateLogIn(w http.ResponseWriter, r *http.Request) {

  var ap *ApiResponse

	apiToken := NewApiToken(r)

	if !apiToken.IsValid() {

    ap = &ApiResponse{Status:"no valid token"}
    w.WriteHeader(http.StatusBadRequest)

	} else {

    ap = &ApiResponse{Status:"OK"}
    w.WriteHeader(http.StatusOK)
  }

  payload, _ := json.Marshal(&ap)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

func (lc LogOutController) CreateLogOut(w http.ResponseWriter, r *http.Request) {

  var ap *ApiResponse

	apiToken := NewApiToken(r)

	if !apiToken.IsValid() {

    ap = &ApiResponse{Status:"no valid token"}
    w.WriteHeader(http.StatusBadRequest)

	} else {

    ap = &ApiResponse{Status:"OK"}
    w.WriteHeader(http.StatusOK)
  }

  payload, _ := json.Marshal(&ap)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
}

func getMongoSession() *mgo.Session {

	session, err := mgo.Dial(MONGO_HOST)

	if err != nil {

		log.Error(err)
		panic(err)
	}

	return session
}

func main()  {

  log.Info(fmt.Sprintf("Starting log-in-out-watcher server v%v on %v:%v", API_VERSION, HOST, PORT))

  r := mux.NewRouter()

  logInController := NewLogInController(getMongoSession())
  logOutController := NewLogOutController(getMongoSession())

  r.Handle(fmt.Sprintf("%v%v", API_PRFIX, "/log/in"), http.HandlerFunc(logInController.CreateLogIn)).Methods("POST")
  r.Handle(fmt.Sprintf("%v%v", API_PRFIX, "/log/out"), http.HandlerFunc(logOutController.CreateLogOut)).Methods("POST")

  log.Info("Listening...")
  http.ListenAndServe(fmt.Sprintf("%v:%v", HOST, PORT), r)
}

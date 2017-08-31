package router

import (
  "net/http"
  "github.com/sergiorb/liow/server/config"
  "github.com/sergiorb/liow/server/entities/api"
  "github.com/sergiorb/liow/server/entities/middleware"
  "github.com/sergiorb/liow/server/models"
  "encoding/json"
  "strings"
  "fmt"
  "context"
)

func checkAPIToken(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    tokenString := r.Header.Get(conf.Api.ApiTokenName)

    log.Debug(fmt.Sprintf("r.Header.Get(\"%v\") => %v",
      conf.Api.ApiTokenName,
      tokenString))

    if len(strings.Trim(tokenString, " ")) == 0 {

      payload, _ := json.Marshal(api.AuthResponse{
        Message:  "Unauthorized",
      })

    	w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusUnauthorized)
      w.Write([]byte(payload))

      return // don't call original handler
    }

    tokenDao := models.NewTokenDao(getMongoSession())
	  defer tokenDao.CloseSession()

    token, err := tokenDao.GetByData(tokenString)

    log.Debug(fmt.Sprintf("tokenDao.GetByData(\"%v\") => %v",
      tokenString,
      token.Id))

    if err != nil {

      log.Debug(fmt.Sprintf("tokenDao.GetByData(\"%v\") => err => %v",
        tokenString,
        err))

      payload, _ := json.Marshal(api.AuthResponse{
        Message:  "Unauthorized",
      })

      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusUnauthorized)
      w.Write([]byte(payload))

      return // don't call original handler
    }

    requestContext := middleware.RequestContext{}

    requestContext.Token  = token

    ctx := context.WithValue(r.Context(), config.REQUEST_CONTEXT_KEY, requestContext)

    h.ServeHTTP(w, r.WithContext(ctx))
  })
}

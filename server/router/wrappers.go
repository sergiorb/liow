package router

import (
  "net/http"
  "github.com/sergiorb/liow/server/entities/api"
  "encoding/json"
  "strings"
)

func checkAPIToken(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    token := r.Header.Get(conf.Api.ApiTokenName)

    if len(strings.Trim(token, " ")) == 0 {

      payload, _ := json.Marshal(api.AuthResponse{
        Message:  "Unauthorized",
      })

    	w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusUnauthorized)
      w.Write([]byte(payload))

      return // don't call original handler
    }
    
    h.ServeHTTP(w, r)
  })
}

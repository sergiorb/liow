package router

import (
  "net/http"
  "github.com/sergiorb/liow/server/entities/api"
  "github.com/sergiorb/liow/server/utils"
  "encoding/json"
  "strings"
  "fmt"
)

func withAuth(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    tokenString := r.Header.Get(conf.Api.ApiTokenName)

    if len(strings.Trim(tokenString, " ")) == 0 {

      payload, _ := json.Marshal(api.AuthResponse{
        Message:  "Unauthorized",
      })

    	w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusUnauthorized)
      w.Write([]byte(payload))

      return // don't call original handler
    }

    ctx, requestContext := utils.LoadOrCreateRequestContext(r)

    if !requestContext.IsAuthenticated() {

      payload, _ := json.Marshal(api.AuthResponse{
        Message:  "Unauthorized",
      })

      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusUnauthorized)
      w.Write([]byte(payload))

      return // don't call original handler
    }

    h.ServeHTTP(w, r.WithContext(ctx))
  })
}

func withRole(role string, h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    ctx, requestContext := utils.LoadOrCreateRequestContext(r)

    log.Debug(fmt.Sprintf("[Token: %v] [Endpoint Role: %v] [Request Role: %v]",
      requestContext.Token.Data,
      role,
      requestContext.Role.Name))

    if requestContext.Role.Name != role {

      payload, _ := json.Marshal(api.AuthResponse{
        Message:  "Forbidden",
      })

      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusForbidden)
      w.Write([]byte(payload))

      return // don't call original handler
    }

    h.ServeHTTP(w, r.WithContext(ctx))
  })
}

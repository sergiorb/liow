package utils

import (
  "net/http"
  "github.com/sergiorb/liow/server/config"
  "github.com/sergiorb/liow/server/entities/middleware"
)

func GetRequestContext(r *http.Request) middleware.RequestContext {

  requestContext, _ := r.Context().Value(config.REQUEST_CONTEXT_KEY).(middleware.RequestContext)

  return requestContext
}

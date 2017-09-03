package utils

import (
  "net/http"
  "github.com/sergiorb/liow/server/config"
  "github.com/sergiorb/liow/server/entities/middleware"
  "github.com/sergiorb/liow/server/entities/authorization"
  "github.com/sergiorb/liow/server/models"
  "fmt"
  "context"
)

func getRC(r *http.Request) (middleware.RequestContext, bool) {

  requestContext, ok := r.Context().Value(config.REQUEST_CONTEXT_KEY).(middleware.RequestContext)

  return requestContext, ok
}

func GetRequestContext(r *http.Request) middleware.RequestContext {

  rc, _ := getRC(r)

  return rc
}

func LoadOrCreateRequestContext(r *http.Request) (context.Context, middleware.RequestContext){

  var ctx context.Context
  var requestContext middleware.RequestContext
  var ok bool
  var tokenString string

  ctx = r.Context()

  tokenString = r.Header.Get(conf.Api.ApiTokenName)

  if requestContext, ok = getRC(r); !ok || !requestContext.Loaded {

    log.Debug(fmt.Sprintf("[Token: %v] Loading requestContext...", tokenString))

    tokenDao := models.NewTokenDao(models.GetMongoSession())
    defer tokenDao.CloseSession()

    token, err := tokenDao.GetByData(tokenString)

    if err != nil {

      log.Debug(fmt.Sprintf("[Token: %v] Can't load token. Error: %v", tokenString, err))

    } else {

      requestContext = middleware.RequestContext{
        Loaded: true,
        Token: token,
        Role : authorization.Role{
          Name: token.Role,
        },
      }

      ctx = context.WithValue(r.Context(), config.REQUEST_CONTEXT_KEY, requestContext)
    }

  } else {

    requestContext = GetRequestContext(r)

    log.Debug(fmt.Sprintf("[Token: %v] RequestContext already loaded", tokenString))
  }

  return ctx, requestContext
}
/*
func LoadOrCreateRequestContext(r *http.Request) middleware.RequestContext {
  /*
  var requestContext middleware.RequestContext
  var ok bool

  tokenString := r.Header.Get(conf.Api.ApiTokenName)

  log.Debug(fmt.Sprintf("[Token: %v] Loading request context...", tokenString))

  if requestContext, ok = getRC(r); ok {

    log.Debug(fmt.Sprintf("[Token: %v] Request context already loaded!", tokenString))

  } else {


    := models.NewTokenDao(models.GetMongoSession())
	  defer tokenDao.CloseSession()

    token, err := tokenDao.GetByData(tokenString)

    if err != nil {

      log.Debug(fmt.Sprintf("[Token: %v] Can't load token", tokenString))

      requestContext = middleware.RequestContext{}

    } else {

      log.Debug(fmt.Sprintf("[Token: %v] Request Context loaded succesfully", tokenString))

      requestContext = middleware.RequestContext{
        Token:  token,
        Role:   authorization.Role{
          Name: token.Role,
        },
      }
    }
  }*/

  /*ctx := context.WithValue(r.Context(), config.REQUEST_CONTEXT_KEY, middleware.RequestContext{})

  log.Debug("====================================================")
  log.Debug(ctx)
  log.Debug("====================================================")

  rc, ok := r.Context().Value(config.REQUEST_CONTEXT_KEY).(middleware.RequestContext)

  return requestContext
}*/

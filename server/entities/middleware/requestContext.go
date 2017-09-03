package middleware

import (
  "github.com/sergiorb/liow/server/models"
  "github.com/sergiorb/liow/server/entities/authorization"
)

type RequestContext struct {

  Loaded  bool
  Token   models.Token
  Role    authorization.Role
}


func (rc *RequestContext) IsAuthenticated() bool {

  return rc.Loaded
}

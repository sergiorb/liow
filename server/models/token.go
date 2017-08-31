package models

import (
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Token struct {
    Id            bson.ObjectId "_id,omitempty"
    creationDate  time.Time     `json:"creationDate"`
    Data          string        `json:"data"`
    Email         string        `json:"email"`
}

type TokenDAO struct {

  session *mgo.Session
}

func NewTokenDao(session *mgo.Session) *TokenDAO {

  return &TokenDAO{
    session: session.Copy(),
  }
}

func (td *TokenDAO) CloseSession() {

  td.session.Close()
}

func (td *TokenDAO) Read(id string) (Token, error) {

  var token Token

  c := td.session.DB(conf.Database.Name).C(TOKEN_COLLECTION_NAME)

  err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&token)

  return token, err
}

func (td *TokenDAO) GetByData(dataString string) (Token, error) {

  var token Token

  c := td.session.DB(conf.Database.Name).C(TOKEN_COLLECTION_NAME)

  err := c.Find(bson.M{"data": dataString}).One(&token)

  return token, err
}

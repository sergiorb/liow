package models

import (
  "time"
  //"fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Token struct {
    Id            bson.ObjectId "_id, omitempty"
    creationUser  User          `json: "creationUser"`
    creationDate  time.Time     `json: "creationDate"`
    data          string        `json: "data"`
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

/*
func NewToken()  {

  now := Time.now()

  return &User{
    CreationDate: now
    UpdateDate: now
    Name: name
    Surname: surname
  }
}

func FindByData(data string) {

}

func (t *Token) isRegistered(session *mgo.Session) bool, error {

  var dbToken Token

  c := session.DB(conf.Database.Name).C(TOKEN_DB_COLLECTION)

  err := c.Find(bson.M{"data": t.data}).One(&dbToken)

  return dbToken != nil, err
}
*/

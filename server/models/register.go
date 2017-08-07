package models

import (
  "time"
  //"fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Register struct {
    Id            bson.ObjectId "_id, omitempty"
    creationToken Token         `json: "creationToken"`
    creationDate  time.Time     `json: "creationDate"`
    Type          string        `json: "type"`
}

type RegisterDAO struct {

  session *mgo.Session
}

func NewRegisterDao(session *mgo.Session) *RegisterDAO {

  return  &RegisterDAO{
    session: session.Copy(),
  }
}

func (rd *RegisterDAO) CloseSession() {

  rd.session.Close()
}

func (rd *RegisterDAO) Read(id string) (Register, error) {

  var register Register

  c := rd.session.DB(conf.Database.Name).C(REGISTER_COLLECTION_NAME)

  err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&register)

  return register, err
}

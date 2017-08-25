package models

import (
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Register struct {
    Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Token         string        `json:"token"`
    CreationDate  time.Time     `json:"creationDate"`
    Data          interface{}   `json:"data"`
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

func (rd *RegisterDAO) Create(register *Register) error {

  register.Id = bson.NewObjectId()
  register.CreationDate = time.Now()

  c := rd.session.DB(conf.Database.Name).C(REGISTER_COLLECTION_NAME)

  return c.Insert(&register)
}

func (rd *RegisterDAO) Read(id string) (Register, error) {

  var register Register

  c := rd.session.DB(conf.Database.Name).C(REGISTER_COLLECTION_NAME)

  err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&register)

  return register, err
}

func (r *Register) GetErrors() map[string]error {

  return nil;
}

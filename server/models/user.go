package models

import (
  "time"
  //"fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id            bson.ObjectId "_id, omitempty"
    creationUser  *User         `json: "creationUser"`
    creationDate  time.Time     `json: "creationDate"`
    Name          string        `json: "name"`
    Surname       string        `json: "surname"`
}

type UserDAO struct {

  session *mgo.Session
}

func NewUserDao(session *mgo.Session) *UserDAO {

  return  &UserDAO{
    session: session.Copy(),
  }
}

func (ud *UserDAO) CloseSession() {

  ud.session.Close()
}

func (ud *UserDAO) Read(id string) (User, error) {

  var user User

  c := ud.session.DB(conf.Database.Name).C(USER_COLLECTION_NAME)

  err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)

  return user, err
}

/*func (ud *UserDAO) Save(u *User) User, error {

  c := ud.session.DB(conf.Database.Name).C(USER_COLLECTION_NAME)
  err := c.Insert(&u)

  return u, err
}*/

/*
func NewUser(name string, surname string)  {

  now := Time.now()

  return &User{
    creationDate: now
    updateDate: now
    name: name
    surname: surname
  }
}

func (u *User) Update(name string, surname string)  {

  u.updateDate = Time.now()
  u.name = name
  u.surname = surname
}*/

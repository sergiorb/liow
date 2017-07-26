package model

type RegisterType struct {
  Id    bson.ObjectId "_id, omitempty"
  Name  string        `json: "name"`
}

type Register struct {
    Id    bson.ObjectId "_id, omitempty"
    Date  time.Time     `json: "date"`
    Type  RegisterType  `json: "type"`
}

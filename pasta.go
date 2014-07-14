package main

import (
  "errors"
  "labix.org/v2/mgo/bson"
  "time"
)

type Pasta struct {
  UID       string `bson:"uid"`
  Content   string `form:"content"`
  CreatedAt time.Time
}

func PastaAll() []Pasta {
  var pastas []Pasta
  DB.C("pastas").Find(nil).All(&pastas)

  return pastas
}

func PastaGet(uid string) (Pasta, error) {
  var pasta Pasta

  // err here not working for some reason
  if err := DB.C("pastas").Find(bson.M{"uid": uid}).One(&pasta); err != nil {
    return pasta, err
  }

  // hack
  if len(pasta.Content) == 0 {
    return pasta, errors.New("Pasta not found")
  }

  return pasta, nil
}

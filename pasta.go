package main

import (
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

func PastaGet(uid string) Pasta {
	var pasta Pasta
	DB.C("pastas").Find(bson.M{"uid": uid}).One(&pasta)

	return pasta
}

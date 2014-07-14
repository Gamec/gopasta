package main

import (
	"labix.org/v2/mgo"
	"os"
)

var DB *mgo.Database

func init() {
	session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if err != nil {
		panic(err)
	}

	DB = session.DB("")
}

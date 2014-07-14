package main

import (
	"labix.org/v2/mgo"
)

var DB *mgo.Database

func init() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	DB = session.DB("gopasta")
}

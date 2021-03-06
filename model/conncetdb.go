package model

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func ConnectUser() (*mgo.Collection, error) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	db := session.DB("tyr-project")
	c := db.C("user")
	fmt.Println(c)
	return c, err
}

func ConnectList() (*mgo.Collection, error) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	db := session.DB("tyr-project")
	c := db.C("list")
	return c, err
}

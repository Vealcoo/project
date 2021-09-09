package model

import "gopkg.in/mgo.v2"

func ConnectUser() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	db := session.DB("tyr-project")
	c := db.C("user")
}

func ConnectList() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	db := session.DB("tyr-project")
	c := db.C("list")
}

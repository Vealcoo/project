package controller

import (
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

func Login(id string, pw string) bool {
	var auth bool = false
	c, err := model.ConnectUser()
	if err != nil {
		panic(err)
	}
	result := model.UserInfo{}
	err = c.Find(bson.M{"userid": id, "userpw": pw}).One(result)
	if err != nil {
		auth = true
	}
	return auth
}

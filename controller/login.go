package controller

import (
	"fmt"
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

func Login(id string, pw string) bool {
	var auth bool = true
	c := model.ConnectUser()
	err := c.Find(bson.M{"userid": id, "userpw": pw})
	if err != nil {
		fmt.Println(err)
		auth = false
	}
	return auth
}

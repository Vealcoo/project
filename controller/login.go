package controller

import (
	"fmt"

	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

func Login(id string, pw string) {
	ConnectUser()
	result := UserInfo{}
	err = c.Find(bson.M{"userid": id, "userpw": pw}).One(&result)
	if err == nil {
		setSession(id)
	}
	fmt.Println(result)
}

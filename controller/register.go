package controller

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func Register(id string, pw string, name string) {
	ConnectUser()
	result := UserInfo{}
	err = c.Find(bson.M{"userid": id}).One(&result)
	if err != nil {
		err = c.Insert(NewUserInfo(id, pw, name))
		if err != nil {

		}
	} else {
		fmt.Println("id exist")
	}
}

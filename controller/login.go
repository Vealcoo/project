package controller

import (
	"fmt"
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

type LoginInfo struct {
	UserId string `json:"userid"`
	UserPw string `json:"userpw"`
}

func Login(id string, pw string) int {
	c, err := model.ConnectUser()
	if err != nil {
		panic(err)
	}
	if id == "" || pw == "" {
		return 1
	}
	result := model.UserInfo{}
	err = c.Find(bson.M{"userid": id, "userpw": pw}).One(&result)
	fmt.Println(result)
	if err != nil {
		return 2
	}
	return 0
}

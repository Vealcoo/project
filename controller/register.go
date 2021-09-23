package controller

import (
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

type RegisterInfo struct {
	UserId   string `json:"userid"`
	UserPw   string `json:"userpw"`
	UserName string `json:"username"`
}

func Register(id string, pw string, name string) int {
	if id == "" || pw == "" || name == "" {
		return 1
	}
	c, err := model.ConnectUser()
	if err != nil {
		return 2
	}
	result := model.UserInfo{}
	err = c.Find(bson.M{"userid": id}).One(&result)
	if err != nil {
		err = c.Insert(model.NewUserInfo(id, pw, name))
		if err != nil {
			return 3
		}
		return 0
	} else {
		return 4
	}
}

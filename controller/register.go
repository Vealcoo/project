package controller

import (
	"fmt"
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

type RegisterInfo struct {
	UserId   string `json:"userid"`
	UserPw   string `json:"userpw"`
	UserName string `json:"username"`
}

func Register(id string, pw string, name string) (error, bool) {
	var auth bool
	c, err := model.ConnectUser()
	if err != nil {
		panic(err)
	}
	result := model.UserInfo{}
	err = c.Find(bson.M{"userid": id}).One(&result)
	if err != nil {
		err = c.Insert(model.NewUserInfo(id, pw, name))
		auth = true
		if err != nil {
			fmt.Println(err)
		}
	} else {
		auth = false
	}
	return err, auth
}

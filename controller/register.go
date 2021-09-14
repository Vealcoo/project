package controller

import (
	"fmt"
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

func Register(id string, pw string, name string) error {
	c := model.ConnectUser()
	result := model.UserInfo{}
	err := c.Find(bson.M{"userid": id}).One(&result)
	if err != nil {
		err = c.Insert(model.NewUserInfo(id, pw, name))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("id exist")
	}
	return nil
}

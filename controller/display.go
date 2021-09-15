package controller

import (
	"fmt"
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

func Display(id string) model.ListInfo {
	c, err := model.ConnectList()
	if err != nil {
		panic(err)
	}
	result := model.ListInfo{}
	err = c.Find(bson.M{"userid": id}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

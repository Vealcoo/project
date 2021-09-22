package controller

import (
	"fmt"
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

type DisplayInfo struct {
	UserId string `json:"userid"`
}

func Display(id string) []model.ListInfo {
	c, err := model.ConnectList()
	if err != nil {
		panic(err)
	}
	var result []model.ListInfo
	err = c.Find(bson.M{"userid": id}).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

package controller

import (
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

type DisplayInfo struct {
	UserId string `json:"userid"`
}

func Display(id string) ([]model.ListInfo, int) {
	var result []model.ListInfo
	c, err := model.ConnectList()
	if err != nil {
		return result, 1
	}
	err = c.Find(bson.M{"userid": id}).All(&result)
	if err != nil {
		return result, 2
	}
	return result, 0
}

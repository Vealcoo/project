package controller

import (
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

type DeleteListInfo struct {
	UserId string `json:"userid"`
	ListId string `json:"listid"`
}

func Delete(listid string) int {
	if listid == "" {
		return 1
	}
	selector := bson.M{"_id": bson.ObjectIdHex(listid)}
	c, err := model.ConnectList()
	if err != nil {
		return 2
	}
	err = c.Remove(selector)
	if err != nil {
		return 3
	}
	return 0
}

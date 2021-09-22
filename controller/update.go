package controller

import (
	"fmt"
	"time"
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

type UpdateListInfo struct {
	ListId      string `json:"listid"`
	UserId      string `json:"userid"`
	ListTitle   string `json:"listtitle"`
	ListContext string `json:"listcontext"`
	StartTime   string `json:"starttime"`
	EndTime     string `json:"endtime"`
}

func Update(listid string, id string, title string, context string, start string, end string, timeup bool) error {
	starttime, _ := time.Parse(time.RFC3339, start)
	endtime, _ := time.Parse(time.RFC3339, end)
	c, err := model.ConnectList()
	if err != nil {
		panic(err)
	}
	selector := bson.M{"_id": bson.ObjectIdHex(listid)}
	new := model.NewListInfo(id, title, context, starttime, endtime, timeup)
	err = c.Update(selector, new)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

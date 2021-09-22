package controller

import (
	"time"
	"tyr-project/model"
)

type InsertListInfo struct {
	UserId      string `json:"userid"`
	ListTitle   string `json:"listtitle"`
	ListContext string `json:"listcontext"`
	StartTime   string `json:"starttime"`
	EndTime     string `json:"endtime"`
}

func Insert(id string, title string, context string, start string, end string, timeup bool) int {
	if id == "" || title == "" || context == "" || start == "" || end == "" {
		return 1
	}
	starttime, _ := time.Parse(time.RFC3339, start)
	endtime, _ := time.Parse(time.RFC3339, end)
	c, err := model.ConnectList()
	if err != nil {
		return 2
	}
	err = c.Insert(model.NewListInfo(id, title, context, starttime, endtime, timeup))
	if err != nil {
		return 3
	}
	return 0
}

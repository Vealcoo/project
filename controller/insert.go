package controller

import (
	"fmt"
	"time"
	"tyr-project/model"
)

func Insert(id string, title string, context string, start string, end string, timeup bool) error {
	starttime, _ := time.Parse(time.RFC3339, start)
	endtime, _ := time.Parse(time.RFC3339, end)
	c, err := model.ConnectList()
	if err != nil {
		panic(err)
	}
	err = c.Insert(model.NewListInfo(id, title, context, starttime, endtime, timeup))
	fmt.Println(c)
	if err != nil {
		fmt.Println("err")
	}
	return nil
}

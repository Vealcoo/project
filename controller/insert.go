package controller

import (
	"fmt"
	"time"
	"tyr-project/model"
)

func Insert(id string, title string, context string, start time.Time, end time.Time, timeup bool) {
	ConnectList()
	err = c.Insert(NewListInfo(id, title, context, start, end, timeup))
	if err != nil {
		fmt.Println(err)
	}
}

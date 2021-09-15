package controller

import (
	"fmt"
	"tyr-project/model"

	"gopkg.in/mgo.v2/bson"
)

func Delete(listid string) error {
	selector := bson.M{"_id": bson.ObjectIdHex(listid)}
	c, err := model.ConnectList()
	if err != nil {
		panic(err)
	}
	err = c.Remove(selector)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

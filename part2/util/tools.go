package util

import (
	"encoding/json"
	"fmt"
)

var ToDoArr = []ToDOThing{
	{Id: 1, Name: "Buy computer", Period: "40m", Status: "Undone"},
	{Id: 2, Name: "Buy banana", Period: "20m", Status: "done"},
	{Id: 3, Name: "Buy peach", Period: "20m", Status: "done"},
	{Id: 4, Name: "Buy melon", Period: "10m", Status: "Undone"},
}

type ToDOThing struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Period string `json:"TimeNeeded"`
	Status string `json:"Status"`
}

func ListToDo(strArr ...ToDOThing) {
	fmt.Println(strArr)
	marshal, err := json.Marshal(strArr)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

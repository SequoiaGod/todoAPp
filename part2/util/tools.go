package util

import (
	"encoding/json"
	"fmt"
)

var ToDoArr = make([]ToDOThing, 0)

type ToDOThing struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Period string `json:"TimeNeeded"`
	Status string `json:"Status"`
}
type Message struct {
	Command  string
	Todo     ToDOThing
	Id       int
	Response chan<- string
}

func ListToDo(strArr ...ToDOThing) {
	fmt.Println(strArr)
	marshal, err := json.Marshal(strArr)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func RemoveItem(id int) ToDOThing {
	for i := 0; i < len(ToDoArr); i++ {
		if ToDoArr[i].Id == id {
			temp := ToDoArr[i]
			ToDoArr = append(ToDoArr[:i], ToDoArr[i+1:]...)
			return temp
		}
	}
	return ToDOThing{}
}

func UpdateItem(todo ToDOThing) {
	for i := 0; i < len(ToDoArr); i++ {
		if ToDoArr[i].Id == todo.Id {
			ToDoArr[i] = todo
			return
		}
	}
	ToDoArr = append(ToDoArr, todo)
}

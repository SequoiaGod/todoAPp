package util

import (
	"encoding/json"
	"fmt"
)

var ToDoArr = make([]ToDOThing, 0)

const CONNECT_STR = "postgresql://postgres:12125412@localhost:5432/todo_list?sslmode=disable"

type ToDOThing struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Period string `json:"TimeNeeded"`
	Status string `json:"Status"`
}

type ToDoPostgreSQL struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Period string `db:"period"`
	Status string `db:"status"`
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

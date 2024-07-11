package util

var ToDoArr []ToDOThing

type ToDOThing struct {
	Name   string `json:"Name"`
	Period string `json:"TimeNeeded"`
	Status string `json:"Status"`
}

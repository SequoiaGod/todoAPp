package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Name   string  `json:"Name"`
	Price  float64 `json:"Price"`
	Volume int64   `json:"Volume"`
}

func ListToDo(itemArr ...Item) {
	fmt.Println(itemArr)
	f, err := os.Create("./TodoThingToJSONFile/tings.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(itemArr); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	//marshal, err := json.Marshal(itemArr)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(marshal))
}

func main() {
	itemArr := []Item{
		{"apple", 3.3, 11},
		{"peach", 4.3, 22},
		{"melon", 23.3, 33},
		{"grapefruit", 34.3, 44},
		{"mango", 32.3, 55},
		{"orange", 333, 66},
		{"strawberry", 98.3, 778},
		{"kiwi fruit", 35, 88},
		{"blueberry", 233.3, 99},
	}
	ListToDo(itemArr...)
}

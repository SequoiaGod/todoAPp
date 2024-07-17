package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Item struct {
	Name   string  `json:"Name"`
	Price  float64 `json:"Price"`
	Volume int64   `json:"Volume"`
}

func main() {
	file, err := os.Open("./part1/12TodoThingToJSONFile/tings.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var todos []Item
	err = json.NewDecoder(file).Decode(&todos)

	for _, todo := range todos {
		fmt.Printf("Name: %s, Price: %f, Volume: %d\n", todo.Name, todo.Price, todo.Volume)
	}
}

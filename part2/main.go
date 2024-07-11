package main

import (
	"encoding/json"
	"fmt"
	"todoAPp/part2/util"
)

func main() {
	s := "{\"Id\": 1, \"Name\": \"Buy computer\", \"Period\": \"40m\", \"Status\": \"Undone\"}"
	data := util.ToDOThing{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

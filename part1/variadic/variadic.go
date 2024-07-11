package main

import (
	"encoding/json"
	"fmt"
)

func ListToDo(strArr ...string) {
	fmt.Println(strArr)
	marshal, err := json.Marshal(strArr)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func main() {
	strArr := []string{"add money", "show balance", "deduct money"}
	ListToDo(strArr...)
}

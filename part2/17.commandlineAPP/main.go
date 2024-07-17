package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"todoAPp/part2/util"
)

var (
	hasCommand = make(chan string, 1)
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	fmt.Println("starting server")
	startServer(ctx)

	<-terminate
	done()
}

func startServer(ctx context.Context) {
	go startCommandLine(ctx)
	go startListenCommand(ctx)
}

func startCommandLine(ctx context.Context) {
	scanner := bufio.NewReader(os.Stdin)

	for {
		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		select {
		case <-ctx.Done():
			return
		default:
			switch input {
			case "create":
				fmt.Println("Please input todo thing to the command line: ")
				item := readSingleLine(scanner)
				hasCommand <- "create*" + item
			case "read":
				fmt.Println("Now all of todo things can be found below： ")
				hasCommand <- "read"
			case "update":
				fmt.Println("Please input the id of todo thing which need be amended: ")
				item := readSingleLine(scanner)
				hasCommand <- "update*" + item
			case "delete":
				fmt.Println("Please input the id of todo thing which need be delete: ")
				item := readSingleLine(scanner)
				hasCommand <- "delete*" + item
			default:
				fmt.Println("Warning bad input!!! Available command is below: ")
				fmt.Println("create   read   update   delete")
			}
		}
	}
}

// {"Id":5,"Name":"test","TimeNeeded":"40m","Status":"Undone"}
func startListenCommand(ctx context.Context) {
	go func() {
		<-ctx.Done()
		close(hasCommand)
	}()
	for val := range hasCommand {
		commandArr := strings.Split(val, "*")
		switch commandArr[0] {
		case "create":
			fmt.Println("---" + commandArr[1])
			util.ToDoArr = append(util.ToDoArr, convertStringToStruct(commandArr[1]))
		case "read":
			util.ListToDo(util.ToDoArr...)
		case "update":
			updateElement(convertStringToStruct(commandArr[1]))
		case "delete":
			deleteElement(commandArr[1])
		}
	}
}

func readSingleLine(scanner *bufio.Reader) string {
	input, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimSpace(input)
}

func convertStringToStruct(str string) util.ToDOThing {
	data := util.ToDOThing{}
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		panic(err)
	}
	return data
}

func deleteElement(id string) {
	for i := 0; i < len(util.ToDoArr); i++ {
		atoi, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		if util.ToDoArr[i].Id == atoi {
			util.ToDoArr = append(util.ToDoArr[:i], util.ToDoArr[i+1:]...)
		}
	}
}

func updateElement(todo util.ToDOThing) {
	for i := 0; i < len(util.ToDoArr); i++ {
		if util.ToDoArr[i].Id == todo.Id {
			util.ToDoArr[i] = todo
		}
	}
}

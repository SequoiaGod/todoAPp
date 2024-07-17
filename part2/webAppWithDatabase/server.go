package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"todoAPp/part2/util"
)

var (
	commandChan = make(chan util.Message)
	db, _       = sql.Open("postgres", util.CONNECT_STR)
)

func initServer(ctx context.Context) {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo", getTodo)
	mux.HandleFunc("POST /todo", updateTodo)
	mux.HandleFunc("DELETE /todo/{key}", deleteTodo)
	mux.HandleFunc("PUT /todo", updateTodo)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	serverRun(server, ctx)
}

func serverRun(server *http.Server, ctx context.Context) {
	var err error
	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	go func() {
		<-ctx.Done()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Println(err)
		}
		log.Println("Server has been stopped...")
	}()
	commandRun(ctx)

}

func commandRun(ctx context.Context) {
	go func() {
		<-ctx.Done()
		close(commandChan)
	}()
	go func() {
		for message := range commandChan {
			switch message.Command {
			case "getAll":
				var str strings.Builder
				str.WriteString("{")
				for _, v := range util.ToDoArr {
					str.WriteString("\n" + fmt.Sprintf("%v", v))
				}
				str.WriteString("\n}")
				message.Response <- str.String()
			case "update":
				util.UpdateItem(message.Todo)
				message.Response <- fmt.Sprintf("%v", message.Todo)
			case "delete":
				item := util.RemoveItem(message.Id)
				message.Response <- fmt.Sprintf("%v", item)
			default:
				log.Println("incorrect command")

			}
		}
	}()

}

func updateTodo(writer http.ResponseWriter, request *http.Request) {
	log.Println("Create/Update a todo thing...")
	response := make(chan string)
	todo := decoder(writer, request)
	commandChan <- util.Message{Command: "update", Todo: todo, Response: response}
	fmt.Fprintf(writer, <-response)
}

func deleteTodo(writer http.ResponseWriter, request *http.Request) {
	key, err := strconv.Atoi(request.PathValue("key"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	response := make(chan string)
	commandChan <- util.Message{Command: "delete", Id: key, Response: response}
	fmt.Fprintln(writer, <-response)

}

func getTodo(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get All of todo things...")
	response := make(chan string)
	commandChan <- util.Message{Command: "getAll", Response: response}
	fmt.Fprintf(writer, <-response)
}

func decoder(writer http.ResponseWriter, request *http.Request) util.ToDOThing {
	decoder := json.NewDecoder(request.Body)
	var todo util.ToDOThing
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	return todo
}

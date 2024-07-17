package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todoAPp/part2/util"
)

func mockCommand(ctx context.Context) {
	go func() {
		<-ctx.Done()
		close(commandChan)
	}()
	go func() {
		for message := range commandChan {
			switch message.Command {
			case "getAll":
				message.Response <- "all data"
			case "update":
				util.UpdateItem(message.Todo)
				message.Response <- "update successfully"
			case "delete":
				message.Response <- fmt.Sprintf("Item has been delete successfully...")
			default:
				log.Println("incorrect command")

			}
		}
	}()
}
func Test_commandRun(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandRun(tt.args.ctx)
		})
	}
}

func TestDecoder(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		expectedToDo   util.ToDOThing
		expectedStatus int
	}{
		{
			name:           "Valid JSON",
			body:           `{"Id": 1, "Name": "Task 1", "TimeNeeded": "1h", "Status": "Pending"}`,
			expectedToDo:   util.ToDOThing{Id: 1, Name: "Task 1", Period: "1h", Status: "Pending"},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid JSON",
			body:           `{"Id": 1, "Name": "Task 1", "TimeNeeded": "1h", "Status": "Pending"`,
			expectedToDo:   util.ToDOThing{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty Body",
			body:           ``,
			expectedToDo:   util.ToDOThing{},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			writer := httptest.NewRecorder()

			decodedToDo := decoder(writer, request)

			if writer.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, writer.Code)
			}

			if writer.Code == http.StatusOK {
				if decodedToDo != tt.expectedToDo {
					t.Errorf("expected %v, got %v", tt.expectedToDo, decodedToDo)
				}
			}
		})
	}
}

func Test_deleteTodo(t *testing.T) {

	var tests = []struct {
		name           string
		key            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Key",
			key:            "1",
			expectedStatus: http.StatusOK,
			expectedBody:   "Item has been delete successfully...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, done := context.WithCancel(context.Background())
			go mockCommand(ctx)
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/todo/%s", tt.key), nil)
			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			deleteTodo(rr, req)
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rr.Code)
			}

			if body := rr.Body.String(); body != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, body)
			}
			defer done()
		})
	}
}

func Test_getTodo(t *testing.T) {

	tests := []struct {
		name           string
		expectedRes    string
		expectedStatus int
	}{
		{
			name:           "test valid",
			expectedStatus: http.StatusOK,
			expectedRes:    "all data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, done := context.WithCancel(context.Background())
			go mockCommand(ctx)
			defer done()
			req := httptest.NewRequest("GET", "/todo", nil)
			writer := httptest.NewRecorder()
			getTodo(writer, req)

			if tt.expectedRes != writer.Body.String() {
				t.Errorf("expected body %q, got %q", tt.expectedRes, writer.Body.String())
			}

			if tt.expectedStatus != writer.Code {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, writer.Code)
			}
		})
	}
}

func Test_initServer(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initServer(tt.args.ctx)
		})
	}
}

func Test_serverRun(t *testing.T) {
	type args struct {
		server *http.Server
		ctx    context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serverRun(tt.args.server, tt.args.ctx)
		})
	}
}

func Test_updateTodo(t *testing.T) {

	tests := []struct {
		name           string
		body           string
		expectedRes    string
		expectedStatus int
	}{
		{
			name:           "test valid",
			body:           `{"Id": 1, "Name": "Task 1", "TimeNeeded": "1h", "Status": "Pending"}`,
			expectedRes:    "update successfully",
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, done := context.WithCancel(context.Background())
			go mockCommand(ctx)
			defer done()
			req := httptest.NewRequest("UPDATE", "/todo", strings.NewReader(tt.body))
			writer := httptest.NewRecorder()
			updateTodo(writer, req)

			if tt.expectedRes != writer.Body.String() {
				t.Errorf("expected body %q, got %q", tt.expectedRes, writer.Body.String())
			}

			if tt.expectedStatus != writer.Code {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, writer.Code)
			}
		})
	}
}

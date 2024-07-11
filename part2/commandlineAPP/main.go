package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
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
}

func startCommandLine(ctx context.Context) {
	scanner := bufio.NewReader(os.Stdin)
	for {
		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(input)
		}
	}
}

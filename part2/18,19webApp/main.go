package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	initServer(ctx)
	<-terminate
	done()
}

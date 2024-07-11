package main

import (
	"context"
	"fmt"
	"sync"
)

var doing = make(chan int, 1)
var fruit = make(map[string]int)

func mutxRefactor() {
	var wg sync.WaitGroup
	var mux sync.Mutex

	fruit["apple"] = 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mux.Lock()
			fruit["apple"] = 10
			mux.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mux.Lock()
			fruit["apple"] = 11
			mux.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(fruit["apple"])
	fmt.Println("Ending ---")
}

func run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case num := <-doing:
			//fmt.Println(num)
			fruit["apple"] = num

		}
	}
}

func channelRefactor() {
	ctx, cancel := context.WithCancel(context.Background())
	fruit["apple"] = 0
	var wg sync.WaitGroup
	go run(ctx)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			doing <- 10
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			doing <- 11
		}
	}()

	wg.Wait()
	cancel()
	fmt.Println(fruit["apple"])
	fmt.Println("Ending ---")
}
func main() {
	channelRefactor()
}

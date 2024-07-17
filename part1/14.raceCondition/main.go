package main

import (
	"fmt"
	"sync"
)

func main() {
	fruit := make(map[string]int)
	fruit["apple"] = 0
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fruit["apple"] = 10
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fruit["apple"] = 11
		}
	}()
	wg.Wait()

	fmt.Println("Ending ---")
}

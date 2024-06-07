package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {

	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("foo")
	}
}

func bar() {

	for i := 0; i < 5; i++ {
		fmt.Println("bar")
	}
}

func main() {

	wg.Add(1) // number of goroutines to wait for

	go foo()
	bar()

	wg.Wait() // Block until all goroutines have completed
}

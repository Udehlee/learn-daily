package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {

	select {

	case <-time.After(3 * time.Second):
		fmt.Println("work done")

	case <-ctx.Done():
		fmt.Println("cancelled", ctx.Err())

	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go doWork(ctx)

	time.Sleep(4 * time.Second)
}

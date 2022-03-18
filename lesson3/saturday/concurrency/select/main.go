package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(7 * time.Second)
		cancel()
	}()

	ch := make(chan int)
	go printValues(ch)

	check(ctx, ch)

	fmt.Println("final write")
}

func check(ctx context.Context, ch chan int) {
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ticker.C:
			writeMessage()
		case i := <-ch:
			log.Println("comes print func", i)
		case <-ctx.Done():
			fmt.Println("Done")
			return
		}
	}
}

func writeMessage() {
	log.Println("this is message")
}

func printValues(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
}

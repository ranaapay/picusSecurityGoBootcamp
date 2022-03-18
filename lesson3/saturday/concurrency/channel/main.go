package main

import (
	"fmt"
	"time"
)

func myFunc(ch chan int) {
	fmt.Println(243 + <-ch)
}

func checkChan(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Wrote ", i, "to chan")
	}
	close(ch)  // You can close the channel.
}

func main() {
	fmt.Println("Start Main")

	//Unbuffered channel
	b := 9
	ch := make(chan int)
	go myFunc(ch)
	ch <- b
	go checkChan(ch)

	for {
		res, ok := <-ch  // You can check channel is open or close
		if ok == false {
			fmt.Println("Channel close", ok)
			break
		}
		fmt.Println("Channel open", ok, "res :", res)
	}

	//Buffered Channel
	chBuf := make(chan int, 2)
	go write(chBuf)
	time.Sleep(2 * time.Second)
	for v := range chBuf {
		fmt.Println("read value ", v, "from ch")
		time.Sleep(2 * time.Second)
	}

}

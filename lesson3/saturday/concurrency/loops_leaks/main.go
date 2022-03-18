package main

import "fmt"

func main() {
	// LOOPS
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	/*
		//if you use for with range in a goroutine, it causes complications.
		//goroutine takes last value always.
		for _, val := range a {
			go func() {
				ch <- val * 2
			}()
		}  // it prints 20 20 20 20 20
	*/
	/*
		for _, v := range a {
			k := v  				//You should assign value to variable
			go func() {
				ch <- k * 2
			}()
		}
	*/
	for _, v := range a {
		go func(v int) {
			ch <- v * 2
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}

	//LEAKS
	for i := range countTo(10) {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	//it prints  0 1 2 3 4 5
	// 6 7 8 9 values are not read. if you work with server and main
	//is not closed that causes leak.
}

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

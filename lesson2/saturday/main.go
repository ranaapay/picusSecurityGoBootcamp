package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	//args
	fmt.Println(os.Args)

	args := os.Args[1:]
	for i, arg := range args {
		fmt.Printf("index: %d, value: %s\n", i, arg)
	}

	//functions
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 4, 6, 7))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{6, 1, 4}...))

	result, err := divide(5, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result, err = divide(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	/*
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	*/

	//map ve slice, pointer şeklinde tutulduğu için değiştirecek
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)
}

func modMap(m map[int]string) {

}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
}

//functions
func sum(a, b int) int {
	return a + b
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, nil
}

type funcType func(int, int) int

var opMap = map[string]funcType{}

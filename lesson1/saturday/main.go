package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var flag bool
	fmt.Println(flag)

	var patika = true
	fmt.Println(patika)

	var number = 6
	fmt.Println(number)

	//Composite Types

	//Arrays
	//var x [3]int
	var y = [3]int{10, 20, 30}
	var x = [...]int{10, 20, 30}
	fmt.Println(y == x)

	var matrix [2][3]int
	fmt.Println(matrix)

	//Slices
	// [...] --> array  [] --> slice
	sliceX := []int{10, 20, 30}
	fmt.Println(len(sliceX))

	var sliceY []int
	sliceY = append(sliceY, 10, 20, 30)
	fmt.Println(sliceY)

	var sliceZ []int
	sliceZ = append(sliceZ, sliceY...)
	fmt.Println(sliceZ)

	var k []int
	fmt.Println(k, len(k), cap(k))
	k = append(k, 10)
	fmt.Println(k, len(k), cap(k))
	k = append(k, 20)
	fmt.Println(k, len(k), cap(k))
	k = append(k, 30)
	fmt.Println(k, len(k), cap(k))
	k = append(k, 40)
	fmt.Println(k, len(k), cap(k))
	k = append(k, 50)
	fmt.Println(k, len(k), cap(k))

	//make
	l := make([]int, 5)
	fmt.Println(l)

	l = append(l, 99)
	fmt.Println(l)

	z := make([]int, 5, 10)
	fmt.Println(z, "cap: ", cap(z))

	//Slicing
	s := []int{1, 2, 3, 4}
	d := s[:2]
	fmt.Println(d)
	f := s[1:]
	fmt.Println(f)
	f[2] = 10
	fmt.Println(s) //[1 2 3 10]

	h := make([]int, len(s))
	num := copy(h, s)
	fmt.Println("s:", s, " h:", h, " num:", num)

	//Maps
	var ax map[string]int
	fmt.Println(ax)

	teams := map[string][]string{
		"Ahmet": []string{"a", "b", "c"},
		"Ayşe":  []string{"e", "f", "g"},
	}
	fmt.Println(teams)
	fmt.Println(teams["Ahmet"])
	teams["Ahmet"] = []string{"k", "l"}
	fmt.Println(teams["Ahmet"])

	//Reading and Writing a Map

	totalWins := map[string]int{}
	totalWins["Ayşe"] = 1
	totalWins["Osman"] = 2
	totalWins["Ayşe"]++
	fmt.Println(totalWins)

	v, ok := totalWins["Osman"]
	fmt.Println(v, ok)
	v, ok = totalWins["Kenan"]
	fmt.Println(v, ok)
	delete(totalWins, "Osman")
	fmt.Println(totalWins)

	//Struct
	type person struct {
		name string
		age  int
		pet  string
	}

	osman := person{
		name: "Osman",
		age:  21,
		pet:  "Zeytin",
	}

	osman.age = 22

	//anonymous structs

	var firstPerson struct {
		name string
		age  int
	}
	firstPerson.name = "Ahmet"
	firstPerson.age = 21

	pet := struct {
		cins string
		name string
	}{
		"kedi",
		"zeytin",
	}

	fmt.Println(firstPerson)
	fmt.Println(pet)

	//Shadowing Variables
	c := 10
	if c > 5 {
		fmt.Println(c)
		c := 5
		fmt.Println(c)
	}
	fmt.Println(c)

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println(m)
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	//switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s. \n", os)
	}

	fmt.Println("when's saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far away")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}

	sal := 10
	switch {
	case sal < 12:
		fmt.Println("good morning")
	case sal < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}

	mal := 5
	switch {
	case mal == 5:
		fmt.Println("esit 5")
		fallthrough
	case mal < 10:
		fmt.Println("kucuk 10")
	}
}

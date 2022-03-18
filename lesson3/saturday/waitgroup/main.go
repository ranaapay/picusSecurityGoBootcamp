package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

/*
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("first")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("second")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("third")
	}()
	wg.Wait()
	fmt.Println("done")
}
*/

func main() {
	urls := []string{
		"https://yolcu360.com",
		"https://google.com",
		"https://yahoo.com",
		"https://twitter.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func() {
			getSiteContent(&wg, url)
		}()
	}
	wg.Wait()

}

func getSiteContent(wg *sync.WaitGroup, url string) {
	resp, _ := http.Get(url)
	defer func() {
		wg.Done()
		resp.Body.Close()
	}()
	respBodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("URL : %s\n", url)
	fmt.Printf(string(respBodyBytes))
}

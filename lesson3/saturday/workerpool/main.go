package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getSiteContent(id int, url <-chan string, result chan<- string) {
	for v := range url {
		resp, _ := http.Get(v)
		respBodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("ID : %d, URL : %s\n", id, v)
		result <- string(respBodyBytes)

		fmt.Printf("Sent to channel : %s\n", v)
		resp.Body.Close()
	}
}

func main() {
	urls := []string{
		"https://yolcu360.com",
		"https://google.com",
		"https://yahoo.com",
		"https://twitter.com",
	}

	url := make(chan string, len(urls))
	result := make(chan string, len(urls))

	for i := 0; i < 3; i++ {
		go getSiteContent(i, url, result)
	}

	for i := 0; i < len(urls); i++ {
		url <- urls[i]
	}
	close(url)
	for i := 0; i < len(urls); i++ {
		a := <-result
		fmt.Println(a)
	}
}

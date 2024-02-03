package main

import (
	"fmt"
	"net/http"
)

func main() {

	urls := []string{
		"https://google.com",
		"https://go.dev",
		"https://amazon.com",
	}

	c := make(chan string, len(urls))

	for _, url := range urls {
		go checkLink(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<- c)
	}

}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- url + " seems to be down"
		return
	}
	c <- url + " seems to be up"
}
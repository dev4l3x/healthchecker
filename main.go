package main

import (
	"fmt"
	"net/http"
	"time"
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

	for link := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(link)
	}

}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "seems to be down")
		c <- url
		return
	}
	fmt.Println(url, "seems to be up")
	c <- url
}
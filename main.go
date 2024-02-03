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

	for _, url := range urls {
		checkLink(url)
	}

}

func checkLink(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "seems to be down")		
	}
	fmt.Println(url, "seems to be up")
}
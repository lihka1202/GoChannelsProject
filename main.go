package main

import (
	"fmt"
	"net/http"
)

func main() {
	holders := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	for _, link := range holders {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("UhOh, website is down: ", link)
	} else {
		fmt.Println("No issues!")
	}
}

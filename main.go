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

	c := make(chan string)

	for _, link := range holders {
		go checkLink(link, c)
	}
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("UhOh, website is down: ", link)
		c <- "Likely to be down"
	} else {
		fmt.Println("No issues!")
		c <- "functional website"
	}
}

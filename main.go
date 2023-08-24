package main

import (
	"fmt"
	"net/http"
	"time"
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

	// // For loop to ensure that most of the functions are being called
	// for {
	// 	go checkLink(<-c, c) //pass the value from the channel
	// }

	//The following for loop is equivalent to the one above

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("UhOh, website is down: ", link)
		c <- link
	} else {
		fmt.Println("No issues!")
		c <- link
	}
}

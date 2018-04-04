package main

import (
	"fmt"
	"net/http"
	"time"
)

// Main Routine is 1 time Blocking call
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // Create new channel

	for _, link := range links {
		// checkLink(link)

		go checkLink(link, c) // child Go Routine with channel
	}

	// 1 time blocking calls
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// repeated blocking calls
	// for {
	// 	go checkLink(<-c, c)
	// }

	// repeated blocking calls -- Alternative loop syntax for other logics
	for l := range c {
		// go checkLink(l, c) // repeated child routines immediately

		// Pause program -- repeated child routines after 5 seconds
		// Using FUNCTION LITERAL in Go
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

		// go repeatedCheckLink(l, c)
	}
}

// func repeatedCheckLink(link string, c chan string) {
// 	time.Sleep(5 * time.Second)
// 	checkLink(link, c)
// }

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down I think!?" // 1time blocking calls

		c <- link // repeated blocking calls
		return
	}

	fmt.Println(link, "is OK!")
	// c <- "Yep, it's up!" // 1time blocking calls

	c <- link // repeated blocking calls
}

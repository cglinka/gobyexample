package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	nonbuff := make(chan string)

	messages <- "buffered"
	messages <- "channel"

	nonbuff <- "a test"
	nonbuff <- "refused"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

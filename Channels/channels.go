package main

import "fmt"

func main() {

	messages := make(chan string, 1)

	go func() {
		messages <- "ping"
		messages <- "pong"
	}()

	msg := <-messages
	fmt.Println(len(messages))
	fmt.Println(cap(messages))
	fmt.Println(msg)

}

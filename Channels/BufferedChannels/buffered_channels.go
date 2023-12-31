package main

import "fmt"

func main() {
	buffer := make(chan int, 3)
	buffer <- 1
	buffer <- 2	
	buffer <- 3

	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
}
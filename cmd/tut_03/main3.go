package main

import "fmt"

func main() {
	fmt.Print("Hello Owlrd!")
	// make channel
	var channel = make(chan int)
	go process(channel)
	// read from channel
	var value = <-channel
	fmt.Println(value)
}

func process(c chan int) {
	c <- 1234
}

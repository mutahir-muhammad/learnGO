package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Hello Owlrd!\n")
	// make channel
	var channel = make(chan int, 5) // now it is a buffered channel
	go process(channel)
	// read from channel
	for i := range channel {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

}
func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("This marks the end of the routine process")

}

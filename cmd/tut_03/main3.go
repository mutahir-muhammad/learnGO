package main

import "fmt"

func main() {
	fmt.Print("Hello Owlrd!\n")
	// make channel
	var channel = make(chan int)
	go process(channel)
	// read from channel
	for i := range channel {
		fmt.Println(i)
	}

}
func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}

}

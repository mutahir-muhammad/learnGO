package main

import "fmt"

type gasEngine struct {
	// fields
	milePerGallon uint8
	gallons       uint8
}

func main() {
	var myEngine gasEngine = gasEngine{50, 70} // struct literal
	fmt.Println(myEngine.milePerGallon, myEngine.gallons)
}

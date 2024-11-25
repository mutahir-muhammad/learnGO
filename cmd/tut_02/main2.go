package main

import "fmt"

type gasEngine struct {
	// fields
	milePerGallon uint8
	gallons       uint8
}

func main() {
	var myEngine gasEngine
	fmt.Println(myEngine.milePerGallon, myEngine.gallons)
}

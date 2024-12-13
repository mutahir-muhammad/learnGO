// This Go program uses goroutines to fetch chicken prices from different websites and sends an alert
// message when a deal is found.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5.0

func main() {
	var chickenChannel = make(chan string)
	var websites = []string{"www.chicken.com", "www.chicken.org", "www.chicken.net"}

	for i := range websites {
		go getPrice(websites[i], chickenChannel)
	}
	sendAlertMsg(chickenChannel)
}

func getPrice(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 10
		if chickenPrice > MAX_CHICKEN_PRICE {
			c <- website
			break
		}
	}
}

func sendAlertMsg(c chan string) {
	fmt.Printf("\n FOund a deal at %s", <-c)
}

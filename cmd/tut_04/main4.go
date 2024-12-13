// This Go program uses goroutines to fetch chicken prices from different websites and sends an alert
// message when a deal is found.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5.0
var MAX_EARBUDS_PRICE int = 30

func main() {
	var chickenChannel = make(chan string)
	var earbudsChannel = make(chan string)
	var websites = []string{"www.chicken.com", "www.chicken.org", "www.chicken.net"}
	var budWebsites = []string{"www.pricoye.com", "www.xcessories.com", "www.xiaomi.com", "www.anker.com"}

	for i := range websites {
		go getPrice(websites[i], chickenChannel)
		go getEarbudPrice(budWebsites[i], earbudsChannel)
	}
	sendAlertMsg(chickenChannel, earbudsChannel)
}

func getPrice(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 10
		if chickenPrice < MAX_CHICKEN_PRICE {
			c <- website
			break
		}
	}
}

func getEarbudPrice(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var earbudPrice = rand.Intn(50)
		if earbudPrice < MAX_EARBUDS_PRICE {
			c <- website
			break
		}
	}
}

func sendAlertMsg(c chan string, e chan string) {
	select {
	case website := <-c:
		fmt.Printf("Chicken is on sale at %s\n", website)
	case website := <-e:
		fmt.Printf("Earbuds are on sale at %s\n", website)
	}

}

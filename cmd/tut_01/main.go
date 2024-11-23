package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var c int = a + b

	fmt.Println("a is", a)
	fmt.Println("b is", b)
	fmt.Println("Sum of a and b is", c)
	fmt.Println("Product of a and b is", a*b)
	fmt.Println("Divide b by a is", b/a)
	fmt.Println("modulo of b by a is", b%a)
	fmt.Println("Hello, World!")

	var myStr string = "This is me GOing!"
	fmt.Println(myStr)

	var myRuneIsRuined rune = 'a'
	fmt.Println(myRuneIsRuined)
	myRuneIsRuined = 'b'
	fmt.Println(myRuneIsRuined)

	printMe()
}

func printMe() {
	fmt.Println("I am a function")
}

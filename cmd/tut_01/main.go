package main

import "fmt"

func main() {
	// basic vars
	var a int = 10
	var b int = 20
	var c int = a + b

	// basic print
	fmt.Println("a is", a)
	fmt.Println("b is", b)
	fmt.Println("Sum of a and b is", c)
	fmt.Println("Product of a and b is", a*b)
	fmt.Println("Divide b by a is", b/a)
	fmt.Println("modulo of b by a is", b%a)
	fmt.Println("Hello, World!")

	// booleans are true and false... ofc...lol
	// string
	var myStr string = "This is me GOing!"
	fmt.Println(myStr)

	// runes give ascii value
	var myRuneIsRuined rune = 'a'
	fmt.Println(myRuneIsRuined)
	myRuneIsRuined = 'b'
	fmt.Println(myRuneIsRuined)

	// functions
	printMe()
	fmt.Println(intDivision(10, 2))
	var resultDivide, resultModulo = multipleReturns(10, 3)
	fmt.Println(resultDivide, resultModulo)
}

func printMe() {
	fmt.Println("I am a function")
}

func intDivision(numerator int, denominator int) int {
	return numerator / denominator
}

func multipleReturns(numerator int, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

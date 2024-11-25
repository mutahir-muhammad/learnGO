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
	ifelseCondition(true)
	showArray()
	makeSlices()
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

func ifelseCondition(someBool bool) {
	if someBool {
		fmt.Println("something is true")
	} else {
		fmt.Println("something is false")
	}
}

func showArray() {
	var myArray [5]uint8 = [5]uint8{1, 2, 3, 4, 5}
	// shorthand can be like below or myArray2 := [...]uint8{7, 8, 9, 10, 11}
	// where the fixed length is inferred based on the number of elements
	myArray2 := [5]uint8{7, 8, 9, 10, 11}
	fmt.Println(myArray)
	fmt.Println("my second array is" + fmt.Sprint(myArray2))
	fmt.Println(len(myArray))
	fmt.Println(myArray[0])
	fmt.Println((myArray[1:4]))
}

func makeSlices() {
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 6)
	fmt.Println(mySlice)
	fmt.Printf("The length is %v and the capacity is %v", len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 7)
	fmt.Printf("\n\nThe length is %v and the capacity is %v", len(mySlice), cap(mySlice))

	// slice can be made using make function or var mySlice2 []int = []int{1, 2, 3}
	var mySlice2 []int = make([]int, 5, 10)
	fmt.Println(mySlice2)
	fmt.Printf("The length is %v and the capacity is %v", len(mySlice2), cap(mySlice2))

}

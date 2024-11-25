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
	makeMaps()
	loopingInMaps()
	loopsInGo()
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
	fmt.Println("Inside showArray function\n")
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
	fmt.Println("inside makeSlices function\n")
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

func makeMaps() {
	fmt.Println("Inside makeMaps function\n")
	//make map using make function
	// where map[string]int means that key is a string
	// and the value is an int
	var myMap map[string]int = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Printf("\nThe map is %v", myMap)
	var myMap2 = map[string]int{"one": 1, "two": 2}
	fmt.Printf("\n\nThe second map is %v, the length is %v", myMap2, len(myMap2))

	// the map will always return something
	fmt.Printf("\n\nuninitialized property in map %v", myMap2["three"])
	var value, err = myMap2["four"]
	if err {
		fmt.Printf("the uninitailized property is %v", value)
	} else {
		fmt.Println("\nInvalid property")
	}

	// delete property by reference using key in key-value pair
	delete(myMap, "one")
	fmt.Printf("myMap is now %v: ", myMap)

}

func loopsInGo() {
	var array [5]uint8 = [5]uint8{1, 2, 3}
	// normal loop with index and value
	for i, v := range array {
		fmt.Printf("\nIndex: %v, Value: %v", i, v)
	}

	// for is while
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}
}

func loopingInMaps() {

	fmt.Println("\ninside Loopin in Maps function")
	var myMap map[int]int = map[int]int{1: 1, 2: 2}
	fmt.Printf("\n\nThe map here is: %v", myMap)

	for name := range myMap {
		fmt.Printf("\nName in myMap %v", name)
	}

}

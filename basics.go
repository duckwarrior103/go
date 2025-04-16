// all go programs are made of packages
// Start running in package main
package main 

import (
	"fmt" // fmt package is for formatting input and output
	"math"
	"math/rand"
)
// use "factored" import statement

func add(x int, y int) int {
	return x + y
}
// functions in Go have their type AFTER the variable name e.g. x int
// functions can taken in 0 - many parameter 

func subtract(x, y int) int {
	return x - y
}
// if consecutive parameters share a type, you only have to define type for the last parameter

func returnMany(x, y int) (int, int, int) {
	return x + y, x - y, x * y
}
// Functions can return any number of results - e.g. three ints

func split(sum int) (first, second int) {
	first = sum / 3 
	second = sum - first 
	return // using a naked return means returning the named values, first and second 
}
// Return values can be named, at the top of the  should be used to document the meaning of the return values

func creatingVariables() {
	// you can declare a list of variables like var ... type
	var word1, word2, word3 string
	word1 = "hello"
	word2 = "blue"
	word3 = "planet"

	fmt.Println(word1, word2, word3)

	// you can initialise variables, if so you can skip type, variable will take type of initializer
	var i, j = 1, 2 // i, j takes values 1, 2 and their type which is int 

	fmt.Println(i, j)

	// short variable declarations (only inside functions)
	k := 3
	fmt.Println(k)
}

func main() {
	fmt.Println("Hello, World!", rand.Intn(10))

	// names exported from a package begin with a capital letter 
	fmt.Println(math.Pi)

	fmt.Println(add(12, 37))

	fmt.Println(returnMany(6, 3))

	fmt.Println(split(10))
}
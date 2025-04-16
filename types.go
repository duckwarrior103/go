package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func demoVariables() (imaginary complex128) {
	// boolean
	var x bool = false

	// string - immutable sequence of characters
	var word string = "hello"

	var num int = 4

	var decimal float32 = 2.0

	fmt.Println(x, word, num, decimal)

	imaginary = cmplx.Sqrt(-64)

	// it is recommended to use int unless you have reasons for sized/unsigned ints

	var emptyBoolean bool
	var emptyInteger int
	// variables declared without an explicit initializer will taken their zero values
	// 0 for numeric, false for boolean, "" for empty string
	fmt.Println(emptyBoolean, emptyInteger)

	return
}

func typeConversion() {
	// T(v) will convert value v to the type T
	f := 42.3
	i := int(f)

	fmt.Println(f, i)
}

func typeInference() {
	// like in other languages, precision of the assigned value determines type
	// for example, 42.3 is float
	v := 3000
	fmt.Printf("v: %T\n", v) //note the use of Printf instead of Println, it is for formatted output
	// format specifiers
}

func constantDeclaration() {
	const Pi = math.Pi
	// constants can be boolean, string, character, numeric
	// you cannot declare constants using :=
	fmt.Println(Pi)
}

func main() {
	fmt.Println(demoVariables())
	typeConversion()
	typeInference()
	constantDeclaration()
}

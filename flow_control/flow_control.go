package main

import (
	"fmt"
	"runtime"
	"time"
)

func forLoops() {
	kyan := 11

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// init and post statements are optional
	for ; kyan < 15; kyan++ {
		// do nothing
	}

	// for loop is while loop in C
	for kyan < 20 {
		kyan += 1
	}

	fmt.Println(kyan)
}

func ifStatements() {
	x := 100
	if x < 200 {
		fmt.Printf("%d is a small number\n", x)
	}

	// you can declare variables in if statements like in for loops
	if x := 8; x <= 10 {
		fmt.Printf("%d is <= 10\n", x)
	}
}

func fakeSqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func switches() {
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	// you can declare variables in case, e.g. switch num := 5; num {}
	// switch cases do not have to be constants nor integers
	// switch cases with no condition
	switch i := true; i {
	case time.Now().Hour() < 12:
		fmt.Println("Good Morning!")
	default:
		fmt.Println("You're not an early worm.")
	}

	// that was logically the same as i:=true, so true == case statement? The switch effectively becomes an if-else chain
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Good Morning!")
	default:
		fmt.Println("You're not an early worm.")
	}
}

func main() {
	// forLoops()

	// defers run when the surrounding function is done executing
	defer fmt.Println("Everything else in main() is finished.")
	ifStatements()
	fmt.Printf("Estimated square root of %f is %f\n", 78.0, fakeSqrt(78.0))
	switches()
}

package main

import (
	"fmt"
	"math"
	"strings"
)

func pointers() {
	i := 42
	var p *int = &i // holds the memory address of i
	fmt.Println(*p)

	// dereferencing / indirecting
	*p = 65
	fmt.Println(i) // i value should now be 65
}

func structs() {
	// a struct is a collection of fields
	// declare a type Vertex which is a struct
	type Vertex struct {
		X int
		Y int
	}

	var point Vertex = Vertex{1, 2}
	point.X = 4 // access struct fields using a dot
	fmt.Println(point)

	// we can access struct fields with a pointer to a struct
	point_alias := &point
	(*point_alias).X = 68
	point_alias.X = 70 // this has the same effect as the one above

	// struct literals
	new_point := Vertex{} //X:0 and Y:0
	new_point = Vertex{X: 3, Y: 1}
	fmt.Println(new_point)

	// init pointer to struct literal
	new_pointer := &Vertex{}
	fmt.Println(new_pointer, new_pointer.X)
}

func arrays() {
	var a [10]int // inits a with int array of size 10
	// array length is part of type, it cannot be resized
	fmt.Println(a)
}

func slices() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:3] // Slice the first three elements

	fmt.Println(arr)
	fmt.Printf("Before append: len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	// Append elements within capacity
	slice = append(slice, 6, 7)
	fmt.Println(arr)
	fmt.Printf("After append: len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	// declaring strings with string literal
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q, len(q), cap(q))
	q = append(q, 18, 19, 20)
	fmt.Println(q, len(q), cap(q))

	// creating slices with make
	// creates array with length of 5
	aMadeWithMake := make([]int, 5)
	fmt.Println(aMadeWithMake)
	// Specify capacity with make
	bMadeWithMake := make([]int, 0, 5) // slice has length 0, but is referencing an array with size 5 (cap 5)
	// underlying array: [0,0,0,0,0], slice: []
	fmt.Println(bMadeWithMake) // should be an empty slice
}

func loopOverSliceOfNumbers() {
	numbers := []int{1, 3, 5, 7, 9}
	for i, v := range numbers {
		fmt.Printf("Element %d is: %d\n", i, v)
	}
}

func Pic(dx, dy int) [][]uint8 {
	thingToReturn := make([][]uint8, dy, dy)
	for y := range thingToReturn {
		row := make([]uint8, dx, dx)
		for x := range thingToReturn[y] {
			thingToReturn[y][x] = uint8((x + y) / 2)
		}
		thingToReturn[y] = row
	}
	return thingToReturn
}

func maps() {
	// we want to create a location to vertex map
	type Vertex struct {
		Lat, Long float64
	}
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)

	// map literals
	m2 := map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408},
	}
	fmt.Println(m2)

	// test key is present
	elem, ok := m2["Apple"]
	fmt.Println(elem, ok) // element is nil of map's element type
	//deleting a key
	delete(m2, "Google")
}

func countWords(s string) map[string]int {
	wordCountMap := make(map[string]int)
	words := strings.Split(s, " ")
	for _, word := range words {
		wordCountMap[word]++
	}
	return wordCountMap
}

func compute3And4(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute3And4(hypot))
}

// what is a closure?
// A closure is a "function value e.g. hypot := from above" which references a variable outside of func body

func adder() func(int) int {
	sum := 1
	// returning a function value: closure, a function that takes in int and returns int
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	current := 0
	next := 0
	return func() int {
		if next == 0 {
			next = 1
			return current
		}
		temp := current
		// get  next value in line and set it current
		current = next
		// compute the new next value in line
		next = temp + next
		return current
	}
}

func main() {
	// fmt.Println("Hello, World!")
	// pointers()
	// structs()
	// arrays()
	// slices()
	// loopOverSliceOfNumbers()
	// pic.Show(Pic)
	// maps()
	// fmt.Println(countWords("I ate a donut. Then I ate another donut."))
	// functionValues()
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

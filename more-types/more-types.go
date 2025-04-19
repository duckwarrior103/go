package main

import "fmt"

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

func main() {
	fmt.Println("Hello, World!")
	pointers()
	structs()
	arrays()
}

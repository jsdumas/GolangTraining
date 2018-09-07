package main

import (
	"fmt"
)

func main() {

	a := 43
	
	//print value of a
	fmt.Println(a)
	//print memory address of a
	fmt.Println(&a)

	//store memory address of a in varirable b
	var b = &a

	//print memory address of a
	fmt.Println(b)

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
}

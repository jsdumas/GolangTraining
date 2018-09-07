package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	
	//*b allows to print value which is stored in memory address of b
	fmt.Println(*b) // 43

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
}

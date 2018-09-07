package main

import "fmt"

//param func type is set up to receive memory adress with *int
func zero(z *int) {
	//print memory adress srored in param z
	fmt.Println(z)
	//change value stored in memory adress of parameter z
	*z = 0
}

func main() {
	x := 5
	//print memory address of x
	fmt.Println(&x)
	//pass memory address of x in parameter
	zero(&x)
	//print value of x
	fmt.Println(x) // x is 0
}

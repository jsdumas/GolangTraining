package main

import "fmt"

//param func type is set up to receive memory adress with *int
func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	//func is passed memory address in param
	zero(&x)
	fmt.Println(x) // x is 0
}

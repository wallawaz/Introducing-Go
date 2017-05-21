package main

import "fmt"

func main() {
	//& -> find the address variable
	y := 1001
	makeZero(&y)
	fmt.Println(y)

	//or
	z := new(int)
	makeOne(z)
	fmt.Println(*z)
}

//* siginfies a pointer
func makeZero(xPointer *int) {
	*xPointer = 0
}
func makeOne(xPointer *int) {
	*xPointer = 1
}

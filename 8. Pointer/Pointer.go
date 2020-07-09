package main

import "fmt"

func main() {
	number1 := 6

	pointer := &number1

	fmt.Println("Address number1: ", &number1)
	fmt.Println("Address pointer: ", &pointer)
	fmt.Println("Pointer point  : ", pointer)

	*pointer = 10
	fmt.Println()
	fmt.Println("number1 value      : ", number1)
	fmt.Println("pointer point value: ", *pointer)

}

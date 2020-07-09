package main

import "fmt"

func sum(number1, number2 int) (int, string) { // passing by value
	return number1 + number2, "success!"
}

func double(number *int) { // passing by reference
	*number *= 2
}

func main() {
	fmt.Println("hello", "world")
	total, msg := sum(1, 5)
	fmt.Println(total)
	fmt.Println(msg)

	number := 5
	fmt.Println("Before:", number)
	double(&number)
	fmt.Println("After:", number)

}

package main

import "fmt"

func main() {
	numbers := [10]int{1, 2, 3, 4}                     // array
	strings := []string{"bob", "maya", "john", "tedy"} //slices - dynamic array

	fmt.Println(numbers[1:5])
	strings = strings[1 : len(strings)-1]
	fmt.Println(strings)
}

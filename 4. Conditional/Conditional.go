package main

import "fmt"

func greaterThanZero(number int) bool {
	if number > 0 {
		return true
	} else {
		return false
	}
}

func main() {

	number := 1

	if number%2 == 0 {
		fmt.Println("number is even")
	} else if number%3 == 0 {
		fmt.Println("number is divideable by 3")
	} else {
		fmt.Println("number is odd")
	}

	switch number {
	case 1:
		fmt.Println("number is one")
	case 2:
		fmt.Println("number is two")
		// fallthrough - continue to next case
	default:
		fmt.Println("number is unknown")
	}

	fmt.Println(greaterThanZero(2))

}

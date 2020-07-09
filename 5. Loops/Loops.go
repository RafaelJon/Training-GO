package main

import "fmt"

func main() {
	// for loop ori
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	number := 0
	// while loop example
	for number < 5 {
		fmt.Println(number)
		number++
	}

	number2 := 0
	// do while example
	for {
		fmt.Println(number2)
		number2++
		if number2 > 15 {
			break
		}
	}
}

package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, value := range number[2:6] {
		fmt.Println("Value:", value)
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello world!")
	/*
		go run - compile and execute, no exe file
		go install - compile and execute, exe file in bin folder
		go build - compile and execute, exe file in the same folder
	*/

	/*
		int
		int8, int16, int32, int64
		float32, float64
		string
		bool
		uint

		var [variable_name] [type] = value
		var [variable_name] = value
		[variable_name] := value
	*/

	// number1 := 45
	// number2 := 32
	// fmt.Println(number1 + number2)
	// fmt.Printf("Type of %d : %T\n", number1, number1)

	// number3 := "120"
	// number4, _ := strconv.Atoi(number3)
	// fmt.Printf("Type of %d : %T\n", number4, number4)

	number5 := math.Ceil(45.7)
	number6 := int(number5)
	fmt.Printf("Type of %d : %T\n", number6, number6)

}

package main

import "fmt"

func main() {

	// dict["key"] : 10
	maps := make(map[string]string)
	maps["hello"] = "hi"
	maps["world"] = "country"

	fmt.Println(maps["hello"])
}

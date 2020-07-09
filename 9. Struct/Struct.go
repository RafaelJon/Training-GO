package main

import "fmt"

type computer struct {
	RAM        int
	Battery    float32
	Proccessor string
	GPU        string
}

func (c computer) printData() {
	fmt.Println("PC spec:")
	fmt.Println("PC RAM       :", c.RAM, "GB")
	fmt.Println("PC Battery   :", c.Battery, "wh")
	fmt.Println("PC Proccessor:", c.Proccessor)
	fmt.Println("PC GPU       :", c.GPU)
}

func (c *computer) upgradeRAM(ram int) {
	c.RAM += ram
	fmt.Println("PC RAM       :", c.RAM, "GB")
}

func main() {
	// pc := computer{
	// 	RAM:        8,
	// 	Battery:    64.5,
	// 	Proccessor: "I7 7700k",
	// }

	pc := computer{8, 64.5, "I7 7700K", "gtx 1050"}
	pc.upgradeRAM(8)
	fmt.Println()
	pc.printData()
}

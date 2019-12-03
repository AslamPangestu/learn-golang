package main

import (
	"fmt"
)

// Car -> Struct = Data Encapusulation.
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

// Print -> Define Method For Struct.
func (c Car) Print() {
	fmt.Println(c)
}

// Drive -> Define Method For Struct.
func (c Car) Drive() {
	fmt.Println("Drive")
}

// GetName -> Define Method For Struct.
func (c Car) GetName() string {
	return c.Name
}

func main() {
	//Initialize Struct
	// c := Car{"Test", 1, 12}
	c := Car{
		Age:     1,
		ModelNo: 12,
		Name:    "Test",
	}
	// var c1 Car
	fmt.Println(c)
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}

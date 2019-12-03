package main

import (
	"fmt"
)

// Mobil Declare interface
type Mobil interface {
	Drive()
	Stop()
}

// NewModel use Mobile Interfacce
func NewModel(arg string) Mobil {
	return &Lambo{arg}
}

// Lambo Car struct
type Lambo struct {
	Model string
}

// Chevy Car struct
type Chevy struct {
	Model string
}

// Drive Method Lambo
func (l *Lambo) Drive() {
	fmt.Println("Lambo Move")
	fmt.Println(l.Model)
}

// Drive Method Chevy
func (c *Chevy) Drive() {
	fmt.Println("Chevy Move")
	fmt.Println(c.Model)
}

// Stop Method Lambo
func (l *Lambo) Stop() {
	fmt.Println("Lambo Stop")
	fmt.Println(l.Model)
}

// Stop Method Chevy
func (c *Chevy) Stop() {
	fmt.Println("Chevy Stop")
	fmt.Println(c.Model)
}

func main() {
	l := NewModel("Test")
	c := Chevy{"Test1"}
	l.Drive()
	c.Drive()
}

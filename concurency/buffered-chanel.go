package main

import (
	"fmt"
)

//Car Type
type Car struct {
	Model string
}

func main() {
	// var c chan int

	c := make(chan *Car, 3)
	// c := make(chan int, 3)

	//Send
	go func() {
		c <- &Car{"1"}
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"}
		close(c)
	}()

	//Sniff
	for i := range c {
		fmt.Println(i.Model)
	}
}

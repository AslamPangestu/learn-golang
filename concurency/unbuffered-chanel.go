package main

import (
	"fmt"
	"time"
)

func main() {
	// var c chan int

	c := make(chan int)

	//Send
	go func() {
		c <- 1
	}()
	//Sniff
	val := <-c
	fmt.Println(val)

	//Send
	go func() {
		c <- 2
	}()
	time.Sleep(time.Second * 2)
	//Sniff
	val = <-c
	fmt.Println(val)
}

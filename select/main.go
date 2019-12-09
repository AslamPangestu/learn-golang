package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, q chan struct{}) {
	for {
		time.Sleep(time.Second)
		//Select Chanel
		select {
		case <-c:
			fmt.Println("Received")
		case <-q:
			fmt.Println("Quit")
			os.Exit(0)
			// break
		}
	}
}

func main() {
	c := make(chan int, 2)
	q := make(chan struct{})

	go func() {
		c <- 1
		// q <- struct{}{}
	}()
	go Select(c, q)
	//Make Infinite
	select {}
}

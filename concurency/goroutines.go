package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("hai")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("hai hai")
	}
}

func main() {
	//Declare Goroutine -> Like async func & run in background
	go heavy()
	go superHeavy()
	fmt.Println("lalal")
	// time.Sleep(time.Second * 5)
	select {}
}

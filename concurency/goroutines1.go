package main

import (
	"fmt"
	"sync"
)

func main() {
	//Wait Groups
	wg := &sync.WaitGroup{}
	//Add, done & Wait

	//Mutex
	// mute := &sync.Mutex{}

	wg.Add(2)
	go func() {
		fmt.Println("Hai")
		//Done Signal
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		//Done Signal
		wg.Done()
	}()
	fmt.Println("Budi")
	wg.Wait()
	fmt.Println("Done")
}

package main

import (
	"fmt"
)

//Create Func
func todo() {
	//Declare Array
	// var arr []int
	// arr[0] = 1
	// arr[1] = 2
	// arr[3] = 3
	// arr[4] = 4
	arr := []int{1, 2, 3, 4}
	arr = append(arr, 5, 6)

	fmt.Println(arr)
}

func main() {
	//Call Function
	todo()
}

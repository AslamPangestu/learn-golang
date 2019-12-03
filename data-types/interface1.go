package main

import (
	"fmt"
)

// interface{} -> Empty Interface

// Anything => Func do anything
func Anything(anything interface{}) {
	fmt.Println(anything)
}
func main() {
	fmt.Println("vim-go")
	Anything("HAI")
	Anything(1)
	Anything(struct{}{})

	// struct{}{} -> Empty Struct
	//Declar Struct
	// struct{ Person string }{"hi"}

	mymap := make(map[string]interface{})

	mymap["name"] = "LALA"
	mymap["age"] = 10
	fmt.Println(mymap)
}

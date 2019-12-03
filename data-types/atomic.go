package main

import (
	"fmt"
	"strings"
)

func main() {
	//Declare Var
	var m1 int
	var s1 string
	m1 = 2
	s1 = "Budi"
	fmt.Println(m1)
	fmt.Println(s1)

	//Declare Many Var
	var (
		m2 = 3
		m3 = 3
	)
	fmt.Println(m2 + m3)

	var m4 int32
	var m5 int64
	//Casting
	fmt.Println(int64(m4) + m5)

	//Declare Simple Var
	m6 := 2
	m7 := 8
	s6 := "Dina"
	fmt.Println(m6 + m7)
	fmt.Println(s6)

	//Compare String
	s2 := "My Name"
	s3 := "Name"
	fmt.Println(strings.Contains(s2, s3))          //Find Text
	fmt.Println(strings.ReplaceAll(s2, "m", "LA")) //Replace All String
	fmt.Println(strings.Split(s2, " "))            // Split by space
	fmt.Println(s2 + s3)                           // Add String
}

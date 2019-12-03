package main

import (
	"fmt"
)

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func main() {
	m1 := 2
	ptr := &m1  //Get Value Address
	val := *ptr //Get Value
	fmt.Println(val)

	//Declar multiple init
	m2, m3 := 2, 3

	fmt.Println(m2, m3)
	swap(&m2, &m3)
	fmt.Println(m2, m3)
}

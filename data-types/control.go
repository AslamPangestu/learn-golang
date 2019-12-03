package main

import (
	"fmt"
)

func main() {
	//if else, for
	f := true
	flag := &f

	//Conditional
	if flag == nil {
		fmt.Println("Null")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	arr := []string{"Budi", "Anduk"}
	//For Loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	//Infinite Loop
	// for {
	// 	fmt.Println("Hai")
	// }
	//Loop Array
	for i, item := range arr {
		fmt.Println(i)
		fmt.Println(item)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "Budi"
	mymap["age"] = 20

	//Loop Map
	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v", k, v)
	}

	//continue, break, switch
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	flag1 := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag1 = false
			break
		} else if i == 1 {
			continue

		}
		fmt.Println(i)
	}
	fmt.Println(flag1)

	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("TGIF")
	case "Mon", "Tue", "Wed":
		fmt.Println("Boring")
	default:
		fmt.Println("Default")
	}

	switch {
	case day == "Fri":
		fmt.Println("TGIF")
		break
	}
}

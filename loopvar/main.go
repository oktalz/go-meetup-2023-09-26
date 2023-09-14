package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var list []int
	for i := 1; i <= 3; i++ {
		list = append(list, i)
	}
	for _, item := range list {
		fmt.Println(item)
	}

	fmt.Println(" ")

	var prints []func()
	for i := 1; i <= 3; i++ {
		prints = append(prints, func() { // HL
			fmt.Println(i) // HL
		}) // HL
	}
	for _, print := range prints {
		print()
	}
	// END OMIT
}

// RUN OMIT
// NUR OMIT

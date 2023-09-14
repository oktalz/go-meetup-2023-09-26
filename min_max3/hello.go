package main

import "fmt"

func main() { // HL
	// START OMIT
	var x, y = 61, 42

	// RUN OMIT
	m1 := min(x, y, 100) // HL
	m2 := max(x, y, 100) // HL
	// NUR OMIT

	fmt.Println(m1)
	fmt.Println(m2)
	// END OMIT
} // HL

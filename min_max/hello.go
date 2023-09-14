package main

import "fmt"

func main() {
	// START OMIT
	var x, y = 61, 42

	// RUN1 OMIT
	m := min(x)    // HL
	n := min(x, y) // HL
	o := max(x, y) // HL
	// NUR1 OMIT

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(min(x, y) == min(y, x))
	//END OMIT
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	// START OMIT
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	i := slices.Index(numbers, 0)    // HL
	j := slices.Index(numbers, 5)    // HL
	k := slices.Index(numbers, 1000) // HL

	fmt.Println(i, j, k)
	// END OMIT
}

// RUN OMIT
// NUR OMIT

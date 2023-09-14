package main

import (
	"fmt"
	"slices"
)

func main() {
	// START OMIT
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbers = slices.Delete(numbers, 1, 2) // HL

	fmt.Println(numbers)
	// END OMIT
}

// RUN OMIT
// NUR OMIT

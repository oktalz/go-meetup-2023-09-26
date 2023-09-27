package main

import (
	"fmt"
)

// START OMIT
func sum(x int) {

	// YES, the code is not optimized

	total := 0
	for i := 1; i <= x; i++ {
		total += i
	}
	fmt.Printf("The sum of the first %d numbers is: %d\n", x, total)
}

func main() {
	sum(10)
}

// END OMIT

package main

import (
	"fmt"
	"maps"
)

func main() {
	// START OMIT
	m1 := map[int]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f"}

	m2 := maps.Clone(m1)

	fmt.Println(maps.Equal(m1, m2))
	fmt.Println(m2)
	// END OMIT
}

// RUN OMIT
// NUR OMIT

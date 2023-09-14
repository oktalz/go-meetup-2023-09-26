package main

import (
	"fmt"
	"maps"
)

func main() {
	// START OMIT
	m1 := map[int]string{0: "x", 6: "g"}
	m2 := map[int]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f"}

	maps.Copy(m1, m2)

	fmt.Println(m1)
	// END OMIT
}

// RUN OMIT
// NUR OMIT

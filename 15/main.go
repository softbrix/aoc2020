package main

import (
		"fmt"
)
 
func main() {
	const MAX = 30000000
	in := []int{14,1,17,0,3,20}
	dict := make(map[int]int, MAX)
	i := 0
	last := 0
	for t := range in {
		dict[t] = i
		i++
	}
	for ; i < MAX; i++ {
		prev_idx, found := dict[last]
		if !found {
      dict[last] = i
      last = 0
    } else {
      lNew := i - prev_idx
      dict[last] = i
      last = lNew
    }
    if ((i % 1000000) == 0) {
      fmt.Print(".")
    }
    if (i == 2019) {
			fmt.Printf("Part 1: %d\n", last)
    }
	}

	fmt.Printf("\nPart 2: %d\n", last)
}
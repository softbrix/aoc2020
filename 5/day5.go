package main

import (
    "bufio"
		"fmt"
    "os"
)

func missing(items []int) int {
	for i := 0; i < 8; i++ {
		found := false
		for _,x := range items {
			found = found || x == i
		}
		if (!found) {
			return i;
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sorted := make(map[int][]int)
	rows := make([]int, 0)
	minR := 200
  maxR := -1
  for scanner.Scan() {
		//l := scanner.Text()
		r := 0
		rd := 64
		s := 0
		sd := 4
		
    for _,c := range scanner.Text() {
			if (c == 'B') {
				r += rd
			}
			rd /= 2
			if (c == 'R') {
				s += sd
				sd /= 2
			} else if (c == 'L') {
				sd /= 2
			}
		}

		rows = append(rows, r * 8 + s)
		sorted[r] = append(sorted[r], s)
		if (minR > r) {
			minR = r
		}
		if (maxR < r) {
			maxR = r
		}
	}
	max := 0
  for _,id := range rows {
		if (id > max) {
			max = id
		} 
	}

	for i,row := range sorted {
		if (len(row) == 7 && minR != i && maxR != i) {
			s := missing(row)
		  fmt.Printf("Part2: %d\n", i*8 + s)
		}
	}

	fmt.Printf("Part1: %d\n", max)
}
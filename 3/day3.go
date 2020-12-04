package main

import (
    "bufio"
		"fmt"
    "os"
)


func B2i(b bool) int {
	if b {
			return 1
	}
	return 0
}
 
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rs := make([][]bool, 0)
	rows := 0;
	cols := 0
  for scanner.Scan() {
		l := scanner.Text()
		cols = len(l)
		row := make([]bool, cols)
		for i,c := range l {
			row[i] = string(c) == "#"
		}
		rs = append(rs, row)
		rows++
	}
	

	if scanner.Err() != nil {
		// handle error.
	}
	

//	Right 1, down 1.
//	Right 3, down 1. (This is the slope you already checked.)
//	Right 5, down 1.
//	Right 7, down 1.
//	Right 1, down 2.

	dxs := []int{1,3,5,7,1}
	dys := []int{1,1,1,1,2}

  tot := 1
	for j := 0; j < len(dxs); j++ {
		dx := dxs[j]
		dy := dys[j]
		sum1 := 0
		x := 0
		for y := 0; y < rows; y += dy {
			sum1 += B2i(rs[y][x%cols])
			x += dx
		}
		if (dx == 3 && dy == 1) {	
			fmt.Printf("Sum1: %d\n", sum1)
		}
		tot *= sum1
	}
	fmt.Printf("Sum2: %d\n", tot)
}
package main

import (
    "bufio"
		"fmt"
		"strings"
		"strconv"
    "os"
)

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}


const ONE = int64(1)
func pow2(i int) int64 {
	return ONE << i
}

func reverse(str string) (result string) { 
	for _, v := range str { 
			result = string(v) + result 
	} 
	return
} 

func indices(newidx int64, floating []int) []int64 {
	if len(floating) == 0 {
		return []int64{newidx}
	} else {
		b0 := floating[len(floating)-1]
		floating = floating[:len(floating)-1]
		return append(indices(newidx, floating), indices(newidx + pow2(b0), floating)...)
	}
}

func main() {
	sum1 := int64(0)
	sum2 := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	ma1 := make(map[int64]int64)
	ma2 := make(map[int64]int64)
	var msk string
  for scanner.Scan() {
		l := scanner.Text()
		v := strings.Split(l, " = ")
		if v[0] == "mask" {
			msk = v[1]
			continue
		}
		idx,_ := strconv.ParseInt(strings.Split(v[0][4:], "]")[0], 10, 64)
		value,_ := strconv.ParseInt(v[1], 10, 64)

		newvalue := int64(0)
		newidx := int64(0)
		floating := make([]int, 0, 0)
		for i,c := range reverse(msk) {
			ibit := idx & pow2(i)
			vbit := value & pow2(i)
			if c == 'X' {
				floating = append(floating, i)
				newvalue += vbit
			} else if c == '1' {
				newidx += pow2(i)
				newvalue += pow2(i)
			} else if c == '0' {
				newidx += ibit
			}
		}
		for _,p := range indices(newidx, floating) {
			ma2[p] = value
		}
		ma1[idx] = newvalue
	}

	for _,value := range ma1 {
		sum1 += value
	}
	for _,value := range ma2 {
		sum2 += value
	}
	
	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}
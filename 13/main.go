package main

import (
    "bufio"
		"fmt"
		"strings"
		"strconv"
		"math"
    "os"
)

type pair struct {
	v int64
	m int64
	used bool
}
func (u *pair) Used() {
  u.used = true
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func mul(pairs []pair) int64 {
	s := int64(1)
	for _,t := range pairs {
		s *= t.m
	}
	return s
}

func main() {
	sum1 := 0
	sum2 := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
	tStart,_ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	times := strings.Split(scanner.Text(), ",")

	pairs := make([]pair, 0, 0)
	for j,t := range times {
		if t == "x" {
			continue;
		}
		x,err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			panic(err)
		}
		pairs = append(pairs, pair{v:int64(j), m:x})
	}
	fmt.Printf("v: %v\n", pairs) 

	maxT := mul(pairs)

	max := &pair{m:0}
	for _,t := range pairs {
		if max.m < t.m {
			max = &t
		}
	}
	max.Used()
	start := max.m - max.v
	fmt.Printf("Max: %v %v\n", *max, start)

	inc := max.m
	found := false
	for a := start; !found && a < maxT; a += inc{
		b := false

		for i,t := range pairs {
			if (a + t.v ) % t.m != 0 {
				b = true
				break
			}
			if !t.used && (a + t.v ) % t.m == 0 {
				inc *= t.m
				pairs[i].Used()
				fmt.Printf("Second: %v %d %d\n", t, inc, a)
			}
		}

		if !b {
			sum2 = a
			found = true;
			for _,t := range pairs {
				fmt.Printf("Q: %d %d %d\n", a, a % t.m, t.v)
			}
		}
		b = false 
	} 
	/*
	Part 1 */
	minV := math.MaxInt32
	for _,t := range times {
		x,err := strconv.Atoi(t)
		if err != nil {
			continue;
		}
		v := ((tStart / x) + 1) * x
		if v < minV {
			minV = v
			sum1 = (minV - tStart ) * x
		}
	}
	
	fmt.Printf("Sum1: %d %d \n", sum1, tStart)
	fmt.Printf("Sum2: %d\n", sum2)
}
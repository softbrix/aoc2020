package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func err(e error) {
	if e != nil {
		panic(e)
	}
}

var dp map[int][]string = make(map[int][]string)

func contains(s string, arr []string) bool {
	for _, e := range arr {
		if e == s {
			return true
		}
	}
	return false
}

func makeMatches(rules []string, idx int) []string {
	if dp[idx] != nil {
		return dp[idx]
	}

	v := rules[idx]
	if v == "\"a\"" {
		return []string{"a"}
	}
	if v == "\"b\"" {
		return []string{"b"}
	}

	res := make([]string, 0, 0)
	s := strings.Split(v, " | ")
	for _, c := range s {
		p := strings.Split(c, " ")
		var opt []string
		for _, pv := range p {
			a, _ := strconv.Atoi(pv)
			match := makeMatches(rules, a)
			if opt == nil {
				opt = match
			} else {
				optnew := make([]string, 0, 0)
				for _, o := range opt {
					for _, m := range match {
						optnew = append(optnew, o+m)
					}
				}
				opt = optnew
			}
		}
		res = append(res, opt...)
	}
	dp[idx] = res
	return res
}

func main() {
	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	rules := make([]string, 150, 150)
	for scanner.Scan() {
		r := scanner.Text()
		if len(r) == 0 {
			break
		}
		//fmt.Println(r)
		v := strings.Split(r, ": ")
		i, _ := strconv.Atoi(v[0])
		rules[i] = v[1]
	}

	fmt.Printf("0: %v\n", rules[0])

	matches := makeMatches(rules, 0)

	valid := make(map[string]bool)

	for _, m := range matches {
		valid[m] = true
	}

	rows := make([]string, 0, 0)

	for scanner.Scan() {
		r := scanner.Text()
		if valid[r] {
			sum1++
		}
		rows = append(rows, r)
	}

	// Part2, 42 is all combinations
	// rules[0] = 8 11
	// rules[8] = "42 | 42 8"
	// rules[11] = "42 31 | 42 11 31"

	//	fmt.Printf("%v\n%v\n", len(dp[31][0]), dp[31])
	//	fmt.Printf("%v\n%v\n", len(dp[42][0]), dp[42])

	if len(dp[42][0]) != len(dp[31][0]) {
		panic("Expected equal block length")
	}

	blkLen := len(dp[42][0])
	for _, r := range rows {
		if ((len(r) - blkLen) % blkLen) == 0 {
			left := r
			i42 := 0
			for ; len(left) > 0; left = left[blkLen:] {
				if !contains(left[:blkLen], dp[42]) {
					break
				}
				i42++
			}
			j31 := 0
			for ; len(left) > 0; left = left[blkLen:] {
				if !contains(left[:blkLen], dp[31]) {
					break
				}
				j31++
			}
			if len(left) == 0 && j31 > 0 && i42 > j31 {
				sum2++
			}
		}
	}

	if scanner.Err() != nil {
		// handle error.
		panic(scanner.Err())
	}

	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}

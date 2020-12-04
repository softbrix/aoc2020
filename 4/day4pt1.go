package main

import (
    "bufio"
		"fmt"
		"strings"
		"regexp"
		"strconv"
    "os"
)

func mkMap() map[string]matcher {
	flds := [...]string{"byr", "iyr", "eyr", "hgt","hcl", "ecl","pid"} // "cid"
	for _,d := range flds {
		m[d] = true
	}
	return m
} 
 
func main() {
	sum1 := 0
	fldi := mkMap()
	scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		if (len(l) == 0) {
			// reset
			fldi = mkMap()
		} else {
			for _,d := range l {
				k := strings.Split(d, ":")[0]
				delete(fldi, k)
			}
			fmt.Println(len(fldi))
			if (len(fldi) == 0) {
				sum1 ++
				fldi = mkMap()
			}
		}
  }

    if scanner.Err() != nil {
      // handle error.
		}
		fmt.Printf("Sum1: %d\n", sum1)
  }
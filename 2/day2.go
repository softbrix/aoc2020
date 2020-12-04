package main

import (
    "bufio"
		"fmt"
		"strings"
		"strconv"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}  
 
func main() {

	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		lo,_ := strconv.Atoi(strings.Split(l[0], "-")[0])
		hi,_ := strconv.Atoi(strings.Split(l[0], "-")[1])
		c := l[1][0:1]
		chars := l[2]
		tot := 0
		for _,d := range chars {
			if string(d) == c {
				tot++
			}
		}
		if (lo <= tot && tot <= hi) {
			sum1 += 1
		}
		if ((string(chars[lo-1]) == c) != (string(chars[hi-1]) == c)) {
			sum2 += 1
		}
  }

    if scanner.Err() != nil {
      // handle error.
		}
		fmt.Printf("Sum1: %d\n", sum1)
		fmt.Printf("Sum2: %d", sum2)
  }
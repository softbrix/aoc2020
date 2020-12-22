package main

import (
    "bufio"
		"fmt"
		"sort"
		"strconv"
    "os"
)

func find(n int, nums []int) bool {
	for _,a := range nums {
		for _,b := range nums {
			if (a+b == n && a != b) {
				return true
			}
		}
	}
	return false
}

func sum(i int,j int, nums []int) int {
	s := 0
	for ; i <= j; i++ {
		s += nums[i]
	}
	return s
}
 
func main() {
	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	nums := make([]int, 0, 0)
	set := make([]int, 0, 0)
	c := 0
	const p = 25 // demo = 5, prod = 25
  for scanner.Scan() {
		l := scanner.Text()
		n,_ := strconv.Atoi(l)
		if (c < p) {
			set = append(set, n)
		} else {
			if !find(n, set) {
				sum1 = n
				break;
			}
			set[c%p] = n
		}
		nums = append(nums, n)
		c++
	}

	for i := len(nums)-1; i >= 0; i-- {
		for j := i-1; j >= 0; j-- {
			sum := sum(j,i,nums)
			if sum == sum1 {
				v := nums[j:i]
				sort.Ints(v)
				sum2 = v[0] + v[i-j-1]
				i = 0
				break
			} else if sum > sum1 {
				break
			}
		}
  }

	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}
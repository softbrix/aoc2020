package main

import (
    "bufio"
		"fmt"
		"sort"
		"strconv"
    "os"
)

func find(n int, nums []int) int {
	for i,a := range nums {
		if (a == n) {
			return i
		}
	}
	return -1
}

func sum(i int,j int, nums []int) int {
	s := 0
	for ; i <= j; i++ {
		s += nums[i]
	}
	return s
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func part2(nums []int) int {
	nums = append([]int{0}, nums...)
	adds := make([]int, len(nums))
	adds[len(nums)-1] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 1; j <= 3; j++ {
			 e := min(i+4, len(nums))
			 id := find(nums[i] + j, nums[i:e])
			 if (id >= 0) {
				adds[i] += adds[i+id] 
			 }
		}
	}
	return adds[0]
}
 
func part1(nums []int) []int {
	adds := []int{0,0,0,0}
	for i := 1; i < len(nums) - 1; i++ {
		adds[nums[i+1] - nums[i]]++
	}
	return adds
}
 
func main() {
	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	nums := make([]int, 0, 0)
  for scanner.Scan() {
		l := scanner.Text()
		n,_ := strconv.Atoi(l)
		nums = append(nums, n)
	}
	


	nums = append(nums, 0)
	sort.Ints(nums)
	nums = append(nums, nums[len(nums) - 1] + 3)

	adds := part1(nums)
	sum1 = adds[1] * adds[3]

	sum2 = part2(nums)

	fmt.Printf("Sum1: %d %d\n", sum1, adds)
	fmt.Printf("Sum2: %d\n", sum2)
}
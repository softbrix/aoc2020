package main

import (
    "bufio"
		"fmt"
//		"sort"
//		"strconv"
    "os"
)

// LANE -1
// FREE 0
// OCCU >0

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

func getSeat(occupants [][]int, di int, dj int) int {
	if (di < 0 || dj < 0) {
		return 0;
	}
	if (di >= len(occupants) || dj >= len(occupants[di])) {
		return 0
	}
	return occupants[di][dj]
}

func adjacent(occupants [][]int, i int, j int) (int,int) {
	busy := 0
	busyFar := 0
	for id := -1; id <= 1; id++ {
		for jd := -1; jd <= 1; jd++ {
			if id == 0 && jd == 0 {
				continue
			}
			di := i + id;
			dj := j + jd;

			seat := getSeat(occupants, di, dj)
			
			if (seat > 0) {
				busy++
			} else {
				a := di
				b := dj
				seat := -1
				for ; seat < 0 && a >= 0 && b >= 0 && a < len(occupants) && b < len(occupants[a]); {
					seat = getSeat(occupants, a, b)
					a += id
					b += jd
				}
				if (i == 0 && j == 0 && di == 0 && dj == 1) {
					fmt.Printf("C3: %d %d %d %d\n", id, jd, seat)
				}
				if seat > 0 {
					busyFar++
				}
			}
			if (i == 0 && j == 0) {
				fmt.Printf("C2: %d %d %d %d\n", id, jd, busy, busyFar)
			}
		}
	}
	return busy, busyFar
}

func occupy2(occupants [][]int, old int, i int, j int) int {
	if old < 0 {
		return -1
	}
	busy,busyFar := adjacent(occupants, i, j)
	if i == 0 {
		fmt.Printf("C1: %d %d %d %d\n", j, old, busy, busyFar)
	}

	busy += busyFar

	if (old == 0 && busy == 0) {
		return 1;
	} else if (old > 0 && busy >= 5) {
		return 0
	}
	return old
}

func occupy(occupants [][]int, old int, i int, j int) int {
	if old < 0 {
		return -1
	}
	busy,_ := adjacent(occupants, i, j)
	if i == -1 {
		fmt.Printf("C1: %d %d %d\n", j, old, busy)
	}

	if (old == 0 && busy == 0) {
		return 1;
	} else if (old > 0 && busy >= 4) {
		return 0
	}
	return old
}
 
func part1(oldOcu [][]int, newOccu [][]int) [][]int {
	for i,row := range oldOcu {
		for j,seat := range row {
			newOccu[i][j] = occupy(oldOcu, seat, i, j)
		}
	}
	//fmt.Printf("ABC %v\n", newOccu)
	//print(newOccu)
	
	return newOccu
}

func part2(oldOcu [][]int, newOccu [][]int) [][]int {
	for i,row := range oldOcu {
		for j,seat := range row {
			newOccu[i][j] = occupy2(oldOcu, seat, i, j)
		}
	}
	//fmt.Printf("ABC %v\n", newOccu)
	print(newOccu)
	
	return newOccu
}

func count(newOccu [][]int) int {
	cnt := 0
	for _,row := range newOccu {
		for _,seat := range row {
			if seat > 0 {
				cnt++
			}
		}
	}
	return cnt
}

func print(oldOcu [][]int) {
	fmt.Println("--------")
	for _,row := range oldOcu {
		runes := make([]rune, 0,0)
		for _,seat := range row {
			runes = append(runes, getRune(seat))
		}
		fmt.Println(string(runes))
	}
	fmt.Println("--------")
}

func getVal(s rune) int {
	if s == '.' {
		return -1
	}
	return 0
}

func getRune(s int) rune {
	if s < 0 {
		return '.'
	} else if s == 0 {
		return 'L'
	}
	return '#'
}

func makeOccupants(seats [][]rune) [][]int {
	occupants := make([][]int, len(seats))
	for i,row := range seats {
		occupants[i] = make([]int, len(row))
		for j,s := range row {
			occupants[i][j] = getVal(s)
		}
	}
	return occupants
}
 
func main() {
	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	seats := make([][]rune, 0, 0)
  for scanner.Scan() {
		l := scanner.Text()
		n := []rune(l)
		seats = append(seats, n)
	}

	fmt.Printf("It: %d\n", seats)

	occupants := makeOccupants(seats)

	last := -1
	res := -2
	
	for i := 0; i < 2048 && last != res; i++ {
		last = res
		occupants = part2(occupants, makeOccupants(seats))
		res = count(occupants); 
		fmt.Printf("It: %d, %d %d\n", i, last, res)
	}
	sum1 = res

	//sum2 = part2(nums)

	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}
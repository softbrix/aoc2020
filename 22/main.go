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

func sum(final []int) int {
	sum1 := 0
	for i, e := range final {
		sum1 += e * (len(final) - i)
	}
	return sum1
}

func join(arr []int) string {
	return strings.Join(strings.Fields(fmt.Sprint(arr)), ",")
}

func clone(arr []int, l int) []int {
	if len(arr) < l {
		l = len(arr)
	}
	c2 := make([]int, l, l)
	for i := 0; i < l; i++ {
		c2[i] = arr[i]
	}
	return c2
}

func play(cards map[int][]int) (int, []int) {
	if len(cards[0]) == 0 {
		return 1, cards[1]
	}
	if len(cards[1]) == 0 {
		return 0, cards[0]
	}
	v0 := cards[0][0]
	v1 := cards[1][0]
	cards[0] = cards[0][1:]
	cards[1] = cards[1][1:]
	if v0 > v1 {
		cards[0] = append(cards[0], []int{v0, v1}...)
	} else {
		cards[1] = append(cards[1], []int{v1, v0}...)
	}
	return play(cards)
}

func play2(cards map[int][]int, game int) (int, []int) {
	hands := make(map[string]bool)
	for i := 1; len(cards[0]) > 0 && len(cards[1]) > 0; i++ {

		s0 := "1:" + join(cards[0])
		s1 := "2:" + join(cards[1])
		key := s0 + s1
		if hands[key] {
			return 0, cards[0]
		}
		hands[key] = true

		v0 := cards[0][0]
		v1 := cards[1][0]
		cards[0] = cards[0][1:]
		cards[1] = cards[1][1:]
		winner := 0
		if v1 > v0 {
			winner = 1
		}
		if len(cards[0]) >= v0 && len(cards[1]) >= v1 {
			// recurse play
			c2 := make(map[int][]int)
			c2[0] = clone(cards[0], v0)
			c2[1] = clone(cards[1], v1)
			game++
			winner, _ = play2(c2, game)
		}
		if winner == 0 {
			cards[0] = append(cards[0], []int{v0, v1}...)
		} else {
			cards[1] = append(cards[1], []int{v1, v0}...)
		}
	}
	if len(cards[1]) > 0 {
		return 1, cards[1]
	} else {
		return 0, cards[0]
	}
}

func main() {
	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	cards := make(map[int][]int)
	cards[0] = make([]int, 0, 0)
	for scanner.Scan() {
		r := scanner.Text()
		if len(r) == 0 {
			i++
			cards[i] = make([]int, 0, 0)
			continue
		}
		if r[0] == 'P' {
			continue
		}
		v, _ := strconv.Atoi(r)
		cards[i] = append(cards[i], v)
	}

	cards2 := make(map[int][]int)
	for k, v := range cards {
		cards2[k] = v
	}

	_, final := play(cards)
	sum1 = sum(final)

	_, final2 := play2(cards2, 1)
	sum2 = sum(final2)

	if scanner.Err() != nil {
		// handle error.
		panic(scanner.Err())
	}

	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}

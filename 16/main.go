package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	min int
	max int
}

type class struct {
	spans []pair
	desc  string
}

func newticket(line string, max int) []int {
	t := []int{}
	for i, s := range strings.Split(line, ",") {
		v, _ := strconv.Atoi(s)
		t = append(t, v)
		if i >= max {
			return t
		}
	}
	return t
}

func isValidVal(v int, spans []pair) bool {
	for _, sp := range spans {
		//fmt.Printf("%d %d %v\n", v, sp, v < sp.min || v > sp.max)
		if v >= sp.min && v <= sp.max {
			return true
		}
	}
	return false
}

func isValidTicket(t []int, spans []pair) bool {
	for _, v := range t {
		if !isValidVal(v, spans) {
			return false
		}
	}
	return true
}

func makeBase(t []int) []bool {
	l := len(t)
	b := make([]bool, l, l)
	for i := 0; i < l; i++ {
		b[i] = true
	}
	return b
}

func findTrue(r []bool) int {
	c := -1
	found := false
	for i, v := range r {
		if v {
			if found {
				return -1
			}
			c = i
			found = true
		}
	}
	return c
}

func main() {
	sum1 := 0
	spans := make([]pair, 0, 0)
	classes := make([]class, 0, 0)
	myTicket := []int{}
	tickets := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		l := strings.Split(line, ": ")
		st := strings.Split(l[1], " or ")
		pairs := []pair{}
		for _, s := range st {
			v := strings.Split(s, "-")
			min, _ := strconv.Atoi(v[0])
			max, _ := strconv.Atoi(v[1])
			p := pair{min: min, max: max}
			pairs = append(pairs, p)
			spans = append(spans, p)
		}
		classes = append(classes, class{spans: pairs, desc: l[0]})
	}
	scanner.Scan()
	line := scanner.Text()
	if line != "your ticket:" {
		panic("Invalid input")
	}
	scanner.Scan()
	myTicket = newticket(scanner.Text(), 1000)
	tickets = append(tickets, myTicket)
	scanner.Scan()
	scanner.Scan()
	line = scanner.Text()
	if line != "nearby tickets:" {
		panic("Invalid input")
	}

	invalidTickets := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		ticket := newticket(line, len(myTicket)-1)

		if isValidTicket(ticket, spans) {
			tickets = append(tickets, ticket)
		} else {
			invalidTickets = append(invalidTickets, ticket)
		}
	}

	fmt.Printf("My: %d\n", myTicket)
	//	fmt.Printf("Tickets: %d\n", tickets)
	//	fmt.Printf("Invalid: %d\n", invalidTickets)

	match := [][]bool{}
	for i, c := range classes {
		match = append(match, makeBase(myTicket))
		for _, t := range tickets {
			for j, _ := range myTicket {
				if !isValidVal(t[j], c.spans) {
					match[i][j] = false
				}
			}
		}
	}

	found := make(map[int]int)
	for len(found) < len(match[0]) {
		for i, r := range match {
			k := findTrue(r)
			if k > -1 {
				found[i] = k

				fmt.Printf("%d %d %s\n", k, i, classes[i].desc)

				for _, m := range match {
					m[k] = false
				}
			}
		}
	}
	sum2 := int64(1)
	for i, k := range found {
		if i < 6 {
			sum2 *= int64(myTicket[k])
			fmt.Printf("Found: %d, %d %d %s\n", i, k, myTicket[k], classes[i].desc)
		}
	}

	for _, t := range invalidTickets {
		for _, v := range t {
			if !isValidVal(v, spans) {
				sum1 += v
			}
		}
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: \n%v \n%v\n", sum2, 5311123569883)
}

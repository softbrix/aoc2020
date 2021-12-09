package main

import (
	"fmt"
)

func err(e error) {
	if e != nil {
		panic(e)
	}
}

type Node struct {
	v    int
	next *Node
}

func printLoop(el *Node) {
	find(0, el)
}

func find(value int, el *Node) *Node {
	start := el
	for el != nil {
		if el.v == value {
			return el
		}
		el = el.next
		if start == el {
			break
		}
	}
	return nil
}

func main() {
	MAX := 1000001
	cards := []int{3, 2, 7, 4, 6, 5, 1, 8, 9}
	//DEMO := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	//cards = DEMO
	index := make([]*Node, MAX, MAX)

	root := Node{}

	last := &root
	for _, c := range cards {
		n := Node{
			v: c,
		}
		last.next = &n
		last = &n
		index[c] = &n
	}
	// Part 2
	for i := 10; i < MAX; i++ {
		n := Node{
			v: i,
		}
		last.next = &n
		last = &n
		index[i] = &n
	}
	LEN := 1000000

	// Close the loop
	first := root.next
	last.next = first

	curr := first

	for loop := 0; loop < 10000000; loop++ {
		pick := []*Node{curr.next, curr.next.next.next}
		curr.next = pick[1].next
		pick[1].next = nil

		var dst *Node
		for j := 1; dst == nil; j++ {
			lbl := (LEN + curr.v - j) % LEN
			if lbl == 0 {
				lbl = LEN
			}
			if find(lbl, pick[0]) != nil {
				continue
			}
			if j > 9 {
				panic("No element found")
			}
			dst = index[lbl]
		}

		// Swappie...
		tmp := dst.next
		dst.next = pick[0]
		pick[1].next = tmp

		curr = curr.next

		if loop%100000 == 0 {
			print(".")
		}
	}

	first = index[1]
	fmt.Printf("%d\n", first.v)
	v := []int{first.next.v, first.next.next.v}
	fmt.Printf("\n%d\n", v)
	fmt.Printf("Sum2: %d\n", v[0]*v[1])
}

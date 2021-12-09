package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func P(x, y, z, w int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z) + "," + strconv.Itoa(w)
}

func loop(ma map[string]bool, it int) int {
	fmt.Printf("Loop: %d %d\n", it, len(ma))
	/*for k, v := range ma {
		fmt.Printf("%s: %v\n", k, v)
	}
	fmt.Printf("==========\n")
	*/
	if it == 6 {
		return len(ma)
	}

	s := []int{-1, 0, 1}
	manext := make(map[string]bool)
	for w := -7; w <= 7; w++ {
		for z := -7; z <= 7; z++ {
			for x := -15; x <= 15; x++ {
				for y := -15; y <= 15; y++ {
					neighbours := 0
					for _, dx := range s {
						for _, dy := range s {
							for _, dz := range s {
								for _, dw := range s {
									if dz != 0 || dx != 0 || dy != 0 || dw != 0 {
										p := P(x+dx, y+dy, z+dz, w+dw)

										if ma[p] {
											// fmt.Printf("p: %s <= %s\n", p, P(x, y, z, w))
											neighbours++
										}
									}
								}
							}
						}
					}
					p := P(x, y, z, w)
					if ma[p] {
						if neighbours == 2 || neighbours == 3 {
							manext[p] = true
						}
					} else if neighbours == 3 {
						manext[p] = true
					}
				}
			}
		}
	}
	it++
	return loop(manext, it)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ma1 := make(map[string]bool)
	y := 0
	xmax := 0
	for scanner.Scan() {
		l := scanner.Text()
		for x, c := range l {
			p := P(x, y, 0, 0)
			if c == '.' {
				//ma1[p] = false
			} else if c == '#' {
				ma1[p] = true
			} else {
				panic("undefined input")
			}
			if x > xmax {
				xmax = x
			}
		}
		y++
	}
	xmax++

	sum1 := loop(ma1, 0)

	fmt.Printf("\nPart 2: %d\n", sum1)
}

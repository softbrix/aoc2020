package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dir int

var E Dir = 1
var W Dir = 2
var NE Dir = 3
var NW Dir = 4
var SE Dir = 5
var SW Dir = 6

func match(c string) Dir {
	switch c {
	case "ne":
		return NE
	case "nw":
		return NW
	case "se":
		return SE
	case "sw":
		return SW
	}
	switch c[:1] {
	case "e":
		return E
	case "w":
		return W
	}
	panic("No match")
}

func delta(dir Dir) (int, int) {
	x, y := 0, 0
	switch dir {
	case E:
		x += 1
	case W:
		x -= 1
	case NW:
		x -= 1
		y -= 1
	case NE:
		y -= 1
	case SW:
		y += 1
	case SE:
		y += 1
		x += 1
	default:
		panic("Unknown option")
	}
	return x, y
}

func adjacent(p Pos, tiles map[Pos]int) int {
	cnt := 0
	for d := E; d <= SW; d++ {
		dx, dy := delta(d)
		pos := Pos{p.x + dx, p.y + dy}
		if tiles[pos] == 1 {
			cnt += 1
		}
	}
	return cnt
}

type Pos struct {
	x, y int
}

func err(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rows := make([][]Dir, 0, 0)
	for scanner.Scan() {
		r := scanner.Text()
		row := make([]Dir, 0, 0)
		for len(r) > 0 {
			c := r
			if len(r) >= 2 {
				c = r[:2]
			}
			v := match(c)
			row = append(row, v)
			if v >= NE {
				r = r[2:]
			} else {
				r = r[1:]
			}
		}
		rows = append(rows, row)
	}
	tiles := make(map[Pos]int)
	for _, row := range rows {
		x, y := 2, 1
		for _, d := range row {
			dx, dy := delta(d)
			x, y = x+dx, y+dy
		}
		p := Pos{x: x, y: y}
		tiles[p] = tiles[p] + 1
	}
	sum1 := 0
	for k, v := range tiles {
		if v%2 == 1 {
			tiles[k] = 1
			sum1++
		} else {
			delete(tiles, k)
		}
	}

	for i := 0; i < 100; i++ {
		newTiles := make(map[Pos]int)
		nearby := make(map[Pos]bool)
		for k, _ := range tiles {
			for d := E; d <= SW; d++ {
				dx, dy := delta(d)
				pos := Pos{k.x + dx, k.y + dy}
				if tiles[pos] == 0 {
					nearby[pos] = true
				}
			}
		}
		// Black
		for k, _ := range tiles {
			cnt := adjacent(k, tiles)
			if cnt > 0 && cnt <= 2 {
				newTiles[k] = 1
			}
		}
		// White
		for k, _ := range nearby {
			if adjacent(k, tiles) == 2 {
				newTiles[k] = 1
			}
		}
		tiles = newTiles
	}

	fmt.Printf("Part1: %v\n", sum1)
	fmt.Printf("Part2: %v\n", len(tiles))
}
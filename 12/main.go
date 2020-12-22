package main

import (
    "bufio"
		"fmt"
		"strconv"
    "os"
)

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func xd(d int) int {
	if d == 1 {
		return 1
	}
	if d == 3 {
		return -1
	}
	return 0
}

func yd(d int) int {
	if d == 2 {
		return 1
	}
	if d == 0 {
		return -1
	}
	return 0
}

func rtx(d int, x int, y int) int {
	if d == 1 {
		return -y
	} else if d == 2 {
		return -x
	} else if d == 3 {
		return y
	} else {
		return x
	}
}

func rty(d int, x int, y int) int {
	if d == 1 {
		return x
	} else if d == 2 {
		return -y
	} else if d == 3 {
		return -x
	} else {
		return y
	}
}

func part1() {
/*	l := scanner.Text()
	d := rune(l[0])
	n,_ := strconv.Atoi(l[1:])
	if d == 'L' {
		if n > 360 {
			fmt.Printf("Full loop %d", n)
		}
		dir -= n / 90
		dir = (dir + 4) % 4
	} else if d == 'R' {
		dir += n / 90
		dir = dir % 4
	} else {
		switch d {
		case 'N': y -= n; break
		case 'S': y += n; break
		case 'W': x -= n; break
		case 'E': x += n; break;
		case 'F': 
			x += xd(dir) * n;
			y += yd(dir) * n;
			break;
		default:
			fmt.Printf("Unknown Step '%s'", d)
		}
	}*/
}

func main() {
	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	// 0 = UP
	// 1 = RIGHT
	// 2 = DOWN
	// 3 = LEFT
	//dir := 1
	dy := -1
	dx := 10
	y := 0
	x := 0
  for scanner.Scan() {
		l := scanner.Text()
		d := rune(l[0])
		n,_ := strconv.Atoi(l[1:])
		if d == 'L' {
			d := -(n / 90)
			d = (d + 4) % 4
			tdy := dy
			dy = rty(d, dx, dy)
			dx = rtx(d, dx, tdy)
		} else if d == 'R' {
			d := n / 90
			tdy := dy
			dy = rty(d, dx, dy)
			dx = rtx(d, dx, tdy)
		} else {
			switch d {
			case 'N': dy -= n; break
			case 'S': dy += n; break
			case 'W': dx -= n; break
			case 'E': dx += n; break;
			case 'F': 
				x += dx * n;
				y += dy * n;
				break;
			default:
				fmt.Printf("Unknown Step '%s'", d)
			}
		}
	}
	
	sum2 = abs(x) + abs(y)

	fmt.Printf("Sum1: %d %d %d \n", sum1, x, y)
	fmt.Printf("Sum2: %d\n", sum2)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OP func(int, int) int

func err(e error) {
	if e != nil {
		panic(e)
	}
}

func mul(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func val(a, _ int) int {
	return a
}

func calc(s string) string {
	if strings.Index(s, "(") > 0 || strings.Index(s, ")") > 0 {
		fmt.Println("Bad str: " + s)
		panic("Can't calc with parentesis")
	}
	sum := 0
	fields := strings.Fields(s)
	newFields := make([]string, 0, 0)
	for i := 0; i < len(fields); i++ {
		c := fields[i]
		if c == "+" {
			last := len(newFields) - 1
			a, _ := strconv.Atoi(newFields[last])
			b, _ := strconv.Atoi(fields[i+1])
			newFields[last] = strconv.Itoa(add(a, b))
			i++
		} else {
			newFields = append(newFields, c)
		}
	}
	var op OP
	for i, c := range newFields {
		if i == 0 {
			sum, _ = strconv.Atoi(c)
			continue
		}
		if c == "+" {
			panic("Unexpected")
		} else if c == "*" {
			op = mul
		} else {
			v, e := strconv.Atoi(c)
			err(e)
			if op != nil {
				sum = op(sum, v)
				op = nil
			} else {
				panic("No operand")
			}
		}
	}
	return strconv.Itoa(sum)
}

func reduce(s string) string {
	for true {
		p := strings.Index(s, ")")
		if p < 0 {
			break
		}
		t := s[:p]
		p2 := strings.LastIndex(t, "(")
		t = calc(t[p2+1:])
		s = s[:p2] + t + s[p+1:]
	}
	return s
}

func main() {
	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		r := scanner.Text()
		fmt.Println(r)
		r = reduce(r)
		s, _ := strconv.Atoi(calc(r))
		sum1 += s

		fmt.Printf("sum: %s ==== %d\n", r, s)
	}

	if scanner.Err() != nil {
		// handle error.
		panic(scanner.Err())
	}

	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}

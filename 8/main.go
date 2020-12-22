package main

import (
    "bufio"
		"fmt"
		"strings"
		"strconv"
    "os"
)

type Inst struct{
	op string
	val int
	run bool
}

func switchOp(op string) string {
	if op == "jmp" {
		return "nop"
	} else if op == "nop" {
		return "jmp"
	}
	return op
}

func seek(inst []Inst) (int, int) {
	I := len(inst)
	acc := 0
	i := 0
	for ; i < I; i++ {
		if (inst[i].run) {
			break;
		}
		inst[i].run = true
		switch inst[i].op {
			case "acc": acc += inst[i].val;
			case "jmp": i += inst[i].val - 1;
		}
	}
	return i, acc
}
 
func main() {

	sum1 := 0
	sum2 := 0
	inst := make([]Inst, 0, 0)
	scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
		l := strings.Fields(scanner.Text())
		val,_ := strconv.Atoi(l[1])
		inst = append(inst, Inst {
			op: l[0],
			val: val,
			run: false,
		})
  }

	if scanner.Err() != nil {
		// handle error.
		panic(scanner.Err())
	}


	// Part 1
	_,sum1 = seek(inst)
	fmt.Printf("========\n")
	// Part 2
	I := len(inst)
	for j := I - 1; j > 0; j-- {
		inst[j].op = switchOp(inst[j].op)
		i := 0
		for ; i < I; i++ {
			inst[i].run = false
		}
		i,sum2 = seek(inst)
		if (i >= I) {
			fmt.Println("FOUND")
			break;
		}
		inst[j].op = switchOp(inst[j].op)
	}
	fmt.Printf("Sum1: %d %d\n", sum1, len(inst))
	fmt.Printf("Sum2: %d\n", sum2)
}
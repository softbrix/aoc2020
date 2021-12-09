package main

import (
	"fmt"
)

func err(e error) {
	if e != nil {
		panic(e)
	}
}

var MASTER uint64 = uint64(20201227)

func pow(base, loop uint64) uint64 {
	v := uint64(1)
	inc := uint64(base)
	for i := uint64(0); i < loop; i++ {
		v *= inc
		v = v % MASTER
	}
	return v
}

func main() {
	BASE := uint64(7)
	//pubKeys := []uint64{5764801, 17807724} // DEMO
	pubKeys := []uint64{9093927, 11001876}
	loop := []uint64{0, 0}

	for l, pub := range pubKeys {
		v := uint64(1)
		for i := uint64(1); i < 10000000000 && loop[l] == 0; i++ {
			v *= BASE
			v = v % MASTER
			if v == pub {
				fmt.Printf("Found: %d", i)
				loop[l] = i
			}
		}
	}

	fmt.Printf("\n%d %d\n", loop, pubKeys)
	fmt.Printf("%d\n", pow(pubKeys[0], loop[1]))
	fmt.Printf("%d\n", pow(pubKeys[1], loop[0]))

}

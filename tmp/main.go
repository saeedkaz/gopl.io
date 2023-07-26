package main

import (
	"fmt"
)

func main() {

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	xx := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", xx)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
	// var x uint8 = 1<<1 | 1<<5
	// var y uint8 = 1<<1 | 1<<2
	// fmt.Printf("1<<2:%08b\n", 5<<2) // 1<<1
	// fmt.Printf("%08b\n", x)         // "00100010", the set {1, 5}
	// fmt.Printf("%08b\n", x)         // "00100010", the set {1, 5}
	// fmt.Printf("%08b\n", x)         // "00100010", the set {1, 5}
	// fmt.Printf("%08b\n", y)         // "00000110", the set {1, 2}
	// fmt.Printf("%08b\n", x&y)       // "00000010", the intersection {1}
	// fmt.Printf("%08b\n", x|y)       // "00100110", the union {1, 2, 5}
	// fmt.Printf("%08b\n", x^y)       // "00100100", the symmetric difference {2, 5}
	// fmt.Printf("%08b\n", x&^y)      // "00100000", the difference {5}
	// for i := uint(0); i < 8; i++ {
	// 	if x&(1<<i) != 0 { // membership test
	// 		fmt.Println(i) // "1", "5"
	// 	}
	// }
	// fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	// fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

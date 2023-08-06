package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
func main() {
	ss := []string{"apple", "banana", "cherry"}
	fmt.Println(k(ss))
	fmt.Println(ss)
	Add(ss)
	fmt.Println(Count(ss))
	fmt.Println(m)
	ss2 := []string{"apple2", "banana2", "cherry2"}
	fmt.Println(k(ss2))
	fmt.Println(ss2)
	Add(ss2)
	fmt.Println(Count(ss2))
	fmt.Println(m)
}

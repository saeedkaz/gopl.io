package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
<<<<<<< HEAD
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s)
	fmt.Printf("%v\n", s)

=======
	var x, y []int
	for i := 0; i < 10; i++ {
		y = append(x, i, i+1)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
>>>>>>> 21794508e9a70671c0b57b8d358874536f53e28c
}

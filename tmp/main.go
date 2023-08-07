package main

import (
	"fmt"

	"gopl.io/ch4/treesort" // import the treesort package
)

func main() {
	values := []int{3, 1, 4, 5, 9, 2, 6, 8, 7} // create a slice of integers
	fmt.Println("Before sorting:", values)     // print the slice before sorting
	treesort.Sort(values)                      // call the Sort function from the treesort package
	fmt.Println("After sorting:", values)      // print the slice after sorting
}

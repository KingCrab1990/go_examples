package main

import (
	"fmt"
	"sort"
)

func main() {
	// created a slice of int
	a := []int{3, 6, 2, 1, 9, 10, 8}
	sort.Ints(a)
	for _, v := range (a) {
		fmt.Println(v)
	}
}

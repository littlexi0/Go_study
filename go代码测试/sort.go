package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	a = append(a, 1)
	a = append(a, 3)
	a = append(a, 2)
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)
}

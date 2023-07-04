package main

import "fmt"

func main() {
	mp := make(map[int]string)
	mp[1] = "one"
	mp[2] = "two"
	mp[3] = "three"
	fmt.Println(mp)
	for k, v := range mp {
		fmt.Println(k, v)
	}
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

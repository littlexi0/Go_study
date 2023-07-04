package main

import "fmt"

type base struct {
	id     int
	gender bool
}

type container struct {
	base
	name string
}

func main() {
	co := container{
		base: base{
			id:     1,
			gender: true,
		},
		name: "container",
	}
	fmt.Println(co.id, co.name)
	fmt.Println(co.base.id)
	fmt.Println(co)
}

package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Println(p.name)
}

func main() {
	var p Person
	p1 := Person{"mike", 15}
	p2 := Person{name: "mike", age: 15}
	p4 := Person{}

	fmt.Println(p, p1, p2, p4)

	p.name = "Smith"
	p.age = 24

	p.PrintName()
}

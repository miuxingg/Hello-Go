package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

type student struct {
	Human
	id                        string
	math, physical, chemistry float32
}

type Speak interface {
	sayHi(name string)
}

func (h Human) sayHi(name string) {
	fmt.Printf("\nHii world! I'm a %v", name)
}

func main() {
	fmt.Println("--------- STRUCT")
	var student1 student
	student1.id = "id1"
	student1.name = "Pham Mai Ngoc Anh"
	student1.age = 16
	student1.math = 8
	student1.physical = 8
	student1.chemistry = 8

	student2 := new(student)
	student2.id = "id2"
	student2.name = "Pham Tran Mai Anh"
	student2.age = 18
	student2.math = 9
	student2.physical = 9
	student2.chemistry = 9

	fmt.Printf("Student 1: %v\n", student1)
	fmt.Printf("Student 2: %v\n", *student2)
	fmt.Printf("Student name 1: %v\n", student1.name)

	var i Speak
	i = student1
	i.sayHi(student1.name)
}

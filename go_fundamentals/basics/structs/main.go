package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person
	var P person
	P.name = "Coolio"
	P.age = 32
	P2 := person{"Frodo", 55}
	fmt.Printf("The person's name is %s\n", P.name)
	fmt.Printf("The hobbit's name is %s\n", P2.name)
	tom.name, tom.age = "Tom", 18
	bob := person{age: 25, name: "Bob"}
	paul := person{"Paul", 43}
	tbOlder, tbdiff := Older(tom, bob)
	tpOlder, tpdiff := Older(tom, paul)
	bpOlder, bpdiff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tbOlder.name, tbdiff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tpOlder.name, tpdiff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bpOlder.name, bpdiff)
}

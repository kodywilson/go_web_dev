package main

import (
	"errors"
	"fmt"
)

func main() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	m := `hello
			world`

	fmt.Println(m)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

	const (
		i      = 100
		pi     = 3.1415
		prefix = "Go_"
	)

	fmt.Println(i, pi, prefix)

	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The second element is %d\n", arr[1])

	// two dimensional array
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(easyArray)

	// maps
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3

	fmt.Println("The third number is: ", numbers["three"])
	fmt.Println("The length of this map is ", len(numbers))
	numbers["one"] = 100
	fmt.Println("Now one is ", numbers["one"])

	delete(numbers, "three")
	fmt.Println("The length of this map is now", len(numbers))

	m1 := make(map[string]string)
	m1["Hello"] = "Bonjour"
	fmt.Println("Hello is", m1["Hello"])
	m2 := m1
	m2["Hello"] = "Salut"
	fmt.Println("Now, hello is", m1["Hello"])

	x := 5
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

}

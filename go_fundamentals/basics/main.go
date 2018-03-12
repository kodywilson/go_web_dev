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
}

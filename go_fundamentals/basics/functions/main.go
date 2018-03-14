package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}

func main() {
	x := 3
	y := 4
	z := 5

	maxxy := max(x, y)
	maxxz := max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, maxxy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, maxxz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))

	a := 3
	b := 4

	xPLUSy, xTIMESy := SumAndProduct(a, b)

	fmt.Printf("%d + %d = %d\n", a, b, xPLUSy)
	fmt.Printf("%d * %d = %d\n", a, b, xTIMESy)

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

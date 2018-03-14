package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
}

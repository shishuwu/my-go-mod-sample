package main

import "fmt"

// return multiple values
func cal(x int, y int) (int, int) {
	sum := x + y
	diff := x - y
	return sum, diff
}

func main() {
	a, b := 5, 3
	sum, diff := cal(a, b)

	fmt.Println("sum", sum)
	fmt.Println("diff",diff)
}

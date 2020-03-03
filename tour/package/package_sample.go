package main

import (
	"fmt"
	"github/shishuwu/my-go-mod-sample/common/calc"
)

func main() {
	x, y := 15, 10
	sum := calc.Do_add(x, y)

	fmt.Println("sum", sum)
}

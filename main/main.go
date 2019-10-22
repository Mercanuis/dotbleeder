package main

import (
	dots2 "dotbleeder/dots"
	"fmt"
)

func main() {
	totalRows := 3
	totalCols := 3

	dot1 := dots2.NewDot(0, 0, 15)
	dot2 := dots2.NewDot(2, 2, 15)

	dots := []dots2.Dot{dot1, dot2}

	s := dots2.NewSheet(totalRows, totalCols, dots)
	fmt.Println(s.SolveSheet())
	fmt.Println()
}

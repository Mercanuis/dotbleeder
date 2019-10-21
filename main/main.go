package main

import (
	dots2 "dotbleeder/dots"
	"fmt"
)

func main() {
	totalRows := 3
	totalCols := 3

	dot1 := dots2.NewDot(0, 0, 10)
	dot2 := dots2.NewDot(1, 2, 10)

	dots := []dots2.Dot{dot1, dot2}

	s := dots2.NewSheet(totalRows, totalCols, dots)
	fmt.Println(s.SolveSheet())
	fmt.Println()

	totalRows = 5
	totalCols = 5

	/*
		   0   1   2   3   4
		0
		1         25
		2
		3
		4
	*/

	dot3 := dots2.NewDot(1, 2, 25)
	dot4 := dots2.NewDot(3, 4, 25)

	dots22 := []dots2.Dot{dot3, dot4}
	s2 := dots2.NewSheet(totalRows, totalCols, dots22)
	fmt.Println(s2.SolveSheet())
}

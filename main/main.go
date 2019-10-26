package main

import (
	dots2 "dotbleeder/dots"
	"fmt"
)

func main() {
	totalRows := 5
	totalCols := 5

	dot1 := dots2.NewDot(0, 0, 10)
	dot2 := dots2.NewDot(3, 4, 10)

	dots := []dots2.Dot{dot1, dot2}

	s := dots2.NewSheet(totalRows, totalCols, dots)
	fmt.Println(s.SolveSheet())
	fmt.Println()
}

package dots

import "fmt"

type Sheet struct {
	Dots    []Dot
	NumRows int
	NumCols int
}

func NewSheet(rows, cols int, dots []Dot) Sheet {
	return Sheet{
		Dots:    dots,
		NumRows: rows,
		NumCols: cols,
	}
}

func (s *Sheet) SolveSheet() int {
	fmt.Println("STARTING SHEET")

	grid := make([][]int, s.NumRows)
	for i := range grid {
		grid[i] = make([]int, s.NumCols)
	}

	//for _, dot := range s.Dots {
	s.doBleedDarkness(grid)
	//	fmt.Println("FINISHED DOT")
	//}
	//fmt.Println("DONE WITH DOTS")

	total := 0
	for i := range grid {
		for j := range grid[i] {
			total += grid[i][j]
		}
	}

	return total
}

func (s *Sheet) doBleedDarkness(grid [][]int) {
	//reset the memo for each dot
	memo := make([][]bool, s.NumRows)
	for i := range memo {
		memo[i] = make([]bool, s.NumCols)
	}

	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = false
		}
	}

	for _, dot := range s.Dots {
		s.bleedDarkness(dot.Row, dot.Col, dot.Darkness, grid, memo)
		fmt.Println("FINISHED DOT")
	}
	fmt.Println("DONE WITH DOTS")

	//s.bleedDarkness(row, col, dark, grid, memo)
}

func (s *Sheet) bleedDarkness(row int, col int, dark int, grid [][]int, memo [][]bool) {
	fmt.Printf("row = %d, col = %d, dark = %d\n", row, col, dark)

	if row < 0 || row >= s.NumRows || col < 0 || col >= s.NumCols {
		fmt.Printf("OUT OF BOUNDS! Dark is at [%d]\n", dark)
		//Out of bounds
		return
	} else if memo[row][col] {
		//We've seen this before, but its possible another path could lead to a higher dark value
		//Make a check and then return
		if dark > grid[row][col] {
			fmt.Printf("UPDATE VISITED CELL: [row %d, col %d] dark is now [%d]\n", row, col, dark)
			grid[row][col] = dark
		}
		return
	} else if grid[row][col] > dark {
		//Current value > dark, return
		fmt.Printf("[row %d, col %d] dark [%d] > [%d]\n", row, col, grid[row][col], dark)
		memo[row][col] = true
		return
	} else {
		//This is a cell to update, update the cell, memo, and then traverse
		fmt.Printf("UPDATE CELL: [row %d, col %d] dark is now [%d]\n", row, col, dark)
		grid[row][col] = dark
		//dark--
		memo[row][col] = true

		//left
		s.bleedDarkness(row-1, col, dark-1, grid, memo)
		//right
		s.bleedDarkness(row+1, col, dark-1, grid, memo)
		//up
		s.bleedDarkness(row, col-1, dark-1, grid, memo)
		//down
		s.bleedDarkness(row, col+1, dark-1, grid, memo)
	}
}

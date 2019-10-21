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
	grid := make([][]int, s.NumRows)
	for i := range grid {
		grid[i] = make([]int, s.NumCols)
	}

	for _, dot := range s.Dots {
		s.doBleedDarkness(dot.Row, dot.Col, dot.Darkness, grid)
		fmt.Println("FINISHED DOT")
	}
	fmt.Println("DONE WITH DOTS")

	total := 0
	for i := range grid {
		for j := range grid[i] {
			total += grid[i][j]
		}
	}

	return total
}

func (s *Sheet) doBleedDarkness(row int, col int, dark int, grid [][]int) {
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

	s.bleedDarkness(row, col, dark, grid, memo)
}

func (s *Sheet) bleedDarkness(row int, col int, dark int, grid [][]int, memo [][]bool) {
	if row < 0 || row >= s.NumRows || col < 0 || col >= s.NumCols {
		//Out of bounds
		return
	} else if memo[row][col] {
		//We've seen this before, but its possible another path could lead to a higher dark value
		//Make a check and then return
		if dark > grid[row][col] {
			grid[row][col] = dark
		}
		return
	} else if grid[row][col] > dark {
		//Current value > dark, return
		memo[row][col] = true
		return
	} else {
		//This is a cell to update, update the cell, memo, and then traverse
		grid[row][col] = dark
		dark--
		memo[row][col] = true

		//left
		s.bleedDarkness(row-1, col, dark, grid, memo)
		//right
		s.bleedDarkness(row+1, col, dark, grid, memo)
		//up
		s.bleedDarkness(row, col-1, dark, grid, memo)
		//down
		s.bleedDarkness(row, col+1, dark, grid, memo)
	}
}

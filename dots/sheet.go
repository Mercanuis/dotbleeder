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

	s.doBleedDarkness(grid)

	total := 0
	for i := range grid {
		for j := range grid[i] {
			total += grid[i][j]
		}
	}

	return total
}

func (s *Sheet) doBleedDarkness(grid [][]int) {
	for _, dot := range s.Dots {
		s.bleedDarkness(dot.Row, dot.Col, dot.Darkness, grid)
		fmt.Println("FINISHED DOT")
	}
	fmt.Println("DONE WITH DOTS")
}

func (s *Sheet) bleedDarkness(row int, col int, dark int, grid [][]int) {
	if row < 0 || row >= s.NumRows || col < 0 || col >= s.NumCols {
		//Out of bounds
		return
	} else if grid[row][col] > dark {
		//Current value > dark, return
		return
	} else {
		//This is a cell to update, update the cell, memo, and then traverse
		grid[row][col] = dark

		//left
		s.bleedDarkness(row-1, col, dark-1, grid)
		//right
		s.bleedDarkness(row+1, col, dark-1, grid)
		//up
		s.bleedDarkness(row, col-1, dark-1, grid)
		//down
		s.bleedDarkness(row, col+1, dark-1, grid)
	}
}

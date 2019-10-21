package dots

type Dot struct {
	Row      int
	Col      int
	Darkness int
}

func NewDot(row, col, darkness int) Dot {
	return Dot{
		Row:      row,
		Col:      col,
		Darkness: darkness,
	}
}

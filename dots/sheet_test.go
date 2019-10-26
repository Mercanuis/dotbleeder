package dots

import (
	"testing"
)

func TestSolveSheet(t *testing.T) {
	testDot1 := NewDot(0, 0, 15)
	testDot2 := NewDot(1, 1, 15)

	cases := map[string]struct {
		testSheet Sheet
		expected  int
	}{
		"2x2": {
			NewSheet(2, 2, []Dot{testDot1, testDot2}),
			58,
		},
		"3x3": {
			NewSheet(3, 3, []Dot{
				NewDot(0, 0, 10), NewDot(1, 2, 10),
			}),
			81,
		},
		"1x3": {
			NewSheet(1, 3, []Dot{
				NewDot(0, 0, 15), NewDot(0, 1, 25), NewDot(0, 2, 39),
			}),
			114,
		},
		"1x1": {
			NewSheet(1, 1, []Dot{
				NewDot(0, 0, 25), NewDot(0, 0, 20),
			}),
			25,
		},
		"4x4": {
			NewSheet(4, 4, []Dot{
				NewDot(1, 2, 15), NewDot(0, 3, 15),
			}),
			210,
		},
		"5x5": {
			NewSheet(5, 5, []Dot{
				NewDot(1, 2, 10), NewDot(3, 4, 10),
			}),
			199,
		},
		"6x3": {
			NewSheet(6, 3, []Dot{
				NewDot(2, 2, 25), NewDot(6, 0, 15), NewDot(4, 1, 30),
			}),
			495,
		},
		"9x9": {
			NewSheet(9, 9, []Dot{
				NewDot(1, 1, 9), NewDot(5, 5, 15), NewDot(6, 9, 10), NewDot(0, 8, 35),
			}),
			2187,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			//expected := tc.expected
			result := tc.testSheet.SolveSheet()
			if tc.expected != result {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

# dotbleeder
This program is a small exercise in recurrsion
The idea is a simple bleeder, imagining dots on a page and how the ink "bleeds" to other cells

For example
```
| | | |
| | | |
| | | | 
```
On an initial 3x3 grid, adding a dot of darkness 10 at [0,0] does the following
```
|10|9 |8 |
|9 |8 |7 |
|8 |7 |6 |
```
Then adding a 2nd dot at [1,2] does the following
```
|10 |9 |9 |
|9 |9 |10 |
|8 |8 |9 |
```
Resulting in a total "Darkness" of 81, because we can only overwrite cells which have a lower dark value than the one that could come in

The exercise uses recurrsion to solve the issue, traversing the grid and keeping track of cells already valued to determine which ones have yet to reach a final value

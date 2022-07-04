package main

import (
	"fmt"
	"github.com/jhleakakos/sudoku-go/board"
)

func main() {

	practiceBoard := board.Board{
		Grid: [9][9]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{2, 3, 4, 5, 6, 7, 8, 9, 0},
			{3, 4, 5, 6, 7, 8, 9, 0, 1},
			{4, 5, 6, 7, 8, 9, 0, 1, 2},
			{5, 6, 7, 8, 9, 0, 1, 2, 3},
			{6, 7, 8, 9, 0, 1, 2, 3, 4},
			{7, 8, 9, 0, 1, 2, 3, 4, 5},
			{8, 9, 0, 1, 2, 3, 4, 5, 6},
		},
	}
	fmt.Print(practiceBoard.GetBoard())

}

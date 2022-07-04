package main

import (
	"fmt"
	"strconv"
)

type Board struct {
	grid [9][9]int
	//grid [9]int
}

func (b *Board) printBoard() {

	fmt.Printf("%+v\n", b)
	fmt.Println(b.grid)

	fmt.Println("| ---   ---   --- |")
	sb := ""
	for i, n := range b.grid {
		sb += "| "
		for j, n := range n {
			//fmt.Printf("j: %d, n: %d\n", j, n)
			sb += strconv.Itoa(n)

			if j == 2 || j == 5 {
				sb += " | "
			}
		}
		sb += " |\n"
		if (i+1)%3 == 0 {
			sb += "| --- | --- | --- |\n"
		}
	}

	fmt.Print(sb)
}

func main() {

	practiceBoard := Board{
		grid: [9][9]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
			{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	practiceBoard.printBoard()

}

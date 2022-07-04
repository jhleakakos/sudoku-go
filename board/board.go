package board

import "strconv"

type Board struct {
	Grid [9][9]int
}

func (b *Board) GetBoard() string {

	sb := "| ---   ---   --- |\n"
	for i, n := range b.Grid {
		sb += "| "
		for j, n := range n {
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

	return sb
}

func (b *Board) CheckIsValidRow(row int) bool {
	rowValues := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
		6: false,
		7: false,
		8: false,
		9: false,
	}

	for _, num := range b.Grid[row] {
		if num < 1 || num > 9 || rowValues[num] == true {
			return false
		}

		rowValues[num] = true
	}

	return true
}

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

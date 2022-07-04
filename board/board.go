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

func (b *Board) CheckIsValidColumn(col int) bool {
	colValues := map[int]bool{
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

	for _, row := range b.Grid {
		if row[col] < 1 || row[col] > 9 || colValues[row[col]] == true {
			return false
		}

		colValues[row[col]] = true
	}

	return true

}

func (b *Board) CheckIsValidSquare(square int) bool {
	squareValues := map[int]bool{
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

	squareStartingIndices := map[int][2]int{
		1: {0, 0},
		2: {0, 3},
		3: {0, 6},
		4: {3, 0},
		5: {3, 3},
		6: {3, 6},
		7: {6, 0},
		8: {6, 3},
		9: {6, 6},
	}

	idxRowStart := squareStartingIndices[square][0]
	idxColStart := squareStartingIndices[square][1]

	for r := idxRowStart; r < idxRowStart+3; r++ {
		for c := idxColStart; c < idxColStart+3; c++ {
			num := b.Grid[r][c]
			if num < 1 || num > 9 || squareValues[num] == true {
				return false
			}
			squareValues[num] = true
		}
	}

	return true
}

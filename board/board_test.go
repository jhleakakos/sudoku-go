package board

import (
	"testing"
)

func TestGetBoard(t *testing.T) {
	testBoard := Board{
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

	got := testBoard.GetBoard()
	expected :=
		"| ---   ---   --- |\n" +
			"| 012 | 345 | 678 |\n" +
			"| 123 | 456 | 789 |\n" +
			"| 234 | 567 | 890 |\n" +
			"| --- | --- | --- |\n" +
			"| 345 | 678 | 901 |\n" +
			"| 456 | 789 | 012 |\n" +
			"| 567 | 890 | 123 |\n" +
			"| --- | --- | --- |\n" +
			"| 678 | 901 | 234 |\n" +
			"| 789 | 012 | 345 |\n" +
			"| 890 | 123 | 456 |\n" +
			"| --- | --- | --- |\n"

	if got != expected {
		t.Error("Problem with printed board matching GetBoard() output")
	}
}

func TestCheckIsValidRow(t *testing.T) {
	// first row is valid
	// second row is invalid because of two 9s
	// third row is invalid because it has a 0
	// fourth row is invalid because it has a 10
	testBoard := Board{
		Grid: [9][9]int{
			{8, 7, 1, 5, 6, 9, 3, 2, 4},
			{9, 7, 1, 5, 6, 9, 3, 2, 4},
			{6, 2, 5, 4, 0, 1, 9, 8, 7},
			{5, 8, 4, 1, 7, 10, 6, 3, 9},
			{2, 9, 7, 3, 8, 6, 4, 5, 1},
			{3, 1, 6, 9, 4, 5, 8, 7, 2},
			{7, 5, 8, 6, 1, 4, 2, 9, 3},
			{9, 4, 2, 8, 5, 3, 7, 1, 6},
			{1, 6, 3, 2, 9, 7, 5, 4, 8},
		},
	}

	expected := testBoard.CheckIsValidRow(0)
	if !expected {
		t.Error("good row returning invalid")
	}

	expected = testBoard.CheckIsValidRow(1)
	if expected {
		t.Error("bad row with duplicate numbers returning valid")
	}

	expected = testBoard.CheckIsValidRow(2)
	if expected {
		t.Error("bad row with a 0 returning valid")
	}

	expected = testBoard.CheckIsValidRow(3)
	if expected {
		t.Error("bad row with a 10 returning valid")
	}
}

func TestCheckIsValidColumn(t *testing.T) {
	// first column is valid
	// second column is invalid because of two 9s
	// third column is invalid because it has a 0
	// fourth column is invalid because it has a 10
	testBoard := Board{
		Grid: [9][9]int{
			{8, 9, 1, 5, 6, 9, 3, 2, 4},
			{4, 3, 9, 7, 2, 8, 1, 6, 5},
			{6, 2, 5, 4, 3, 1, 9, 8, 7},
			{5, 8, 4, 1, 7, 2, 6, 3, 9},
			{2, 9, 0, 3, 8, 6, 4, 5, 1},
			{3, 1, 6, 9, 4, 5, 8, 7, 2},
			{7, 5, 8, 6, 1, 4, 2, 9, 3},
			{9, 4, 2, 10, 5, 3, 7, 1, 6},
			{1, 6, 3, 2, 9, 7, 5, 4, 8},
		},
	}

	expected := testBoard.CheckIsValidColumn(0)
	if !expected {
		t.Error("good column returning invalid")
	}

	expected = testBoard.CheckIsValidColumn(1)
	if expected {
		t.Error("bad column with duplicate numbers returning valid")
	}

	expected = testBoard.CheckIsValidColumn(2)
	if expected {
		t.Error("bad column with a 0 returning valid")
	}

	expected = testBoard.CheckIsValidColumn(3)
	if expected {
		t.Error("bad column with a 10 returning valid")
	}
}

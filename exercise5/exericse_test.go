package exercise5

import "testing"

func TestPass_Coord(t *testing.T) {
	for data := range dataProvider() {
		row, column := data.input.coord()
		if row != data.expectRow || column != data.expectColumn {
			t.Errorf("Expected %d, %d. %d, %d given.", data.expectRow, data.expectColumn, row, column)
		}
	}
}

type data struct {
	expectRow int
	expectColumn int
	input Pass
}

func dataProvider() chan data {
	ch := make(chan data)
	go func() {
		ch <- data {expectRow: 70, expectColumn: 7, input: parsePass("BFFFBBFRRR")}
		ch <- data {expectRow: 14, expectColumn: 7, input: parsePass("FFFBBBFRRR")}
		ch <- data {expectRow: 102, expectColumn: 4, input: parsePass("BBFFBBFRLL")}
		ch <- data {expectRow: 44, expectColumn: 5, input: parsePass("FBFBBFFRLR")}
		close(ch)
	}()

	return ch
}
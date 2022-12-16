package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix interface {
	Rows() [][]int
	Cols() [][]int
	Set(row int, col int, val int) bool
}

type matrix [][]int

func (m matrix) Rows() [][]int {
	rowsCopy := make([][]int, len(m))

	for rowIdx, row := range m {
		rowsCopy[rowIdx] = make([]int, len(row))
		copy(rowsCopy[rowIdx], row)
	}

	return rowsCopy
}

func (m matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))

	for _, row := range m {
		for colIdx, elem := range row {
			cols[colIdx] = append(cols[colIdx], elem)
		}
	}

	return cols
}

func (m matrix) Set(row int, col int, val int) bool {
	if row < 0 || row > len(m)-1 || col < 0 || col > len(m[0])-1 {
		return false
	}

	m[row][col] = val
	return true
}

func New(in string) (matrix, error) {

	fullRows := strings.Split(in, "\n")

	matrix := make([][]int, len(fullRows))

	for rowIdx, row := range fullRows {
		rowByElement := strings.Split(strings.TrimSpace(row), " ")

		if rowIdx != 0 && len(matrix[rowIdx-1]) != len(rowByElement) {
			return nil, errors.New("wrong row length")
		}

		matrix[rowIdx] = make([]int, len(rowByElement))

		for colIdx, elem := range rowByElement {
			converted, err := strconv.Atoi(elem)

			if err != nil {
				return nil, err
			}

			matrix[rowIdx][colIdx] = converted
		}
	}

	return matrix, nil
}

package matrix

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Matrix [][]int

type Pair [2]int

func New(s string) (*Matrix, error) {
	matrix := Matrix{}

	if s == "" {
		return &Matrix{}, nil
	}

	for _, rawRow := range strings.Split(s, "\n") {
		row := []int{}
		for _, rawValue := range strings.Split(rawRow, " ") {
			value, err := strconv.ParseInt(rawValue, 10, 32)
			if err != nil {
				return nil, err
			}
			row = append(row, int(value))
		}
		matrix = append(matrix, row)
	}
	return &matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	if len(*m) == 0 {
		return []Pair{}
	}

	results := []Pair{}

	rowMax := make(map[int]int)
	for r, row := range *m {
		rowMax[r] = slices.Max(row)
	}

	colMin := make(map[int]int)
	for c := range (*m)[0] {
		minVal := math.MaxInt
		for r := range *m {
			if val := (*m)[r][c]; val < minVal {
				minVal = val
			}
		}
		colMin[c] = minVal
	}

	for r, row := range *m {
		for c, val := range row {
			if val == rowMax[r] && val == colMin[c] {
				results = append(results, [2]int{r + 1, c + 1})
			}
		}
	}

	return results
}

package cryptosquare

import (
	"math"
	"unicode"
)

func normalise(m string) string {
	result := ""
	for _, char := range m {
		char := unicode.ToLower(char)

		isAlpha := 'a' <= char && char <= 'z'
		isNum := '0' <= char && char <= '9'
		if isAlpha || isNum {
			result = result + string(char)
		}
	}
	return result
}

func getSquareSize(n int) (int, int) {
	size := math.Sqrt(float64(n))
	nRows := int(math.Floor(size))
	nCols := int(math.Ceil(size))

	if nRows*nCols < n {
		nRows += 1
	}

	if nRows*nCols < n {
		nCols += 1
	}

	return nRows, nCols
}

func buildSquare(m string) [][]string {
	nRows, nCols := getSquareSize(len(m))

	square := [][]string{}
	for i := 0; i < nRows; i++ {
		square = append(square, make([]string, nCols))
		for j := 0; j < nCols; j++ {
			square[i][j] = " "
		}
	}

	i := 0
	row := 0
	col := 0
	for i < len(m) {
		if col > nCols-1 {
			row += 1
			col = 0
		}

		square[row][col] = string(m[i])

		i += 1
		col += 1
	}
	return square
}

func readSquare(sq [][]string) string {
	result := ""

	for c := 0; c < len(sq[0]); c++ {
		for r := 0; r < len(sq); r++ {
			result += string(sq[r][c])
		}

		result += " "
	}

	result = result[:len(result)-1]
	return result
}

func Encode(pt string) string {
	m := normalise(pt)

	if len(m) == 0 {
		return m
	}

	sq := buildSquare(m)
	result := readSquare(sq)

	return result
}

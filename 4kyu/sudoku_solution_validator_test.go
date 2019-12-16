package _kyu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Validate3x3(m [][]int, s int, e int, size int) bool {
	hor, ver := 0, 0
	for i := s; i < s+size; i++ {
		for j := e; j < e+size; j++ {
			hor += m[i][j]
			ver += m[j][i]
		}
	}
	// This code can have some hole...
	if hor != 45 || ver != 45 {
		return false
	}
	return true
}

func ValidateSolution(m [][]int) bool {
	// Check horizontal, vertical
	for i := 0; i < 9; i++ {
		hor, ver := 0, 0
		for j := 0; j < 9; j++ {
			hor += m[i][j]
			ver += m[j][i]
		}
		if hor != 45 || ver != 45 {
			return false
		}
	}
	// Check 3x3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !Validate3x3(m, i, j, 3) {
				return false
			}
		}
	}
	return true
}

func TestSudokuSolutionValidator(t *testing.T) {

	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	var testFalse = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 0, 3, 4, 8},
		{1, 0, 0, 3, 4, 2, 5, 6, 0},
		{8, 5, 9, 7, 6, 1, 0, 2, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 0, 1, 5, 3, 7, 2, 1, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 0, 0, 4, 8, 1, 1, 7, 9},
	}

	var testFalse2 = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{2, 3, 1, 5, 6, 4, 8, 9, 7},
		{3, 1, 2, 6, 4, 5, 9, 7, 8},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{5, 6, 4, 8, 9, 7, 2, 3, 1},
		{6, 4, 5, 9, 7, 8, 3, 1, 2},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{8, 9, 7, 2, 3, 1, 5, 6, 4},
		{9, 7, 8, 3, 1, 2, 6, 4, 5},
	}

	assert.Equal(t, ValidateSolution(testTrue), true)
	assert.Equal(t, ValidateSolution(testFalse), false)
	assert.Equal(t, ValidateSolution(testFalse2), false)
}

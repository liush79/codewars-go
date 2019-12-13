package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func RowSumOddNumbers(n int) (res int) {
	res = 1
	for i, j := 1, 2; i < n; i, j = i+1, j+2 {
		res += j
	}
	base := res
	for i, j := 1, 2; i < n; i, j = i+1, j+2 {
		res += base + j
	}
	return
}

func TestSumOfOddNumbers(t *testing.T) {
	assert.Equal(t, RowSumOddNumbers(1), 1)
	assert.Equal(t, RowSumOddNumbers(2), 8)
	assert.Equal(t, RowSumOddNumbers(5), 125)
	assert.Equal(t, RowSumOddNumbers(13), 2197)
	assert.Equal(t, RowSumOddNumbers(19), 6859)
	assert.Equal(t, RowSumOddNumbers(41), 68921)
	assert.Equal(t, RowSumOddNumbers(42), 74088)
	assert.Equal(t, RowSumOddNumbers(74), 405224)
}

package _kyu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func PositiveSum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}
	return sum
}

func TestSumOfPositive(t *testing.T) {
	assert.Equal(t, PositiveSum([]int{1, -2, 3, 4, 5}), 13)
}

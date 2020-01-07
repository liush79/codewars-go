package _kyu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsDivisible(n, x, y int) bool {
	return n % x == 0 && n % y == 0
}

func TestIsNDivisible(t *testing.T) {
	assert.Equal(t, IsDivisible(3, 3, 4),false)
	assert.Equal(t, IsDivisible(12, 3, 4),true)
	assert.Equal(t, IsDivisible(8, 3, 4),false)
	assert.Equal(t, IsDivisible(48, 3, 4),true)
	assert.Equal(t, IsDivisible(100, 5, 10),true)
	assert.Equal(t, IsDivisible(100, 5, 3),false)
	assert.Equal(t, IsDivisible(4, 4, 2),true)
	assert.Equal(t, IsDivisible(5, 2, 3),false)
	assert.Equal(t, IsDivisible(17, 17, 17),true)
	assert.Equal(t, IsDivisible(17, 1, 17),true)
}

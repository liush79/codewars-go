package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func EquableTriangle(a, b, c int) bool {
	length := a + b + c
	d := length / 2
	return d*(d-a)*(d-b)*(d-c) == length*length
}

func TestCheckTriangle(t *testing.T) {
	assert.Equal(t, EquableTriangle(5, 12, 13), true)
	assert.Equal(t, EquableTriangle(2, 3, 4, ), false)
}

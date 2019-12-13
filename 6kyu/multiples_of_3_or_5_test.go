package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Multiple3And5(number int) (sum int) {
	for i := 1; i < number; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return
}

func TestMultipleOf3Or5(t *testing.T) {
	assert.Equal(t, Multiple3And5(10), 23)
	assert.Equal(t, Multiple3And5(20), 78)
}

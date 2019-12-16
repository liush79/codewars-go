package kata

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// solution 보니까.. best, clever 방법은
// strconv 안쓰고 했음 10 보다 크냐 작으냐로
func DigitalRoot(n int) int {
	if n / 10 < 1 {
		return n
	}
	sum := 0
	for _, sn := range strconv.Itoa(n) {
		num, _ := strconv.Atoi(string(sn)); sum += num
	}
	return DigitalRoot(sum)
}

func TestSumOfDigit(t *testing.T) {
	assert.Equal(t, DigitalRoot(16), 7)
	assert.Equal(t, DigitalRoot(195), 6)
	assert.Equal(t, DigitalRoot(992), 2)
	assert.Equal(t, DigitalRoot(167346), 9)
	assert.Equal(t, DigitalRoot(0), 0)
}

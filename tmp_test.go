package kata

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Is_valid_ip(ip string) bool {
	nums := strings.Split(ip, ".")
	if len(nums) != 4 {
		return false
	}
	for _, num := range nums {
		if len(num) != 1 && strings.HasPrefix(num, "0") {
			return false
		}
		if n, err:= strconv.Atoi(num); err != nil {
			return false
		} else if 0 > n || n > 255 {
			return false
		}
	}
	return true
}
func TestTmp(t *testing.T) {
	assert.Equal(t, Is_valid_ip("12.255.56.1"),true)
	assert.Equal(t, Is_valid_ip(""),false)
	assert.Equal(t, Is_valid_ip("abc.def.ghi.jkl"),false)
	assert.Equal(t, Is_valid_ip("123.456.789.0"),false)
	assert.Equal(t, Is_valid_ip("12.34.56"),false)
	assert.Equal(t, Is_valid_ip("12.34.56 .1"),false)
	assert.Equal(t, Is_valid_ip("12.34.56.-1"),false)
	assert.Equal(t, Is_valid_ip("123.045.067.089"),false)
	assert.Equal(t, Is_valid_ip("127.1.1.0"),true)
	assert.Equal(t, Is_valid_ip("0.0.0.0"),true)
	assert.Equal(t, Is_valid_ip("0.34.82.53"),true)
	assert.Equal(t, Is_valid_ip("192.168.1.300"),false)
}

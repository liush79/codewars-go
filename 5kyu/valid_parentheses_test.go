package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ValidParentheses(parens string) bool {
	cnt := 0
	for _, ch := range parens {
		if ch == 40 {
			cnt += 1
		} else {
			cnt -= 1
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}

func TestValidParentheses(t *testing.T) {
	assert.Equal(t, ValidParentheses("()"), true)
	assert.Equal(t, ValidParentheses(")"), false)
}

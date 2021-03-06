package kata

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func ValidParentheses(parens string) bool {
	if !strings.HasPrefix(parens, "(") {
		return false
	}
	cnt := 0
	for _, ch := range parens {
		if ch == 123 {
			cnt += 1
		} else {
			cnt -= 1
		}
	}
	return cnt == 0
}

func TestTmp(t *testing.T) {
	assert.Equal(t, ValidParentheses("()"),true)
	assert.Equal(t, ValidParentheses(")"),false)
}

package kata

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)


// 조금만 더 생각해봐라
// for 문 한번으로도 가능하다
func duplicate_count(s1 string) (res int) {
	m := map[int32]int{}
	for _, s := range strings.ToLower(s1) {
		m[s] += 1
	}
	for _, val := range m {
		if val != 1 {
			res += 1
		}
	}
	return
}


func TestCountingDuplicates(t *testing.T) {
	assert.Equal(t, duplicate_count("abcde"), 0)
	assert.Equal(t, duplicate_count("abcdea"), 1)
	assert.Equal(t, duplicate_count("abcdeaB11"), 3)
	assert.Equal(t, duplicate_count("indivisibility"), 1)
}

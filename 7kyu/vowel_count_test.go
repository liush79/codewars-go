package _kyu

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func GetCount(str string) (count int) {
	for i := 0; i < len(str); i++ {
		if strings.ContainsAny("aeiou", string(str[i])) {
			count++
		}
	}
	return
}

func TestVowelCount(t *testing.T) {
	assert.Equal(t, GetCount("abracadabra"), 5)
}
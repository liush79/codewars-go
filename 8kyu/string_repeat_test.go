package _kyu

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)


func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func TestVowelCount(t *testing.T) {
	assert.Equal(t, RepeatStr(4, "a"), "aaaa")
	assert.Equal(t, RepeatStr(3, "hello "),"hello hello hello ")
	assert.Equal(t, RepeatStr(2, "abc"),"abcabc")
}

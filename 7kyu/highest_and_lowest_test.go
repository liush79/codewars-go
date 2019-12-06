package _kyu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func HighAndLow(in string) string {
	instr := strings.Split(in, " ")
	num, _ := strconv.Atoi(instr[0])
	min, max := num, num
	for i := range instr {
		num, _ = strconv.Atoi(instr[i])
		if max < num {
			max = num
		} else if min > num {
			min = num
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}

func TestHighAndLow(t *testing.T) {
	assert.Equal(t, HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"), "42 -9")
}

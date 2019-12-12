package _kyu_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)


func getString(s string, length int) (res string) {
	if length != 1 {
		res += "-"
	}

	for i := 0; i < length; i++ {
		if i == 0 {
			res += strings.ToUpper(s)
		} else {
			res += s
		}
	}
	return
}

func Accum(s string) (result string) {
	for i := 0; i < len(s); i++ {
		result += getString(strings.ToLower(string(s[i])), i + 1)
	}
	return
}

func dotest(t *testing.T, s string, exp string) {
	var ans = Accum(s)
	assert.Equal(t, ans, exp)
}

func TestAccum(t *testing.T) {
	dotest(t, "ZpglnRxqenU", "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu")
	dotest(t, "NyffsGeyylB", "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb")
	dotest(t, "MjtkuBovqrU", "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu")
	dotest(t, "EvidjUnokmM", "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm")
	dotest(t, "HbideVbxncC", "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc")
}

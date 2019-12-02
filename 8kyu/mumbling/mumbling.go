package mumbling

import (
	"strings"
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
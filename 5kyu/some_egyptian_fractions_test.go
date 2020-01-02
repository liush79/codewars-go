package kata

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"strings"
	"testing"
)

func Decompose(s string) []string {
	res := make([]string, 0)
	numer, demon := 0, 0
	if strings.Contains(s, "/") {
		numer, _ = strconv.Atoi(strings.Split(s, "/")[0])
		demon, _ = strconv.Atoi(strings.Split(s, "/")[1])
	} else if strings.Contains(s, ".") {
		numer, _ = strconv.Atoi(strings.Split(s, ".")[1])
		demon = int(math.Pow10(len(strings.Split(s, ".")[1])))
		if strings.Split(s, ".")[0] != "0" {
			res = append(res, strings.Split(s, ".")[0])
		}
	} else {
		if s == "0" { return res }
		res = append(res, s)
		return res
	}

	if numer >= demon {
		res = append(res, fmt.Sprintf("%d", numer / demon))
		numer = numer % demon
	} else if numer > 0 && demon % numer == 0 {
		demon = demon / numer
		numer = 1
	}

	for demon > 0 && numer > 0  {
		nextNum := demon / numer + 1
		if (nextNum - 1) * numer == demon {
			nextNum -= 1
		}
		res = append(res, fmt.Sprintf("1/%d", nextNum))
		numer = numer * nextNum - demon
		demon *= nextNum
	}

	return res
}

func TestSomeEgyptianFractions(t *testing.T) {
	assert.Equal(t, Decompose("21/23"), []string{"1/2", "1/3", "1/13", "1/359", "1/644046"})
	assert.Equal(t, Decompose("12/4"), []string{"3"})
	assert.Equal(t, Decompose("0.66"), []string{"1/2", "1/7", "1/59", "1/5163", "1/53307975"})
	assert.Equal(t, Decompose("0"), []string{})
	assert.Equal(t, Decompose("1"), []string{"1"})
	assert.Equal(t, Decompose("1.25"), []string{"1", "1/4"})
	assert.Equal(t, Decompose("50/4187"), []string{"1/84", "1/27055", "1/1359351420"})
}

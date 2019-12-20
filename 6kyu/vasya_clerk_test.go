package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Tickets(peopleInLine []int) string {
	pocket := map[int]int{25: 0, 50: 0, 100: 0}
	for _, val := range peopleInLine {
		switch val {
		case 50:
			pocket[25] -= 1
		case 100:
			pocket[25] -= 1
			if pocket[50] > 0 {
				pocket[50] -= 1
			} else {
				pocket[25] -= 2
			}
		}
		pocket[val] += 1
		if pocket[25] < 0 || pocket[50] < 0 {
			return "NO"
		}
	}
	return "YES"
}

func TestTickets(t *testing.T) {
	assert.Equal(t, Tickets([]int{25, 25, 50}), "YES")
	assert.Equal(t, Tickets([]int{25, 100}), "NO")
	assert.Equal(t, Tickets([]int{25, 25, 50, 50, 100}), "NO")
	assert.Equal(t, Tickets([]int{25, 25, 50, 50, 25, 100}), "YES")
	assert.Equal(t, Tickets([]int{25, 50, 25, 25, 25, 25, 25, 50, 25, 25, 100, 25, 25, 25, 25, 100, 25, 100, 25, 25}), "YES")
}

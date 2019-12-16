package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func DirReduc(arr []string) []string {
	for {
		find := false
		for i := 0; i < len(arr)-1; i++ {
			switch arr[i] {
			case "NORTH": if arr[i+1] == "SOUTH" { find = true }
			case "SOUTH": if arr[i+1] == "NORTH" { find = true }
			case "EAST": if arr[i+1] == "WEST" { find = true }
			case "WEST": if arr[i+1] == "EAST" { find = true }
			}
			if find {
				arr = append(arr[:i], arr[i+2:]...)
				break
			}
		}
		if !find {
			return arr
		}
	}
}

func TestDirectionsReduction(t *testing.T) {
	var a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	assert.Equal(t, DirReduc(a), []string{"WEST"})
	a = []string{"NORTH", "WEST", "SOUTH", "EAST"}
	assert.Equal(t, DirReduc(a), []string{"NORTH", "WEST", "SOUTH", "EAST"})
	a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
	assert.Equal(t, DirReduc(a), []string{"NORTH"})
}

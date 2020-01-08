package kata

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func FindOutlier(integers []int) int {
	var odd, even *int
	for i, integer := range integers {
		if even != nil && odd != nil {
			if integer % 2 == 0 {
				return *odd
			}
			return *even
		}
		if integer % 2 == 0 {
			even = &integers[i]
		} else {
			odd = &integers[i]
		}
	}
	return integers[len(integers)-1]
}

func TestFindOutlier(t *testing.T) {
	assert.Equal(t, FindOutlier([]int{2, 6, 8, -10, 3}),3)
	assert.Equal(t, FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781}),206847684)
	assert.Equal(t, FindOutlier([]int{math.MaxInt32, 0, 1}),0)
}

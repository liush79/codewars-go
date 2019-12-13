package _kyu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func combat(health, damage float64) float64 {
	if health < damage {
		return 0
	}
	return health - damage
}

func TestGrasshopper(t *testing.T) {
	assert.Equal(t, combat(100.0, 50.0), 50.0)
}

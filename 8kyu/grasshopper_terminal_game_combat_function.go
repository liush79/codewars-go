package _kyu

import (
	"fmt"
)

func combat(health, damage float64) float64 {
	if health < damage {
		return 0
	}
	return health - damage
}

func main() {
	fmt.Println(combat(100.0, 50.0))
}

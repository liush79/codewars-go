package main

import "fmt"

func PositiveSum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}
	return sum
}

func main() {
	fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}))
}

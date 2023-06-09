package main

import "fmt"

func sum(nums ...int) int { // HL
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2))    // HL
	fmt.Println(sum(1, 2, 3)) // HL

	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...)) // HL
}

package slices

import "fmt"

func Print(numbers []int) {
	for _, number := range numbers {
		fmt.Printf("%d,", number)
	}
	fmt.Println()
}

func Change(numbers []int) {
	for i, number := range numbers {
		numbers[i] = number + 1
	}
}
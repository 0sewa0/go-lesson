package arrays

import "fmt"

func Print(fiveNumbers [5]int) {
	for _, number := range fiveNumbers {
		fmt.Printf("%d,", number)
	}
	fmt.Println()
}

func Change(fiveNumbers [5]int) {
	for i, number := range fiveNumbers {
		fiveNumbers[i] = number + 1
	}
}
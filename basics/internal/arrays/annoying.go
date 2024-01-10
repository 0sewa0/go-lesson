package arrays

import "fmt"

func Print(fiveNumbers [5]int) {
	for _, number := range fiveNumbers {
		fmt.Println(number)
	}
}
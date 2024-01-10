package main

import (
	"fmt"

	"go.lesson/basics/internal/arrays"
	_ "go.lesson/basics/internal/inits"
	"go.lesson/basics/internal/returns"
)

func onlySlices() {
	numbers := [5]int{1, 2, 3, 4}
	// numbers := []int{1,2,3}
	// numbers := make([]int, 5)
	arrays.Print(numbers)
}

func main() {
	fmt.Println("Go basics")

	onlySlices()

	what, is, this := returns.Annoying()
	fmt.Printf("%s, %s, %d\n", what, is, this)

	animal := returns.NotAnnoying()
	fmt.Println(animal)

}

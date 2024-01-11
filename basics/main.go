package main

import (
	"fmt"

	"go.lesson/basics/internal/arrays"
	_ "go.lesson/basics/internal/inits"
	"go.lesson/basics/internal/returns"
	"go.lesson/basics/internal/slices"
)

func onlySlices() {
	numbers := [5]int{1, 2, 3, 4}
	// numbers := []int{1,2,3}
	// numbers := make([]int, 5)
	arrays.Print(numbers)
	arrays.Change(numbers)
	arrays.Print(numbers)

	slice := numbers[:]
	slices.Print(slice)
	slices.Change(slice)
	slices.Print(slice)
	slice = append(slice, 1)


}

func main() {
	fmt.Println("Go basics")

	onlySlices()

	what, is, this := returns.Annoying()
	fmt.Printf("%s, %s, %d\n", what, is, this)

	animal := returns.NotAnnoying()
	fmt.Println(animal)
}

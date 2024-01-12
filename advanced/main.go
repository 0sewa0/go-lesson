package main

import (
	"fmt"
	"go.lesson/advanced/internal/snowman"
	// "go.lesson/basics/internal/slices" // Would this work?
)

func main() {
	methods()
}

func methods() {

	// Constructors with optional params

	// Value receiver - Builder Pattern
	builder := snowman.NewBuilder()
	builderMan := builder.
		WithName("build-man").
		WithCrookedNose().
		WithCarrotNose().
		WithSize(2).
		WithNoseLength(3).
		WithStoneEyes().
		WithEyeSize(1). // WithEmptyEyes(). // Would this work?
		Build()
	fmt.Println(builderMan)

	// Pointer receiver - Options Pattern
	optionsMan := snowman.New(
		"options-man",
		snowman.Size(11),
		snowman.MetalNose(),
		snowman.CrookedNose(),
		snowman.NoseLength(6),
		snowman.WithCoalEyes(),
		snowman.EyeSize(3),
	)
	fmt.Println(optionsMan)
}

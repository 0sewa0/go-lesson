package main

import (
	"errors"
	"fmt"
	"go.lesson/advanced/internal/concurrency"
	"go.lesson/advanced/internal/snowman"
	// "go.lesson/basics/internal/slices" // Would this work?
)

// FYI: https://codingchallenges.fyi/blog/learn-go/
// this is not based on it, but just as an extra
func main() {
	methods()
	types()
	concurrent()
}

func methods() {

	// Constructors with optional params
	fmt.Println("---About methods---")

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

func types() {
	fmt.Println("\n\n---About types---")

	godzilla := snowman.New("GODZILLA", snowman.Size(99999999), snowman.NoseLength(2))
	pinocchio := snowman.New("PINOCCHIO", snowman.Size(1), snowman.NoseLength(100000))
	teddyBear := snowman.NewBuilder().WithName("🧸").WithSize(2).WithNoseLength(2).Build()

	snowThings := []snowman.MelterHugger{
		godzilla,
		pinocchio,
		&teddyBear, // ???
	}

	isPirate := false
	for _, thing := range snowThings {
		err := thing.Hug(10) // what about "Range uses values instead of references"?
		if err != nil {
			switch err.(type) { // why is Goland complaining?
			case snowman.GotStabbedError:
				isPirate = true
			case snowman.TooSmallToHugError:
				fmt.Printf("😢\n")
			}
			fmt.Print(err)
		}
	}
	fmt.Printf("Are you a pirate? %v\n", isPirate)

	// Which will work?
	//meltAll(snowThings)
	//meltAll(snowThings.([]snowman.Melter))
	var melters []snowman.Melter
	for _, thing := range snowThings {
		melters = append(melters, thing.(snowman.Melter))
	}
	meltAll(melters)
}

func meltAll(things []snowman.Melter) {
	fmt.Println("\n-Start melt-")
	startTemp := -4
	melted := 0
MELT:
	for melted != len(things) { // PLS, don't do this. 😅
		startTemp += 1
		for _, thing := range things {
			err := thing.Melt(startTemp)

			var tooCold snowman.TooColdToMeltError
			if errors.As(err, &tooCold) {
				fmt.Printf("It is not cold enough to melt %s, (current temp: %d) \n", tooCold.Name, startTemp)
				continue MELT
			}
			melted += 1
		}
	}

}

// don't use interface{}, use any (it's the same)
func usingAny(thing any) {
	_, ok := thing.(snowman.Melter)
	if ok {
		fmt.Println("Its a Melter")
	}

	_, ok = thing.(snowman.Hugger)
	if ok {
		fmt.Println("Its a Hugger")
	}
}

func concurrent() {
	fmt.Println("\n\n---About concurrency---")

	fmt.Println("\n-With Mutex-")
	counter := &concurrency.CharCounter{}
	fmt.Println(counter.Count('a', "vjhasbhdbahbdakxbnkajbkahjbkanjkhabxjkhabkabvjvajhkabkljabkjhbkjabjhagvjabjavjabj"))

	fmt.Println("\n-With Channels-")

	count := concurrency.CountChars('a', "vjhasbhdbahbdakxbnkajbkahjbkanjkhabxjkhabkabvjvajhkabkljabkjhbkjabjhagvjabjavjabj")
	fmt.Println(count)

	fmt.Println("\n-About Contexts-")
	concurrency.ContextMaze()
}

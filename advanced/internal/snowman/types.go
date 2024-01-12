package snowman

import "fmt"

const (
	maxNameLength = 10
)

type Melter interface {
	Melt(int) error
}

type TooColdToMeltError string

func (cold TooColdToMeltError) Error() string {
	return fmt.Sprintf("That is too cold to melt anything! %s lives another day", cold)
}

// ---

type Hugger interface {
	Hug(int) error
}

type TooSmallToHugError string

func (small TooSmallToHugError) Error() string {
	return fmt.Sprintf("You are too small to hug %s! Disappointing", small)
}

type GotStabbedError string

func (stab GotStabbedError) Error() string {
	return fmt.Sprintf("You got stabbed in the eye by %s's nose! You are a pirate now", stab)
}

// ---

type nose struct {
	material NoseMaterial
	length   uint
	crooked  bool
}

type eye struct {
	material EyeMaterial
	size     uint
}

type snowman struct {
	name string
	size uint
	nose nose
	eyes eye
}

func (s snowman) Melt(temp int) error {
	if temp < 0 {
		return TooColdToMeltError(s.name)
	}
	fmt.Printf("You melted %s, how could you!?", s.name)
	return nil
}

func (s snowman) Hug(armLength uint) error {
	if s.size < s.nose.length {
		return GotStabbedError(s.name)
	}
	if armLength < s.size/2 {
		return TooSmallToHugError(s.name)
	}
	fmt.Printf("You hugged %s, nice.", s.name)
	return nil
}

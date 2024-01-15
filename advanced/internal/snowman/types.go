package snowman

import "fmt"

const (
	maxNameLength = 10
)

type Melter interface {
	Melt(int) error
	IsMelted() bool
}

type TooColdToMeltError struct {
	Name string
}

func (cold TooColdToMeltError) Error() string {
	return fmt.Sprintf("That is too cold to melt anything! %s lives another day\n", cold.Name)
}

// ---

type Hugger interface {
	Hug(uint) error
	IsHugged() bool
}

type TooSmallToHugError struct {
	Name string
}

func (small TooSmallToHugError) Error() string {
	return fmt.Sprintf("You are too small to hug %s! Disappointing\n", small.Name)
}

type GotStabbedError struct {
	Name string
}

func (stab GotStabbedError) Error() string {
	return fmt.Sprintf("You got stabbed in the eye by %s's nose! You are a pirate now\n", stab.Name)
}

// ---

type MelterHugger interface {
	Melter
	Hugger
}

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

	melted bool
	hugged bool
}

func (s *snowman) Melt(temp int) error {
	if temp < 0 {
		return TooColdToMeltError{s.name}
	}
	fmt.Printf("You melted %s, how could you!?\n", s.name)
	s.melted = true
	return nil
}

func (s *snowman) IsMelted() bool {
	return s.melted
}

func (s *snowman) Hug(armLength uint) error {
	if s.size < s.nose.length {
		return GotStabbedError{s.name}
	}
	if armLength < s.size/2 {
		return TooSmallToHugError{s.name}
	}
	fmt.Printf("You hugged %s, nice.\n", s.name)
	s.hugged = true
	return nil
}

func (s *snowman) IsHugged() bool {
	return s.hugged
}

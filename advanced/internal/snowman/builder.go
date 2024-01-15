package snowman

import "strings"

type Builder struct {
	snowman snowman
}

func (b Builder) WithName(name string) Builder {
	if len(name) > maxNameLength {
		name = name[:maxNameLength]
	}
	b.snowman.name = strings.ToTitle(name)
	return b
}

func (b Builder) WithSize(size uint) Builder {
	b.snowman.size = size
	return b
}

func (b Builder) WithWoodNose() Builder {
	b.snowman.nose.material = WOOD
	return b
}

func (b Builder) WithCarrotNose() Builder {
	b.snowman.nose.material = CARROT
	return b
}

func (b Builder) WithCrookedNose() Builder {
	b.snowman.nose.crooked = true
	return b
}

func (b Builder) WithNoseLength(length uint) Builder {
	b.snowman.nose.length = length
	return b
}

func (b Builder) WithStoneEyes() Builder {
	b.snowman.eyes.material = STONE
	return b
}

func (b Builder) WithEyeSize(size uint) Builder {
	b.snowman.eyes.size = size
	return b
}

// Mixing receiver types == Bad
func (b *Builder) WithEmptyEyes() *Builder {
	b.snowman.eyes.material = EMPTY
	return b
}

func (b Builder) Build() snowman {
	return b.snowman
} // Why not return a reference?

func NewBuilder() Builder {
	return Builder{}
}

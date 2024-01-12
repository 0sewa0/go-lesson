package snowman

import "strings"

type Option func(*snowman)

func New(name string, options ...Option) *snowman { // Improve ?
	if len(name) > maxNameLength {
		name = name[:maxNameLength]
	}

	snow := &snowman{
		name: strings.ToTitle(name),
	}

	for _, option := range options {
		option(snow)
	}
	return snow
}

func Size(size uint) Option {
	return func(snowman *snowman) {
		snowman.size = size
	}
}

func MetalNose() Option {
	return func(snowman *snowman) {
		snowman.nose.material = METAL
	}
}

func PlasticNose() Option {
	return func(snowman *snowman) {
		snowman.nose.material = PLASTIC
	}
}

func NoseLength(length uint) Option {
	return func(snowman *snowman) {
		snowman.nose.length = length
	}
}

func CrookedNose() Option {
	return func(snowman *snowman) {
		snowman.nose.crooked = true
	}
}

func WithCoalEyes() Option {
	return func(snowman *snowman) {
		snowman.eyes.material = COAL
	}
}

func EyeSize(size uint) Option {
	return func(snowman *snowman) {
		snowman.eyes.size = size
	}
}

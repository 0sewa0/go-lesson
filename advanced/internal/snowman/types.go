package snowman

const (
	maxNameLength = 10
)

type Melter interface {
	Melt(int) error
}

type Hugger interface {
	Hug(int) error
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
}

func (s snowman) Melt(temp int) error {
	return nil // TODO: Add TooColdError
}

func (s snowman) Hug(armLength int) error {
	return nil // TODO: Add TooSmallError
}

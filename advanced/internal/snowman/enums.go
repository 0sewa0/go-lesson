package snowman

import "fmt"

type NoseMaterial int

var _ fmt.Stringer = NoseMaterial(0)

const (
	WOOD NoseMaterial = iota
	METAL
	CARROT
	PLASTIC
)

func (material NoseMaterial) String() string {
	switch material {
	case WOOD:
		return "Wooden"
	case METAL:
		return "Metallic"
	case CARROT:
		return "Carrot"
	case PLASTIC:
		return "Plastic"
	default:
		return "Unknown"
	}
}

type EyeMaterial int

var _ fmt.Stringer = EyeMaterial(0)

const (
	EMPTY EyeMaterial = iota
	COAL
	STONE
)

func (material EyeMaterial) String() string {
	switch material {
	case EMPTY:
		return "Empty"
	case COAL:
		return "Coal"
	case STONE:
		return "Stone"
	default:
		return "Unknown"
	}
}

//func conversion() { // Would this work?
//	var material int
//	nose := material.(NoseMaterial)
//}

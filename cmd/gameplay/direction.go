package gameplay

import "github.com/veandco/go-sdl2/sdl"

type KurinDirection uint8

const KurinDirectionNorth = KurinDirection(1)
const KurinDirectionEast = KurinDirection(2)
const KurinDirectionSouth = KurinDirection(0)
const KurinDirectionWest = KurinDirection(3)

func GetFacingDirection(from sdl.FPoint, to sdl.FPoint) KurinDirection {
	if to.Y < from.Y {
		return KurinDirectionNorth
	}
	if to.X > from.X {
		return KurinDirectionEast
	}
	if to.Y > from.Y {
		return KurinDirectionSouth
	}
	if to.X < from.X {
		return KurinDirectionWest
	}

	return KurinDirectionNorth
}

package common

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinDirection uint8

const (
	KurinDirectionNorth = KurinDirection(1)
	KurinDirectionEast  = KurinDirection(2)
	KurinDirectionSouth = KurinDirection(0)
	KurinDirectionWest  = KurinDirection(3)
)

func GetFacingDirection(from sdl.Point, to sdl.Point) KurinDirection {
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

func GetFacingDirectionF(from sdl.FPoint, to sdl.FPoint) KurinDirection {
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

func GetPositionInDirectionV(from sdlutils.Vector3, direction KurinDirection) sdlutils.Vector3 {
	return sdlutils.Vector3{
		Base: GetPositionInDirection(from.Base, direction),
		Z:    from.Z,
	}
}

func GetPositionInDirectionFV(from sdlutils.Vector3, direction KurinDirection) sdlutils.FVector3 {
	return sdlutils.FVector3{
		Base: sdlutils.PointToFPoint(GetPositionInDirection(from.Base, direction)),
		Z:    from.Z,
	}
}

func GetPositionInDirectionFVCenter(from sdlutils.Vector3, direction KurinDirection) sdlutils.FVector3 {
	return sdlutils.FVector3{
		Base: sdlutils.PointToFPointCenter(GetPositionInDirection(from.Base, direction)),
		Z:    from.Z,
	}
}

func GetPositionInDirection(from sdl.Point, direction KurinDirection) sdl.Point {
	switch direction {
	case KurinDirectionNorth:
		return sdl.Point{
			X: from.X,
			Y: from.Y - 1,
		}
	case KurinDirectionEast:
		return sdl.Point{
			X: from.X + 1,
			Y: from.Y,
		}
	case KurinDirectionSouth:
		return sdl.Point{
			X: from.X,
			Y: from.Y + 1,
		}
	case KurinDirectionWest:
		return sdl.Point{
			X: from.X - 1,
			Y: from.Y,
		}
	}

	return sdl.Point{}
}

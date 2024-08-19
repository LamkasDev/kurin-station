package common

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type Direction uint8

const (
	DirectionNorth = Direction(1)
	DirectionEast  = Direction(2)
	DirectionSouth = Direction(0)
	DirectionWest  = Direction(3)
)

func GetFacingDirection(from sdl.Point, to sdl.Point) Direction {
	if to.Y < from.Y {
		return DirectionNorth
	}
	if to.X > from.X {
		return DirectionEast
	}
	if to.Y > from.Y {
		return DirectionSouth
	}
	if to.X < from.X {
		return DirectionWest
	}

	return DirectionNorth
}

func GetFacingDirectionF(from sdl.FPoint, to sdl.FPoint) Direction {
	if to.Y < from.Y {
		return DirectionNorth
	}
	if to.X > from.X {
		return DirectionEast
	}
	if to.Y > from.Y {
		return DirectionSouth
	}
	if to.X < from.X {
		return DirectionWest
	}

	return DirectionNorth
}

func GetPositionInDirectionV(from sdlutils.Vector3, direction Direction) sdlutils.Vector3 {
	return sdlutils.Vector3{
		Base: GetPositionInDirection(from.Base, direction),
		Z:    from.Z,
	}
}

func GetPositionInDirectionFV(from sdlutils.Vector3, direction Direction) sdlutils.FVector3 {
	return sdlutils.FVector3{
		Base: sdlutils.PointToFPoint(GetPositionInDirection(from.Base, direction)),
		Z:    from.Z,
	}
}

func GetPositionInDirectionFVCenter(from sdlutils.Vector3, direction Direction) sdlutils.FVector3 {
	return sdlutils.FVector3{
		Base: sdlutils.PointToFPointCenter(GetPositionInDirection(from.Base, direction)),
		Z:    from.Z,
	}
}

func GetPositionInDirection(from sdl.Point, direction Direction) sdl.Point {
	switch direction {
	case DirectionNorth:
		return sdl.Point{
			X: from.X,
			Y: from.Y - 1,
		}
	case DirectionEast:
		return sdl.Point{
			X: from.X + 1,
			Y: from.Y,
		}
	case DirectionSouth:
		return sdl.Point{
			X: from.X,
			Y: from.Y + 1,
		}
	case DirectionWest:
		return sdl.Point{
			X: from.X - 1,
			Y: from.Y,
		}
	}

	return sdl.Point{}
}

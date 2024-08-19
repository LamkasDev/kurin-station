package sdlutils

import (
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/arl/math32"
	"github.com/veandco/go-sdl2/sdl"
)

func FPointToPoint(point sdl.FPoint) sdl.Point {
	return sdl.Point{
		X: int32(point.X), Y: int32(point.Y),
	}
}

func FPointToPointFloored(point sdl.FPoint) sdl.Point {
	return sdl.Point{
		X: int32(math32.Floor(point.X)), Y: int32(math32.Floor(point.Y)),
	}
}

func AddFPoints(a sdl.FPoint, b sdl.FPoint) sdl.FPoint {
	return sdl.FPoint{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func SubtractFPoints(a sdl.FPoint, b sdl.FPoint) sdl.FPoint {
	return sdl.FPoint{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func MultiplyFPoints(a sdl.FPoint, b sdl.FPoint) sdl.FPoint {
	return sdl.FPoint{
		X: a.X * b.X,
		Y: a.Y * b.Y,
	}
}

func DivideFPoint(a sdl.FPoint, b float32) sdl.FPoint {
	return sdl.FPoint{
		X: a.X / b,
		Y: a.Y / b,
	}
}

func DivideFPoints(a sdl.FPoint, b sdl.FPoint) sdl.FPoint {
	return sdl.FPoint{
		X: a.X / b.X,
		Y: a.Y / b.Y,
	}
}

func CompareFPoints(a sdl.FPoint, b sdl.FPoint) bool {
	return a.X == b.X && a.Y == b.Y
}

func GetDistanceSimpleF(a sdl.FPoint, b sdl.FPoint) float32 {
	return mathutils.MaxFloat32(mathutils.AbsFloat32(a.X-b.X), mathutils.AbsFloat32(a.Y-b.Y))
}

func GetDistanceF(a sdl.FPoint, b sdl.FPoint) float32 {
	return mathutils.AbsFloat32(a.X-b.X) + mathutils.AbsFloat32(a.Y-b.Y)
}

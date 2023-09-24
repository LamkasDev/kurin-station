package sdlutils

import (
	"github.com/adam-lavrik/go-imath/i32"
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

func PointToFPoint(point sdl.Point) sdl.FPoint {
	return sdl.FPoint{
		X: float32(point.X), Y: float32(point.Y),
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

func DivideFPoints(a sdl.FPoint, b sdl.FPoint) sdl.FPoint {
	return sdl.FPoint{
		X: a.X / b.X,
		Y: a.Y / b.Y,
	}
}

func DivideFPointByFloat(a sdl.FPoint, b float32) sdl.FPoint {
	return sdl.FPoint{
		X: a.X / b,
		Y: a.Y / b,
	}
}

func GetDistanceSimple(a sdl.Point, b sdl.Point) int32 {
	return i32.Max(i32.Abs(a.X-b.X), i32.Abs(a.Y-b.Y))
}

func GetDistance(a sdl.Point, b sdl.Point) int32 {
	return i32.Abs(a.X-b.X) + i32.Abs(a.Y-b.Y)
}

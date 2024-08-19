package sdlutils

import (
	"math"

	"github.com/adam-lavrik/go-imath/i32"
	"github.com/arl/math32"
	"github.com/veandco/go-sdl2/sdl"
)

func PointToFPoint(point sdl.Point) sdl.FPoint {
	return sdl.FPoint{
		X: float32(point.X), Y: float32(point.Y),
	}
}

func PointToFPointCenter(point sdl.Point) sdl.FPoint {
	return sdl.FPoint{
		X: float32(point.X) + 0.5, Y: float32(point.Y) + 0.5,
	}
}

func AddPoints(a sdl.Point, b sdl.Point) sdl.Point {
	return sdl.Point{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func SubtractPoints(a sdl.Point, b sdl.Point) sdl.Point {
	return sdl.Point{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func MultiplyPoints(a sdl.Point, b sdl.Point) sdl.Point {
	return sdl.Point{
		X: a.X * b.X,
		Y: a.Y * b.Y,
	}
}

func DividePoint(point sdl.Point, n float32) sdl.Point {
	return sdl.Point{
		X: int32(float32(point.X) / n),
		Y: int32(float32(point.Y) / n),
	}
}

func DividePoints(a sdl.Point, b sdl.Point) sdl.Point {
	return sdl.Point{
		X: a.X / b.X,
		Y: a.Y / b.Y,
	}
}

func IsPointZero(a sdl.Point) bool {
	return a.X == 0 && a.Y == 0
}

func ComparePoints(a sdl.Point, b sdl.Point) bool {
	return a.X == b.X && a.Y == b.Y
}

func GetDistanceSimple(a sdl.Point, b sdl.Point) int32 {
	return i32.Max(i32.Abs(a.X-b.X), i32.Abs(a.Y-b.Y))
}

func GetDistance(a sdl.Point, b sdl.Point) int32 {
	return i32.Abs(a.X-b.X) + i32.Abs(a.Y-b.Y)
}

func RotatePoint(point sdl.Point, center sdl.Point, angle float32) sdl.Point {
	x1 := float32(point.X - center.X)
	y1 := float32(point.Y - center.Y)

	rad := -angle * (math.Pi / 180)
	x2 := x1*math32.Cos(rad) - y1*math32.Sin(rad)
	y2 := x1*math32.Sin(rad) + y1*math32.Cos(rad)

	return sdl.Point{
		X: int32(x2) + center.X,
		Y: int32(y2) + center.Y,
	}
}

package sdlutils

import (
	"math"

	"github.com/LamkasDev/kurin/cmd/common/mathutils"
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

func PointToFPointCenter(point sdl.Point) sdl.FPoint {
	return sdl.FPoint{
		X: float32(point.X) + 0.5, Y: float32(point.Y) + 0.5,
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

func GetDistanceSimpleF(a sdl.FPoint, b sdl.FPoint) float32 {
	return mathutils.MaxFloat32(mathutils.AbsFloat32(a.X-b.X), mathutils.AbsFloat32(a.Y-b.Y))
}

func GetDistance(a sdl.Point, b sdl.Point) int32 {
	return i32.Abs(a.X-b.X) + i32.Abs(a.Y-b.Y)
}

func GetDistanceF(a sdl.FPoint, b sdl.FPoint) float32 {
	return mathutils.AbsFloat32(a.X-b.X) + mathutils.AbsFloat32(a.Y-b.Y)
}

func RotatePoint(point sdl.Point, center sdl.Point, angle float32) sdl.Point {
	x1 := float32(point.X - center.X);
	y1 := float32(point.Y - center.Y);

	rad := -angle * (math.Pi/180);
	x2 := x1 * math32.Cos(rad) - y1 * math32.Sin(rad);
	y2 := x1 * math32.Sin(rad) + y1 * math32.Cos(rad);

	return sdl.Point{
		X: int32(x2) + center.X,
		Y: int32(y2) + center.Y,
	}
}

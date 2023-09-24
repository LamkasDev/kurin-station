package mathutils

import "github.com/veandco/go-sdl2/sdl"

func Lerp(a float32, b float32, f float32) float32 {
	return a*(1.0-f) + (b * f)
}

func LerpFPoint(a sdl.FPoint, b sdl.FPoint, f float32) sdl.FPoint {
	return sdl.FPoint{
		X: Lerp(a.X, b.X, f),
		Y: Lerp(a.Y, b.Y, f),
	}
}

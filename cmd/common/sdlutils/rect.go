package sdlutils

import "github.com/veandco/go-sdl2/sdl"

func FRectToRect(frect sdl.FRect) sdl.Rect {
	return sdl.Rect{
		X: int32(frect.X), Y: int32(frect.Y), W: int32(frect.W), H: int32(frect.H),
	}
}

func RectToFRect(rect sdl.Rect) sdl.FRect {
	return sdl.FRect{
		X: float32(rect.X), Y: float32(rect.Y), W: float32(rect.W), H: float32(rect.H),
	}
}

func AddRects(a sdl.Rect, b sdl.Rect) sdl.Rect {
	return sdl.Rect{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		W: a.W + b.W,
		H: a.H + b.H,
	}
}

func AddRectAndPoint(a sdl.Rect, b sdl.Point) sdl.Rect {
	return sdl.Rect{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		W: a.W,
		H: a.H,
	}
}

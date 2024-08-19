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

func ScaleRectCentered(rect sdl.Rect, n float32) sdl.Rect {
	return sdl.Rect{
		X: rect.X + int32((float32(rect.W)*(1-n))/2),
		Y: rect.Y + int32((float32(rect.H)*(1-n))/2),
		W: int32(float32(rect.W) * n),
		H: int32(float32(rect.H) * n),
	}
}

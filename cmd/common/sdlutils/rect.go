package sdlutils

import "github.com/veandco/go-sdl2/sdl"

func IsVectorInsideRect(pos sdl.Rect, rect sdl.Rect) bool {
	return pos.X >= rect.X && pos.Y >= rect.Y && pos.X <= rect.X+rect.W && pos.Y <= rect.Y+rect.H
}

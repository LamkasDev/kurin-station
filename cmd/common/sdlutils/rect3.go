package sdlutils

import "github.com/veandco/go-sdl2/sdl"

type Rect3 struct {
	Base sdl.Rect
	Z    uint8
}

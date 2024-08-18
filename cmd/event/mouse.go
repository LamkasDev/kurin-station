package event

import "github.com/veandco/go-sdl2/sdl"

type Mouse struct {
	PressedLeft  bool
	PressedRight bool

	PendingLeft   *sdl.Point
	PendingRight  *sdl.Point
	PendingScroll int32
	Delta         sdl.Point

	Cursor  sdl.SystemCursor
	Cursors map[sdl.SystemCursor]*sdl.Cursor
}

func NewMouse() Mouse {
	return Mouse{
		PressedLeft:   false,
		PressedRight:  false,
		PendingLeft:   nil,
		PendingRight:  nil,
		PendingScroll: 0,
		Delta:         sdl.Point{},
		Cursor:        sdl.SYSTEM_CURSOR_ARROW,
		Cursors: map[sdl.SystemCursor]*sdl.Cursor{
			sdl.SYSTEM_CURSOR_ARROW: sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_ARROW),
			sdl.SYSTEM_CURSOR_HAND:  sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_HAND),
		},
	}
}

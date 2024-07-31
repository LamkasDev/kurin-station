package event

import "github.com/veandco/go-sdl2/sdl"

type KurinMouse struct {
	PendingLeft  *sdl.Point
	PendingRight *sdl.Point
	Cursor sdl.SystemCursor
	Cursors map[sdl.SystemCursor]*sdl.Cursor
	Scroll       int32
}

func NewKurinMouse() KurinMouse {
	return KurinMouse{
		PendingLeft:  nil,
		PendingRight: nil,
		Cursor: sdl.SYSTEM_CURSOR_ARROW,
		Cursors: map[sdl.SystemCursor]*sdl.Cursor{
			sdl.SYSTEM_CURSOR_ARROW: sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_ARROW),
			sdl.SYSTEM_CURSOR_HAND: sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_HAND),
		},
		Scroll:       0,
	}
}

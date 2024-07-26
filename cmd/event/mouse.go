package event

import "github.com/veandco/go-sdl2/sdl"

type KurinMouse struct {
	PendingLeft  *sdl.Point
	LastLeft  *sdl.Point
	PendingRight *sdl.Point
	LastRight  *sdl.Point
	Scroll       int32
}

func NewKurinMouse() KurinMouse {
	return KurinMouse{
		PendingLeft:  nil,
		LastLeft: nil,
		PendingRight: nil,
		LastRight: nil,
		Scroll:       0,
	}
}

package event

import "github.com/veandco/go-sdl2/sdl"

type KurinMouse struct {
	PendingLeft  *sdl.Point
	PendingRight *sdl.Point
	Scroll       int32
}

func NewKurinMouse() KurinMouse {
	return KurinMouse{
		PendingLeft:  nil,
		PendingRight: nil,
		Scroll:       0,
	}
}

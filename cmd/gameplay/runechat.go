package gameplay

import (
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRunechat struct {
	Color   sdl.Color
	Message string
	Ticks   uint32
	Data    interface{}
}

func NewKurinRunechat(message string) *KurinRunechat {
	return &KurinRunechat{
		Color:   sdl.Color{R: 255, G: 255, B: 255},
		Message: message,
		Ticks:   360,
	}
}

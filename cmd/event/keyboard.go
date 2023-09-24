package event

import "github.com/veandco/go-sdl2/sdl"

type KurinKeyboard struct {
	Pressed   map[sdl.Keycode]bool
	Pending   *sdl.Keycode
	Input     string
	InputMode bool
}

func NewKurinKeyboard() KurinKeyboard {
	return KurinKeyboard{
		Pressed:   map[sdl.Keycode]bool{},
		Pending:   nil,
		Input:     "",
		InputMode: false,
	}
}

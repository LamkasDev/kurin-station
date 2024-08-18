package event

import "github.com/veandco/go-sdl2/sdl"

type Keyboard struct {
	Pressed   map[sdl.Keycode]bool
	Pending   *sdl.Keycode
	Input     string
	InputMode bool
}

func NewKeyboard() Keyboard {
	return Keyboard{
		Pressed:   map[sdl.Keycode]bool{},
		Pending:   nil,
		Input:     "",
		InputMode: false,
	}
}

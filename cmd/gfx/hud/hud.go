package hud

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinHUDElement struct {
	Path        string
	Hovered     bool
	GetPosition KurinHUDElementGetPosition
	Click       KurinHUDElementClick
}

type KurinHUDElementGetPosition func(window sdl.Point) sdl.Point
type KurinHUDElementClick func(game *gameplay.KurinGame)

var KurinHUDElementHandLeft = KurinHUDElement{
	GetPosition: func(window sdl.Point) sdl.Point {
		return sdl.Point{X: int32(float32(window.X) / 2), Y: window.Y - 72}
	},
	Click: func(game *gameplay.KurinGame) {
		if game.SelectedCharacter == nil {
			return
		}
		game.SelectedCharacter.ActiveHand = gameplay.KurinHandLeft
	},
}

var KurinHUDElementHandRight = KurinHUDElement{
	GetPosition: func(window sdl.Point) sdl.Point {
		return sdl.Point{X: int32(float32(window.X)/2) - 64, Y: window.Y - 72}
	},
	Click: func(game *gameplay.KurinGame) {
		if game.SelectedCharacter == nil {
			return
		}
		game.SelectedCharacter.ActiveHand = gameplay.KurinHandRight
	},
}

var KurinHUDElementPDA = KurinHUDElement{
	GetPosition: func(window sdl.Point) sdl.Point {
		return sdl.Point{X: window.X - 72, Y: 8}
	},
	Click: func(game *gameplay.KurinGame) {

	},
}

var KurinHUDElements = []*KurinHUDElement{&KurinHUDElementHandLeft, &KurinHUDElementHandRight, &KurinHUDElementPDA}

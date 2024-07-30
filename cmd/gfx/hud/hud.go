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
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: int32(float32(windowSize.X) / 2), Y: windowSize.Y - 72}
	},
	Click: func(game *gameplay.KurinGame) {
		if game.SelectedCharacter == nil {
			return
		}
		game.SelectedCharacter.ActiveHand = gameplay.KurinHandLeft
	},
}

var KurinHUDElementHandRight = KurinHUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: int32(float32(windowSize.X)/2) - 64, Y: windowSize.Y - 72}
	},
	Click: func(game *gameplay.KurinGame) {
		if game.SelectedCharacter == nil {
			return
		}
		game.SelectedCharacter.ActiveHand = gameplay.KurinHandRight
	},
}

var KurinHUDElementPDA = KurinHUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: windowSize.X - 72, Y: 8}
	},
	Click: func(game *gameplay.KurinGame) {

	},
}

var KurinHUDElements = []*KurinHUDElement{&KurinHUDElementHandLeft, &KurinHUDElementHandRight, &KurinHUDElementPDA}

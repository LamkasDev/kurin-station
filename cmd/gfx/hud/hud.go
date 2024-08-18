package hud

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinHUDElement struct {
	Path        string
	Hovered     bool
	GetPosition KurinHUDElementGetRect
	Click       KurinHUDElementClick
}

type (
	KurinHUDElementGetRect func(window sdl.Point) sdl.Point
	KurinHUDElementClick   func()
)

var KurinHUDElementHandLeft = KurinHUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: int32(float32(windowSize.X) / 2), Y: windowSize.Y - 72}
	},
	Click: func() {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return
		}
		gameplay.GameInstance.SelectedCharacter.ActiveHand = gameplay.KurinHandLeft
	},
}

var KurinHUDElementHandRight = KurinHUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: int32(float32(windowSize.X)/2) - 64, Y: windowSize.Y - 72}
	},
	Click: func() {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return
		}
		gameplay.GameInstance.SelectedCharacter.ActiveHand = gameplay.KurinHandRight
	},
}

var KurinHUDElementPDA = KurinHUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: windowSize.X - 72, Y: 8}
	},
	Click: func() {
	},
}

var KurinHUDElementCredit = KurinHUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: 0, Y: 76}
	},
	Click: func() {
	},
}

var KurinHUDElementGoals = KurinHUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: 8, Y: 8}
	},
	Click: func() {
	},
}

var KurinHUDElements = []*KurinHUDElement{&KurinHUDElementHandLeft, &KurinHUDElementHandRight, &KurinHUDElementPDA}

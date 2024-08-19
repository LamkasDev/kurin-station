package hud

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type HUDElement struct {
	Path        string
	Hovered     bool
	GetPosition HUDElementGetRect
	Click       HUDElementClick
}

type (
	HUDElementGetRect func(window sdl.Point) sdl.Point
	HUDElementClick   func()
)

var HUDElementHandLeft = HUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: int32(float32(windowSize.X) / 2), Y: windowSize.Y - 72}
	},
	Click: func() {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return
		}
		gameplay.GameInstance.SelectedCharacter.ActiveHand = gameplay.HandLeft
	},
}

var HUDElementHandRight = HUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: int32(float32(windowSize.X)/2) - 64, Y: windowSize.Y - 72}
	},
	Click: func() {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return
		}
		gameplay.GameInstance.SelectedCharacter.ActiveHand = gameplay.HandRight
	},
}

var HUDElementPDA = HUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: windowSize.X - 72, Y: 8}
	},
	Click: func() {
	},
}

var HUDElementCredit = HUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: 0, Y: 76}
	},
	Click: func() {
	},
}

var HUDElementGoals = HUDElement{
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: 8, Y: 8}
	},
	Click: func() {
	},
}

var HUDElements = []*HUDElement{&HUDElementHandLeft, &HUDElementHandRight, &HUDElementPDA}

package hud

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type HUDElement struct {
	Path        string
	Anchor      gfx.UIAnchor
	Scale       sdl.FPoint
	Hovered     bool
	GetPosition HUDElementGetRect
	Click       HUDElementClick
}

type (
	HUDElementGetRect func(window sdl.Point) sdl.Point
	HUDElementClick   func()
)

var HUDElementHandLeft = HUDElement{
	Path:   "hand_l",
	Anchor: gfx.UIAnchorBottomCenter,
	Scale:  sdl.FPoint{X: 2, Y: 2},
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: int32(gfx.RendererInstance.Context.WindowScale.X * 32), Y: int32(gfx.RendererInstance.Context.WindowScale.X * 16)}
	},
	Click: func() {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return
		}
		gameplay.GetInventory(gameplay.GameInstance.SelectedCharacter).ActiveHand = gameplay.HandLeft
	},
}

var HUDElementHandRight = HUDElement{
	Path:   "hand_r",
	Anchor: gfx.UIAnchorBottomCenter,
	Scale:  sdl.FPoint{X: 2, Y: 2},
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: -int32(gfx.RendererInstance.Context.WindowScale.X * 32), Y: int32(gfx.RendererInstance.Context.WindowScale.X * 16)}
	},
	Click: func() {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return
		}
		gameplay.GetInventory(gameplay.GameInstance.SelectedCharacter).ActiveHand = gameplay.HandRight
	},
}

var HUDElementPDA = HUDElement{
	Path:   "pda",
	Anchor: gfx.UIAnchorTopRight,
	Scale:  sdl.FPoint{X: 2, Y: 2},
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: 0, Y: 8}
	},
	Click: func() {
	},
}

var HUDElementGoals = HUDElement{
	Scale: sdl.FPoint{X: 2, Y: 2},
	GetPosition: func(windowSize sdl.Point) sdl.Point {
		return sdl.Point{X: int32(gfx.RendererInstance.Context.WindowScale.X * 8), Y: int32(gfx.RendererInstance.Context.WindowScale.X * 8)}
	},
	Click: func() {
	},
}

var HUDElements = []*HUDElement{&HUDElementHandLeft, &HUDElementHandRight, &HUDElementPDA}

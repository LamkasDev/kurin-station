package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"github.com/veandco/go-sdl2/sdl"
)

type KitsuneRendererFrameContext struct {
	LayoutPosition      sdl.Rect
	LayoutForgiveMargin node.KitsuneElementRectLTRB
	LayoutDirection     node.KitsuneLayoutDirection

	PageSize         sdl.Rect
	InspectedElement *node.KitsuneElement
}

func NewKitsuneRendererContext() *KitsuneRendererFrameContext {
	return &KitsuneRendererFrameContext{
		LayoutPosition:      sdl.Rect{},
		LayoutForgiveMargin: node.KitsuneElementRectLTRB{},
		LayoutDirection:     node.KitsuneLayoutDirectionRow,

		PageSize:         sdl.Rect{},
		InspectedElement: nil,
	}
}

type KitsuneRendererWindowContext struct {
	WindowSize    sdl.Rect
	MousePosition sdl.Rect
}

func NewKitsuneWindowContext() *KitsuneRendererWindowContext {
	return &KitsuneRendererWindowContext{
		WindowSize:    sdl.Rect{},
		MousePosition: sdl.Rect{},
	}
}

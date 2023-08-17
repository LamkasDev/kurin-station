package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKitsuneElementRect(renderer *KitsuneRenderer, element *node.KitsuneElement) sdl.Rect {
	return sdl.Rect{
		X: renderer.FrameContext.LayoutPosition.X,
		Y: renderer.FrameContext.LayoutPosition.Y,
		W: GetKitsuneElementVariableValue(element.Size.Width),
		H: GetKitsuneElementVariableValue(element.Size.Height),
	}
}

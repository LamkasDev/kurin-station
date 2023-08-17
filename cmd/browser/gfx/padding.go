package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/node"
)

func PrependKitsuneElementPadding(renderer *KitsuneRenderer, element *node.KitsuneElement) {
	renderer.FrameContext.LayoutPosition.X += GetKitsuneElementVariableValue(element.Padding.Left)
	renderer.FrameContext.LayoutPosition.Y += GetKitsuneElementVariableValue(element.Padding.Top)
	renderer.FrameContext.PageSize.H += GetKitsuneElementVariableValue(element.Padding.Top)
}

func AppendKitsuneElementPadding(renderer *KitsuneRenderer, element *node.KitsuneElement) {
	renderer.FrameContext.LayoutPosition.X += GetKitsuneElementVariableValue(element.Padding.Right)
	renderer.FrameContext.LayoutPosition.Y += GetKitsuneElementVariableValue(element.Padding.Bottom)
	renderer.FrameContext.PageSize.H += GetKitsuneElementVariableValue(element.Padding.Bottom)
}

package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"github.com/veandco/go-sdl2/sdl"
	"modernc.org/mathutil"
)

func PrependKitsuneElementMargin(renderer *KitsuneRenderer, element *node.KitsuneElement) {
	renderer.FrameContext.LayoutPosition.X += GetKitsuneElementVariableValue(element.Margin.Left)

	marginTop := mathutil.MaxInt32(0, GetKitsuneElementVariableValue(element.Margin.Top)-GetKitsuneElementVariableValue(renderer.FrameContext.LayoutForgiveMargin.Bottom))
	renderer.FrameContext.LayoutPosition.Y += marginTop
	renderer.FrameContext.PageSize.H += marginTop
}

func AppendKitsuneElementMargin(renderer *KitsuneRenderer, element *node.KitsuneElement) {
	renderer.FrameContext.LayoutPosition.X += GetKitsuneElementVariableValue(element.Margin.Right)

	marginBottom := GetKitsuneElementVariableValue(element.Margin.Bottom)
	renderer.FrameContext.LayoutPosition.Y += marginBottom
	renderer.FrameContext.PageSize.H += marginBottom
	renderer.FrameContext.LayoutForgiveMargin = element.Margin
}

func GetKitsuneElementMarginRect(renderer *KitsuneRenderer, element *node.KitsuneElement) sdl.Rect {
	return sdl.Rect{
		X: element.CachedRect.X - GetKitsuneElementVariableValue(element.Margin.Left),
		Y: element.CachedRect.Y - GetKitsuneElementVariableValue(element.Margin.Top),
		W: GetKitsuneElementVariableValue(element.Size.Width) + GetKitsuneElementVariableValue(element.Margin.Left) + GetKitsuneElementVariableValue(element.Margin.Right),
		H: GetKitsuneElementVariableValue(element.Size.Height) + GetKitsuneElementVariableValue(element.Margin.Top) + GetKitsuneElementVariableValue(element.Margin.Bottom),
	}
}

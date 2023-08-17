package contextgfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/gfx"
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"modernc.org/mathutil"
)

func CalculateKitsuneElementSize(element *node.KitsuneElement) node.KitsuneElementSize {
	switch node.GetKitsuneElementLayoutDirection(element) {
	case node.KitsuneLayoutDirectionColumn:
		return node.NewKitsuneElementSize(
			CalculateKitsuneElementWidth(element)+gfx.GetKitsuneElementVariableValue(element.OwnSize.Width),
			mathutil.MaxInt32(PickKitsuneElementHeight(element), gfx.GetKitsuneElementVariableValue(element.OwnSize.Height)),
		)
	case node.KitsuneLayoutDirectionRow:
		return node.NewKitsuneElementSize(
			mathutil.MaxInt32(PickKitsuneElementWidth(element), gfx.GetKitsuneElementVariableValue(element.OwnSize.Width)),
			CalculateKitsuneElementHeight(element)+gfx.GetKitsuneElementVariableValue(element.OwnSize.Height),
		)
	}

	return node.KitsuneElementSize{}
}

func PickKitsuneElementWidth(element *node.KitsuneElement) int32 {
	width := int32(0)
	for _, child := range element.Children {
		childWidth := gfx.GetKitsuneElementVariableValue(child.Size.Width)
		if childWidth > width {
			width = childWidth
		}
	}

	return width
}

func CalculateKitsuneElementWidth(element *node.KitsuneElement) int32 {
	width := int32(0)
	for _, child := range element.Children {
		width += gfx.GetKitsuneElementVariableValue(child.Size.Width)
	}

	return width
}

func PickKitsuneElementHeight(element *node.KitsuneElement) int32 {
	height := int32(0)
	for _, child := range element.Children {
		childHeight := gfx.GetKitsuneElementVariableValue(child.Size.Height)
		if childHeight > height {
			height = childHeight
		}
	}

	return height
}

func CalculateKitsuneElementHeight(element *node.KitsuneElement) int32 {
	height := int32(0)
	for _, child := range element.Children {
		height += gfx.GetKitsuneElementVariableValue(child.Size.Height)
	}

	return height
}

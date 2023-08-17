package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"github.com/LamkasDev/kitsune/cmd/common/constants"
	"github.com/LamkasDev/kitsune/cmd/common/sdlutils"
)

func GetKitsuneElementVariableValue(value node.KitsuneElementValueVariable) int32 {
	switch value.Type {
	case node.KitsuneElementValueVariableFixed:
		return value.Value
	}

	return -1
}

func ApplyKitsuneElementSize(renderer *KitsuneRenderer, element *node.KitsuneElement) {
	switch renderer.FrameContext.LayoutDirection {
	case node.KitsuneLayoutDirectionRow:
		renderer.FrameContext.LayoutPosition.Y += GetKitsuneElementVariableValue(element.OwnSize.Height)
		renderer.FrameContext.PageSize.H += GetKitsuneElementVariableValue(element.OwnSize.Height)
	case node.KitsuneLayoutDirectionColumn:
		renderer.FrameContext.LayoutPosition.X += GetKitsuneElementVariableValue(element.OwnSize.Width)
	}
}

func IsMouseOverKitsuneElement(renderer *KitsuneRenderer, element *node.KitsuneElement) bool {
	return sdlutils.IsVectorInsideRect(renderer.WindowContext.MousePosition, GetKitsuneElementRect(renderer, element))
}

func IsKitsuneElementVisible(renderer *KitsuneRenderer, element *node.KitsuneElement) bool {
	return renderer.FrameContext.LayoutPosition.Y >= constants.ApplicationTabHeight-element.Size.Height.Value
}

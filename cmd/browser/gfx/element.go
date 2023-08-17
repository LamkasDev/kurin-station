package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/context"
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"github.com/LamkasDev/kitsune/cmd/common/constants"
	"github.com/LamkasDev/kitsune/cmd/common/sdlutils"
)

func RenderKitsuneRendererElementTree(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender, element *node.KitsuneElement) *error {
	previousX := renderer.FrameContext.LayoutPosition.X
	if err := RenderKitsuneRendererElement(renderer, rcontext, element); err != nil {
		return err
	}

	for _, child := range element.Children {
		if err := RenderKitsuneRendererElementTree(renderer, rcontext, child); err != nil {
			return err
		}
	}

	renderer.FrameContext.LayoutPosition.X = previousX
	if renderer.FrameContext.LayoutDirection == node.KitsuneLayoutDirectionColumn {
		renderer.FrameContext.LayoutPosition.Y += element.CachedRect.H
		renderer.FrameContext.PageSize.H += element.CachedRect.H
		renderer.FrameContext.LayoutDirection = node.KitsuneLayoutDirectionRow
	}

	return nil
}

func RenderKitsuneRendererElement(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender, element *node.KitsuneElement) *error {
	PrependKitsuneElementMargin(renderer, element)
	PrependKitsuneElementPadding(renderer, element)

	if IsKitsuneElementVisible(renderer, element) {
		elementRect := GetKitsuneElementRect(renderer, element)
		element.CachedRect = elementRect
		if IsMouseOverKitsuneElement(renderer, element) && element.OwnSize.Width.Value > 0 {
			marginRect := GetKitsuneElementMarginRect(renderer, element)

			sdlutils.SetDrawColor(renderer.Renderer, constants.MarginColor)
			renderer.Renderer.FillRect(&marginRect)
			sdlutils.SetDrawColor(renderer.Renderer, constants.BoundingBoxFillColor)
			renderer.Renderer.FillRect(&elementRect)
			sdlutils.SetDrawColor(renderer.Renderer, constants.BoundingBoxBorderColor)
			renderer.Renderer.DrawRect(&elementRect)
		}

		switch element.Data.(type) {
		case *node.KitsuneElementTextData:
			if err := renderer.Renderer.Copy(element.Data.(*node.KitsuneElementTextData).Text, nil, &elementRect); err != nil {
				return &err
			}
		case *node.KitsuneElementLinkData:
			if err := renderer.Renderer.Copy(element.Data.(*node.KitsuneElementLinkData).Base.Text, nil, &elementRect); err != nil {
				return &err
			}
		}

		renderer.FrameContext.LayoutDirection = node.GetKitsuneElementLayoutDirection(element)
	}

	ApplyKitsuneElementSize(renderer, element)
	AppendKitsuneElementPadding(renderer, element)
	AppendKitsuneElementMargin(renderer, element)

	return nil
}

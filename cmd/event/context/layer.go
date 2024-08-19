package actions

import (
	"fmt"
	"math"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/context"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLayerContextData struct {
	ContextLayer *gfx.RendererLayer
}

func NewEventLayerContext(contextLayer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerContext,
		Process: ProcessEventLayerContext,
		Data: &EventLayerContextData{
			ContextLayer: contextLayer,
		},
	}
}

func LoadEventLayerContext(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerContext(layer *event.EventLayer) error {
	if gfx.RendererInstance.Context.State != gfx.RendererContextStateNone {
		return nil
	}
	data := layer.Data.(*EventLayerContextData)
	contextData := data.ContextLayer.Data.(*context.RendererLayerContextData)
	if contextData.Position != nil {
		contextData.HoveredItem = -1
		if gfx.RendererInstance.Context.MousePosition.InRect(&sdl.Rect{X: contextData.Position.X, Y: contextData.Position.Y, W: context.RendererLayerContextDataItemWidth, H: int32(len(contextData.Items)) * context.RendererLayerContextDataItemHeight}) {
			hovered := int(math.Floor((float64(gfx.RendererInstance.Context.MousePosition.Y) - float64(contextData.Position.Y)) / context.RendererLayerContextDataItemHeight))
			if hovered >= 0 && hovered < len(contextData.Items) && !contextData.Items[hovered].Disabled {
				contextData.HoveredItem = hovered
			}
		}
	}
	if event.EventManagerInstance.Mouse.PendingRight != nil {
		tile := gameplay.GetTileAt(&gameplay.GameInstance.Map, sdlutils.Vector3{Base: *event.EventManagerInstance.Mouse.PendingRight, Z: gameplay.GameInstance.SelectedCharacter.Position.Z})
		if tile != nil {
			position := gfx.RendererInstance.Context.MousePosition
			contextData.Position = &position
			contextData.Items = []context.RendererLayerContextDataItem{
				{
					Text:     fmt.Sprintf("Tile %d_%d", tile.Position.Base.X, tile.Position.Base.Y),
					Disabled: true,
				},
				{
					Text:    "Inspect",
					OnClick: func() {},
				},
			}
			event.EventManagerInstance.Mouse.PendingRight = nil
		}
	}
	if event.EventManagerInstance.Mouse.PendingLeft != nil {
		if contextData.HoveredItem != -1 {
			contextData.Items[contextData.HoveredItem].OnClick()
			event.EventManagerInstance.Mouse.PendingLeft = nil
		}
		contextData.Position = nil
		contextData.Items = []context.RendererLayerContextDataItem{}
		contextData.HoveredItem = -1
	}

	return nil
}

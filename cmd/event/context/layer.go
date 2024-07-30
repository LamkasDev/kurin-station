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

type KurinEventLayerContextData struct {
	ContextLayer *gfx.KurinRendererLayer
}

func NewKurinEventLayerContext(contextLayer *gfx.KurinRendererLayer) *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerContext,
		Process: ProcessKurinEventLayerContext,
		Data: KurinEventLayerContextData{
			ContextLayer: contextLayer,
		},
	}
}

func LoadKurinEventLayerContext(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerContext(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	if manager.Renderer.RendererContext.State != gfx.KurinRendererContextStateNone {
		return nil
	}
	
	data := layer.Data.(KurinEventLayerContextData).ContextLayer.Data.(context.KurinRendererLayerContextData)
	if data.Position != nil {
		data.HoveredItem = -1
		if manager.Renderer.RendererContext.MousePosition.InRect(&sdl.Rect{X: data.Position.X, Y: data.Position.Y, W: context.KurinRendererLayerContextDataItemWidth, H: int32(len(data.Items))*context.KurinRendererLayerContextDataItemHeight}) {
			hovered := int(math.Floor((float64(manager.Renderer.RendererContext.MousePosition.Y)-float64(data.Position.Y))/context.KurinRendererLayerContextDataItemHeight))
			if hovered >= 0 && hovered < len(data.Items) && !data.Items[hovered].Disabled {
				data.HoveredItem = hovered
			}
		}
	}
	if manager.Mouse.PendingRight != nil {
		tile := gameplay.GetTileAt(&game.Map, sdlutils.Vector3{Base: *manager.Mouse.PendingRight, Z: game.SelectedCharacter.Position.Z})
		if tile != nil {
			position := manager.Renderer.RendererContext.MousePosition
			data.Position = &position
			data.Items = []context.KurinRendererLayerContextDataItem{
				{
					Text: fmt.Sprintf("Tile %d_%d", tile.Position.Base.X, tile.Position.Base.Y),
					Disabled: true,
				},
				{
					Text: "Inspect",
					OnClick: func() {},
				},
			}
			manager.Mouse.PendingRight = nil
		}
	}
	if manager.Mouse.PendingLeft != nil {
		if data.HoveredItem != -1 {
			data.Items[data.HoveredItem].OnClick()
			manager.Mouse.PendingLeft = nil
		}
		data.Position = nil
		data.Items = []context.KurinRendererLayerContextDataItem{}
		data.HoveredItem = -1
	}
	layer.Data.(KurinEventLayerContextData).ContextLayer.Data = data

	return nil
}

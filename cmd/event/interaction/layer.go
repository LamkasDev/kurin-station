package interaction

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/LamkasDev/kurin/cmd/gfx/species"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerInteractionData struct {
	ItemLayer *gfx.RendererLayer
}

func NewKurinEventLayerInteraction(itemLayer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadKurinEventLayerInteraction,
		Process: ProcessKurinEventLayerInteraction,
		Data: &KurinEventLayerInteractionData{
			ItemLayer: itemLayer,
		},
	}
}

func LoadKurinEventLayerInteraction(layer *event.EventLayer) error {
	return nil
}

func ProcessKurinEventLayerInteraction(layer *event.EventLayer) error {
	if gfx.RendererInstance.Context.CameraMode != gfx.KurinRendererCameraModeCharacter || event.EventManagerInstance.Keyboard.InputMode {
		return nil
	}
	if event.EventManagerInstance.Mouse.PendingLeft != nil {
		gameplay.InteractKurinCharacter(gameplay.GameInstance.SelectedCharacter, *event.EventManagerInstance.Mouse.PendingLeft)
		event.EventManagerInstance.Mouse.PendingLeft = nil
	}

	// TODO: add hovered tile and object

	data := layer.Data.(*KurinEventLayerInteractionData)
	itemData := data.ItemLayer.Data.(*item.KurinRendererLayerItemData)
	gameplay.GameInstance.HoveredItem = nil

	for _, currentItem := range gameplay.GameInstance.Map.Items {
		if !gameplay.CanKurinCharacterInteractWithItem(gameplay.GameInstance.SelectedCharacter, currentItem) {
			continue
		}
		graphic := itemData.Items[currentItem.Type]
		hoveredOffset := gfx.GetHoveredOffset(&gfx.RendererInstance.Context, item.GetKurinItemRect(data.ItemLayer, currentItem))
		hoveredOffset = sdlutils.RotatePoint(hoveredOffset, sdl.Point{X: graphic.Textures[0].Surface.W / 2, Y: graphic.Textures[0].Surface.H / 2}, float32(currentItem.Transform.Rotation))
		if gfx.IsHoveredOffsetSolid(graphic.Textures[0], hoveredOffset) {
			gameplay.GameInstance.HoveredItem = currentItem
			event.EventManagerInstance.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
		}
	}

	gameplay.GameInstance.HoveredCharacter = nil
	for _, currentCharacter := range gameplay.GameInstance.Characters {
		if !gameplay.CanKurinCharacterInteractWithCharacter(gameplay.GameInstance.SelectedCharacter, currentCharacter) {
			continue
		}
		hoveredOffset := gfx.GetHoveredOffset(&gfx.RendererInstance.Context, species.GetKurinCharacterRect(currentCharacter))
		if hoveredOffset.InRect(&sdl.Rect{W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y}) {
			gameplay.GameInstance.HoveredCharacter = currentCharacter
		}
	}

	return nil
}

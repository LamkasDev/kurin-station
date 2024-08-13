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
	ItemLayer *gfx.KurinRendererLayer
}

func NewKurinEventLayerInteraction(itemLayer *gfx.KurinRendererLayer) *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerInteraction,
		Process: ProcessKurinEventLayerInteraction,
		Data: KurinEventLayerInteractionData{
			ItemLayer: itemLayer,
		},
	}
}

func LoadKurinEventLayerInteraction(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	return nil
}

func ProcessKurinEventLayerInteraction(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	if manager.Renderer.Context.CameraMode != gfx.KurinRendererCameraModeCharacter || manager.Keyboard.InputMode {
		return nil
	}
	if manager.Mouse.PendingLeft != nil {
		gameplay.InteractKurinCharacter(gameplay.KurinGameInstance.SelectedCharacter, *manager.Mouse.PendingLeft)
		manager.Mouse.PendingLeft = nil
	}

	// TODO: add hovered tile and object

	data := layer.Data.(KurinEventLayerInteractionData)
	gameplay.KurinGameInstance.HoveredItem = nil
	for _, currentItem := range gameplay.KurinGameInstance.Map.Items {
		if !gameplay.CanKurinCharacterInteractWithItem(gameplay.KurinGameInstance.SelectedCharacter, currentItem) {
			continue
		}
		graphic := data.ItemLayer.Data.(item.KurinRendererLayerItemData).Items[currentItem.Type]
		hoveredOffset := gfx.GetHoveredOffset(&manager.Renderer.Context, item.GetKurinItemRect(manager.Renderer, data.ItemLayer, currentItem))
		hoveredOffset = sdlutils.RotatePoint(hoveredOffset, sdl.Point{X: graphic.Textures[0].Surface.W/2, Y: graphic.Textures[0].Surface.H/2}, float32(currentItem.Transform.Rotation))
		if gfx.IsHoveredOffsetSolid(graphic.Textures[0], hoveredOffset) {
			gameplay.KurinGameInstance.HoveredItem = currentItem
			manager.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
		}
	}

	gameplay.KurinGameInstance.HoveredCharacter = nil
	for _, currentCharacter := range gameplay.KurinGameInstance.Characters {
		if !gameplay.CanKurinCharacterInteractWithCharacter(gameplay.KurinGameInstance.SelectedCharacter, currentCharacter) {
			continue
		}
		hoveredOffset := gfx.GetHoveredOffset(&manager.Renderer.Context, species.GetKurinCharacterRect(manager.Renderer, currentCharacter))
		if hoveredOffset.InRect(&sdl.Rect{W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y}) {
			gameplay.KurinGameInstance.HoveredCharacter = currentCharacter
		}
	}

	return nil
}

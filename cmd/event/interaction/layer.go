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

func LoadKurinEventLayerInteraction(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerInteraction(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	if manager.Renderer.WindowContext.CameraMode != gfx.KurinRendererCameraModeCharacter || manager.Keyboard.InputMode {
		return nil
	}
	if manager.Mouse.PendingLeft != nil {
		gameplay.InteractKurinCharacter(game.SelectedCharacter, game, *manager.Mouse.PendingLeft)
		manager.Mouse.PendingLeft = nil
	}

	itemLayer := layer.Data.(KurinEventLayerInteractionData).ItemLayer
	game.HoveredItem = nil
	for _, currentItem := range game.Items {
		if !gameplay.CanKurinCharacterInteractWithItem(game.SelectedCharacter, currentItem) {
			continue
		}
		graphic := itemLayer.Data.(item.KurinRendererLayerItemData).Items[currentItem.Type]
		hoveredOffset := gfx.GetHoveredOffset(&manager.Renderer.WindowContext, item.GetKurinItemRect(manager.Renderer, itemLayer, game, currentItem))
		if hoveredOffset.InRect(&sdl.Rect{W: graphic.Texture.Base.Size.W, H: graphic.Texture.Base.Size.H}) {
			hoveredColor := sdlutils.GetPixelAt(graphic.Texture, hoveredOffset)
			if hoveredColor != nil && sdlutils.IsColorVisible(*hoveredColor) {
				game.HoveredItem = currentItem
			}
		}
	}

	game.HoveredCharacter = nil
	for _, currentCharacter := range game.Characters {
		if !gameplay.CanKurinCharacterInteractWithCharacter(game.SelectedCharacter, currentCharacter) {
			continue
		}
		hoveredOffset := gfx.GetHoveredOffset(&manager.Renderer.WindowContext, species.GetKurinCharacterRect(manager.Renderer, currentCharacter))
		if hoveredOffset.InRect(&sdl.Rect{W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y}) {
			game.HoveredCharacter = currentCharacter
		}
	}

	return nil
}

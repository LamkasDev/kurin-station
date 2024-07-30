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
	Cursors map[sdl.SystemCursor]*sdl.Cursor
}

func NewKurinEventLayerInteraction(itemLayer *gfx.KurinRendererLayer) *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerInteraction,
		Process: ProcessKurinEventLayerInteraction,
		Data: KurinEventLayerInteractionData{
			ItemLayer: itemLayer,
			Cursors: map[sdl.SystemCursor]*sdl.Cursor{
				sdl.SYSTEM_CURSOR_ARROW: sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_ARROW),
				sdl.SYSTEM_CURSOR_HAND: sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_HAND),
			},
		},
	}
}

func LoadKurinEventLayerInteraction(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerInteraction(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	if manager.Renderer.RendererContext.CameraMode != gfx.KurinRendererCameraModeCharacter || manager.Keyboard.InputMode {
		return nil
	}
	if manager.Mouse.PendingLeft != nil {
		gameplay.InteractKurinCharacter(game.SelectedCharacter, game, *manager.Mouse.PendingLeft)
		manager.Mouse.PendingLeft = nil
	}

	// TODO: add hovered tile and object

	data := layer.Data.(KurinEventLayerInteractionData)
	cursor := sdl.SYSTEM_CURSOR_ARROW

	game.HoveredItem = nil
	for _, currentItem := range game.Map.Items {
		if !gameplay.CanKurinCharacterInteractWithItem(game.SelectedCharacter, currentItem) {
			continue
		}
		graphic := data.ItemLayer.Data.(item.KurinRendererLayerItemData).Items[currentItem.Type]
		hoveredOffset := gfx.GetHoveredOffset(&manager.Renderer.RendererContext, item.GetKurinItemRect(manager.Renderer, data.ItemLayer, game, currentItem))
		hoveredOffset = sdlutils.RotatePoint(hoveredOffset, sdl.Point{X: graphic.Textures[0].Surface.W/2, Y: graphic.Textures[0].Surface.H/2}, float32(currentItem.Transform.Rotation))
		if gfx.IsHoveredOffsetSolid(graphic.Textures[0], hoveredOffset) {
			game.HoveredItem = currentItem
			cursor = sdl.SYSTEM_CURSOR_HAND
		}
	}

	game.HoveredCharacter = nil
	for _, currentCharacter := range game.Characters {
		if !gameplay.CanKurinCharacterInteractWithCharacter(game.SelectedCharacter, currentCharacter) {
			continue
		}
		hoveredOffset := gfx.GetHoveredOffset(&manager.Renderer.RendererContext, species.GetKurinCharacterRect(manager.Renderer, currentCharacter))
		if hoveredOffset.InRect(&sdl.Rect{W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y}) {
			game.HoveredCharacter = currentCharacter
		}
	}

	sdl.SetCursor(data.Cursors[cursor])
	return nil
}

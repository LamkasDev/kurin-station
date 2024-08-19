package interaction

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/LamkasDev/kurin/cmd/gfx/species"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLayerInteractionData struct {
	ItemLayer *gfx.RendererLayer
}

func NewEventLayerInteraction(itemLayer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerInteraction,
		Process: ProcessEventLayerInteraction,
		Data: &EventLayerInteractionData{
			ItemLayer: itemLayer,
		},
	}
}

func LoadEventLayerInteraction(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerInteraction(layer *event.EventLayer) error {
	if gfx.RendererInstance.Context.CameraMode != gfx.RendererCameraModeCharacter || event.EventManagerInstance.Keyboard.InputMode {
		return nil
	}
	data := layer.Data.(*EventLayerInteractionData)
	itemData := data.ItemLayer.Data.(*item.RendererLayerItemData)

	gameplay.GameInstance.HoveredTile = nil
	gameplay.GameInstance.HoveredObject = nil
	mousePosition := render.ScreenToWorldPosition(gfx.RendererInstance.Context.MousePosition)
	tile := gameplay.GetTileAt(&gameplay.GameInstance.Map, sdlutils.Vector3{Base: mousePosition, Z: 0})
	if tile != nil {
		gameplay.GameInstance.HoveredTile = tile
		gameplay.GameInstance.HoveredObject = gameplay.GetObjectAtTile(tile)
	}

	gameplay.GameInstance.HoveredItem = nil
	for _, currentItem := range gameplay.GameInstance.Map.Items {
		if !gameplay.CanCharacterInteractWithItem(gameplay.GameInstance.SelectedCharacter, currentItem) {
			continue
		}
		graphic := itemData.Items[currentItem.Type]
		hoveredOffset := gfx.GetHoveredOffset(&gfx.RendererInstance.Context, item.GetItemRect(data.ItemLayer, currentItem))
		hoveredOffset = sdlutils.RotatePoint(hoveredOffset, sdl.Point{X: graphic.Textures[0].Surface.W / 2, Y: graphic.Textures[0].Surface.H / 2}, float32(currentItem.Transform.Rotation))
		if gfx.IsHoveredOffsetSolid(graphic.Textures[0], hoveredOffset) {
			gameplay.GameInstance.HoveredItem = currentItem
			event.EventManagerInstance.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
		}
	}

	gameplay.GameInstance.HoveredCharacter = nil
	for _, currentCharacter := range gameplay.GameInstance.Characters {
		if !gameplay.CanCharacterInteractWithCharacter(gameplay.GameInstance.SelectedCharacter, currentCharacter) {
			continue
		}
		hoveredOffset := gfx.GetHoveredOffset(&gfx.RendererInstance.Context, species.GetCharacterRect(currentCharacter))
		if hoveredOffset.InRect(&sdl.Rect{W: gameplay.TileSize.X, H: gameplay.TileSize.Y}) {
			gameplay.GameInstance.HoveredCharacter = currentCharacter
		}
	}

	if event.EventManagerInstance.Mouse.PendingLeft != nil {
		gameplay.InteractCharacter(gameplay.GameInstance.SelectedCharacter, *event.EventManagerInstance.Mouse.PendingLeft)
		event.EventManagerInstance.Mouse.PendingLeft = nil
	}

	return nil
}

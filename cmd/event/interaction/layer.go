package interaction

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/LamkasDev/kurin/cmd/gfx/mob"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
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
	tile := gameplay.GetTileAt(gameplay.GameInstance.Map, sdlutils.Vector3{Base: mousePosition, Z: gameplay.GameInstance.SelectedZ})
	if tile != nil {
		gameplay.GameInstance.HoveredTile = tile
		gameplay.GameInstance.HoveredObject = gameplay.GetObjectAtTile(tile)
	}

	gameplay.GameInstance.HoveredItem = nil
	for _, currentItem := range gameplay.GameInstance.Map.Items {
		if !currentItem.Template.CanPickup || !gameplay.CanMobInteractWithItem(gameplay.GameInstance.SelectedCharacter, currentItem) {
			continue
		}
		graphic := itemData.Items[currentItem.Type]
		hoveredOffset := gfx.GetHoveredOffset(&gfx.RendererInstance.Context, item.GetItemRect(data.ItemLayer, currentItem, graphic.Textures[0].Base))
		hoveredOffset = sdlutils.RotatePoint(hoveredOffset, sdl.Point{X: graphic.Textures[0].Surface.W / 2, Y: graphic.Textures[0].Surface.H / 2}, float32(currentItem.Transform.Rotation))
		if gfx.IsHoveredOffsetSolid(graphic.Textures[0], hoveredOffset) {
			gameplay.GameInstance.HoveredItem = currentItem
			event.EventManagerInstance.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
		}
	}

	gameplay.GameInstance.HoveredMob = nil
	for _, currentMob := range gameplay.GameInstance.Map.Mobs {
		if !gameplay.CanMobInteractWithMob(gameplay.GameInstance.SelectedCharacter, currentMob) {
			continue
		}
		hoveredOffset := gfx.GetHoveredOffset(&gfx.RendererInstance.Context, mob.GetMobRect(currentMob))
		if hoveredOffset.InRect(&sdl.Rect{W: gameplay.TileSize.X, H: gameplay.TileSize.Y}) {
			gameplay.GameInstance.HoveredMob = currentMob
		}
	}

	if event.EventManagerInstance.Mouse.PendingLeft != nil {
		gameplay.InteractCharacter(gameplay.GameInstance.SelectedCharacter, *event.EventManagerInstance.Mouse.PendingLeft)
		event.EventManagerInstance.Mouse.PendingLeft = nil
	}

	return nil
}

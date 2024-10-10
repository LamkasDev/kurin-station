package keybinds

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/hud"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLayerHUDData struct {
	HudLayer  *gfx.RendererLayer
	ItemLayer *gfx.RendererLayer
	Cursors   map[sdl.SystemCursor]*sdl.Cursor
}

func NewEventLayerHUD(hudLayer *gfx.RendererLayer, itemLayer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerHUD,
		Process: ProcessEventLayerHUD,
		Data: &EventLayerHUDData{
			HudLayer:  hudLayer,
			ItemLayer: itemLayer,
			Cursors: map[sdl.SystemCursor]*sdl.Cursor{
				sdl.SYSTEM_CURSOR_ARROW: sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_ARROW),
				sdl.SYSTEM_CURSOR_HAND:  sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_HAND),
			},
		},
	}
}

func LoadEventLayerHUD(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerHUD(layer *event.EventLayer) error {
	if gameplay.GameInstance.SelectedCharacter == nil {
		return nil
	}

	data := layer.Data.(*EventLayerHUDData)
	itemData := data.ItemLayer.Data.(*item.RendererLayerItemData)
	hudData := data.HudLayer.Data.(*hud.RendererLayerHUDData)
	hudData.HoveredItem = nil

	lhand := gameplay.GameInstance.SelectedCharacter.Data.(*gameplay.MobCharacterData).Inventory.Hands[gameplay.HandLeft]
	lhandRect := gfx.GetUIRect(hudData.Icons["hand_l"].Texture, hud.HUDElementHandLeft.GetPosition(gfx.RendererInstance.Context.WindowSize), hud.HUDElementHandLeft.Scale, hud.HUDElementHandLeft.Anchor)
	if lhand != nil {
		hoveredOffset := sdlutils.DividePoints(sdlutils.SubtractPoints(gfx.RendererInstance.Context.MousePosition, sdl.Point{X: lhandRect.X, Y: lhandRect.Y}), sdlutils.FPointToPoint(sdlutils.MultiplyFPoints(hud.HUDElementHandLeft.Scale, gfx.RendererInstance.Context.WindowScale)))
		if gfx.IsHoveredOffsetSolid(itemData.Items[lhand.Type].Textures[0], hoveredOffset) {
			event.EventManagerInstance.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
			hudData.HoveredItem = lhand
			if event.EventManagerInstance.Mouse.PendingLeft != nil {
				lhand.Template.OnHandInteraction(lhand)
				event.EventManagerInstance.Mouse.PendingLeft = nil
			}
		}
	}
	rhand := gameplay.GameInstance.SelectedCharacter.Data.(*gameplay.MobCharacterData).Inventory.Hands[gameplay.HandRight]
	rhandRect := gfx.GetUIRect(hudData.Icons["hand_r"].Texture, hud.HUDElementHandRight.GetPosition(gfx.RendererInstance.Context.WindowSize), hud.HUDElementHandRight.Scale, hud.HUDElementHandRight.Anchor)
	if rhand != nil {
		hoveredOffset := sdlutils.DividePoints(sdlutils.SubtractPoints(gfx.RendererInstance.Context.MousePosition, sdl.Point{X: rhandRect.X, Y: rhandRect.Y}), sdlutils.FPointToPoint(sdlutils.MultiplyFPoints(hud.HUDElementHandRight.Scale, gfx.RendererInstance.Context.WindowScale)))
		if gfx.IsHoveredOffsetSolid(itemData.Items[rhand.Type].Textures[0], hoveredOffset) {
			event.EventManagerInstance.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
			hudData.HoveredItem = rhand
			if event.EventManagerInstance.Mouse.PendingLeft != nil {
				rhand.Template.OnHandInteraction(rhand)
				event.EventManagerInstance.Mouse.PendingLeft = nil
			}
		}
	}

	for _, element := range hud.HUDElements {
		rect := gfx.GetUIRect(hudData.Icons[element.Path].Texture, element.GetPosition(gfx.RendererInstance.Context.WindowSize), element.Scale, element.Anchor)
		if gfx.RendererInstance.Context.MousePosition.InRect(&rect) {
			element.Hovered = true
			if event.EventManagerInstance.Mouse.PendingLeft != nil {
				element.Click()
				event.EventManagerInstance.Mouse.PendingLeft = nil
			}
		} else {
			element.Hovered = false
		}
	}

	return nil
}

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

type KurinEventLayerHUDData struct {
	HudLayer *gfx.KurinRendererLayer
	ItemLayer *gfx.KurinRendererLayer
	Cursors map[sdl.SystemCursor]*sdl.Cursor
}

func NewKurinEventLayerHUD(hudLayer *gfx.KurinRendererLayer, itemLayer *gfx.KurinRendererLayer) *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerHUD,
		Process: ProcessKurinEventLayerHUD,
		Data:    KurinEventLayerHUDData{
			HudLayer: hudLayer,
			ItemLayer: itemLayer,
			Cursors: map[sdl.SystemCursor]*sdl.Cursor{
				sdl.SYSTEM_CURSOR_ARROW: sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_ARROW),
				sdl.SYSTEM_CURSOR_HAND: sdl.CreateSystemCursor(sdl.SYSTEM_CURSOR_HAND),
			},
		},
	}
}

func LoadKurinEventLayerHUD(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerHUD(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	data := layer.Data.(KurinEventLayerHUDData)
	itemData := data.ItemLayer.Data.(item.KurinRendererLayerItemData)

	hudData := data.HudLayer.Data.(hud.KurinRendererLayerHUDData)
	hudData.HoveredItem = nil
	lhand := game.SelectedCharacter.Inventory.Hands[gameplay.KurinHandLeft]
	if lhand != nil {
		hoveredOffset := sdlutils.DividePoints(gfx.GetHoveredOffsetUnscaled(&manager.Renderer.Context, hud.KurinHUDElementHandLeft.GetPosition(manager.Renderer.Context.WindowSize)), sdl.Point{X: 2, Y: 2})
		if gfx.IsHoveredOffsetSolid(itemData.Items[lhand.Type].Textures[0], hoveredOffset) {
			manager.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
			hudData.HoveredItem = lhand
			if manager.Mouse.PendingLeft != nil {
				lhand.Interact(lhand, game);
				manager.Mouse.PendingLeft = nil;
			}
		}
	}
	rhand := game.SelectedCharacter.Inventory.Hands[gameplay.KurinHandRight]
	if rhand != nil {
		hoveredOffset := sdlutils.DividePoints(gfx.GetHoveredOffsetUnscaled(&manager.Renderer.Context, hud.KurinHUDElementHandRight.GetPosition(manager.Renderer.Context.WindowSize)), sdl.Point{X: 2, Y: 2})
		if gfx.IsHoveredOffsetSolid(itemData.Items[rhand.Type].Textures[0], hoveredOffset) {
			manager.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
			hudData.HoveredItem = rhand
			if manager.Mouse.PendingLeft != nil {
				rhand.Interact(rhand, game);
				manager.Mouse.PendingLeft = nil;
			}
		}
	}
	data.HudLayer.Data = hudData

	for _, element := range hud.KurinHUDElements {
		pos := element.GetPosition(manager.Renderer.Context.WindowSize)
		if manager.Renderer.Context.MousePosition.InRect(&sdl.Rect{X: pos.X, Y: pos.Y, W: 64, H: 64}) {
			element.Hovered = true
			if manager.Mouse.PendingLeft != nil {
				element.Click(game)
				manager.Mouse.PendingLeft = nil
			}
		} else {
			element.Hovered = false
		}
	}

	return nil
}

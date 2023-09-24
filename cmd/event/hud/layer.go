package keybinds

import (
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx/hud"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerHUDData struct {
}

func NewKurinEventLayerHUD() *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerHUD,
		Process: ProcessKurinEventLayerHUD,
		Data:    KurinEventLayerHUDData{},
	}
}

func LoadKurinEventLayerHUD(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerHUD(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	for _, element := range hud.KurinHUDElements {
		pos := element.GetPosition(manager.Renderer.WindowContext.WindowSize)
		if manager.Renderer.WindowContext.MousePosition.InRect(&sdl.Rect{X: pos.X, Y: pos.Y, W: 64, H: 64}) {
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

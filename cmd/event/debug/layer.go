package keybinds

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/debug"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerDebugData struct {
	Layer *gfx.KurinRendererLayer
}

func NewKurinEventLayerDebug(layer *gfx.KurinRendererLayer) *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerDebug,
		Process: ProcessKurinEventLayerDebug,
		Data: KurinEventLayerDebugData{
			Layer: layer,
		},
	}
}

func LoadKurinEventLayerDebug(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerDebug(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	if manager.Keyboard.Pending != nil {
		switch *manager.Keyboard.Pending {
		case sdl.K_p:
			actionsData := layer.Data.(KurinEventLayerDebugData).Layer.Data.(debug.KurinRendererLayerDebugData)
			actionsData.Path = gameplay.FindKurinPath(&game.Map.Pathfinding, game.SelectedCharacter.Position, sdlutils.Vector3{Base: sdl.Point{X: 0, Y: 0}, Z: 0})
			layer.Data.(KurinEventLayerDebugData).Layer.Data = actionsData
		default:
			return nil
		}
		manager.Keyboard.Pending = nil
	}

	return nil
}

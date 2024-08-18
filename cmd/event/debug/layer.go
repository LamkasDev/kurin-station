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
	Layer *gfx.RendererLayer
}

func NewKurinEventLayerDebug(layer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadKurinEventLayerDebug,
		Process: ProcessKurinEventLayerDebug,
		Data: &KurinEventLayerDebugData{
			Layer: layer,
		},
	}
}

func LoadKurinEventLayerDebug(layer *event.EventLayer) error {
	return nil
}

func ProcessKurinEventLayerDebug(layer *event.EventLayer) error {
	if event.EventManagerInstance.Keyboard.Pending == nil {
		return nil
	}
	data := layer.Data.(*KurinEventLayerDebugData).Layer.Data.(*debug.KurinRendererLayerDebugData)
	switch *event.EventManagerInstance.Keyboard.Pending {
	case sdl.K_p:
		data.Path = gameplay.FindKurinPath(&gameplay.GameInstance.Map.Pathfinding, gameplay.GameInstance.SelectedCharacter.Position, sdlutils.Vector3{Base: sdl.Point{X: 0, Y: 0}, Z: 0})
		event.EventManagerInstance.Keyboard.Pending = nil
	case sdl.K_o:
		if len(gameplay.GameInstance.Narrator.Objectives) > 0 {
			gameplay.CompleteKurinNarratorObjective(gameplay.GameInstance.Narrator)
		}
		event.EventManagerInstance.Keyboard.Pending = nil
	default:
		return nil
	}

	return nil
}

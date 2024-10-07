package keybinds

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/debug"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLayerDebugData struct {
	Layer *gfx.RendererLayer
}

func NewEventLayerDebug(layer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerDebug,
		Process: ProcessEventLayerDebug,
		Data: &EventLayerDebugData{
			Layer: layer,
		},
	}
}

func LoadEventLayerDebug(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerDebug(layer *event.EventLayer) error {
	if event.EventManagerInstance.Keyboard.Pending == nil {
		return nil
	}
	data := layer.Data.(*EventLayerDebugData).Layer.Data.(*debug.RendererLayerDebugData)
	switch *event.EventManagerInstance.Keyboard.Pending {
	case sdl.K_p:
		data.Path = gameplay.FindPath(&gameplay.GameInstance.Map.Pathfinding, gameplay.GameInstance.SelectedCharacter.Position, sdlutils.Vector3{Base: sdl.Point{X: 0, Y: 0}, Z: gameplay.GameInstance.Map.BaseZ})
		event.EventManagerInstance.Keyboard.Pending = nil
	case sdl.K_o:
		if len(gameplay.GameInstance.Narrator.Objectives) > 0 {
			gameplay.CompleteNarratorObjective(gameplay.GameInstance.Narrator)
		}
		event.EventManagerInstance.Keyboard.Pending = nil
	case sdl.K_F1:
		gameplay.GameInstance.Godmode = !gameplay.GameInstance.Godmode
	default:
		return nil
	}

	return nil
}

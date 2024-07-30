package actions

import (
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/actions"
	"github.com/LamkasDev/kurin/cmd/gfx/tool"
	"github.com/adam-lavrik/go-imath/ix"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type KurinEventLayerActionsData struct {
	ActionsLayer *gfx.KurinRendererLayer
	ToolLayer    *gfx.KurinRendererLayer
}

func NewKurinEventLayerActions(actionsLayer *gfx.KurinRendererLayer, toolLayer *gfx.KurinRendererLayer) *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerActions,
		Process: ProcessKurinEventLayerActions,
		Data: KurinEventLayerActionsData{
			ActionsLayer: actionsLayer,
			ToolLayer:    toolLayer,
		},
	}
}

func LoadKurinEventLayerActions(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerActions(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	if manager.Keyboard.Pending == nil {
		return nil
	}
	if manager.Renderer.RendererContext.State == gfx.KurinRendererContextStateActions {
		ProcessKurinEventLayerActionsInput(manager, layer, game)
		manager.Keyboard.Pending = nil
		return nil
	}
	switch *manager.Keyboard.Pending {
	case sdl.K_t:
		StartKurinEventLayerActionsInput(manager, layer, actions.KurinActionModeSay)
		manager.Keyboard.Pending = nil
		break
	case sdl.K_b:
		StartKurinEventLayerActionsInput(manager, layer, actions.KurinActionModeBuild)
		manager.Keyboard.Pending = nil
		break
	}

	return nil
}

func StartKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, mode actions.KurinActionMode) {
	data := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData)
	data.Mode = mode
	layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data = data
	
	manager.Keyboard.InputMode = true
	manager.Renderer.RendererContext.State = gfx.KurinRendererContextStateActions
}

func ProcessKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) {
	data := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData)
	input := data.Input
	switch *manager.Keyboard.Pending {
	case sdl.K_ESCAPE:
		EndKurinEventLayerActionsInput(manager, layer, game)
		return
	case sdl.K_RETURN:
		ExecuteKurinEventLayerActionsInput(manager, layer, game)
		return
	case sdl.K_UP:
		data.Index = ix.Max(ix.Min(data.Index-1, len(actions.GetMenuStructureGraphics(&data))-1), 0)
		layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data = data
		return
	case sdl.K_DOWN:
		data.Index = ix.Max(ix.Min(data.Index+1, len(actions.GetMenuStructureGraphics(&data))-1), 0)
		layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data = data
		return
	case sdl.K_BACKSPACE:
		if len(input) > 0 {
			input = input[:len(input)-1]
		}
	default:
		input += manager.Keyboard.Input
	}

	currentStructures := actions.GetMenuStructureGraphics(&data)
	if data.Index < len(currentStructures) {
		currentStructure := currentStructures[data.Index]
		data.Input = input
		newStructures := actions.GetMenuStructureGraphics(&data)
		data.Index = ix.Max(slices.Index(newStructures, currentStructure), 0)
	} else {
		data.Input = input
		data.Index = 0
	}
	layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data = data
}

func ExecuteKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) {
	data := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData)
	input := data.Input
	EndKurinEventLayerActionsInput(manager, layer, game)

	switch data.Mode {
	case actions.KurinActionModeBuild:
		manager.Renderer.RendererContext.State = gfx.KurinRendererContextStateTool
		toolData := layer.Data.(KurinEventLayerActionsData).ToolLayer.Data.(tool.KurinRendererLayerToolData)
		toolData.Mode = tool.KurinToolModeBuild
		toolData.Prefab = gameplay.NewKurinObject(actions.GetMenuStructureGraphics(&data)[data.Index].Template.Id)
		layer.Data.(KurinEventLayerActionsData).ToolLayer.Data = toolData
	case actions.KurinActionModeSay:
		if len(input) == 0 {
			return
		}
		gameplay.CreateKurinRunechatMessage(&game.RunechatController, gameplay.NewKurinRunechatCharacter(game.SelectedCharacter, input))
	}
}

func EndKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) {
	data := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData)
	data.Input = ""
	data.Index = 0
	layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data = data

	manager.Renderer.RendererContext.State = gfx.KurinRendererContextStateNone
	manager.Keyboard.InputMode = false
}

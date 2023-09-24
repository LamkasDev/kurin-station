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
	if manager.Renderer.WindowContext.State == gfx.KurinRendererContextStateActions {
		ProcessKurinEventLayerActionsInput(manager, layer, game)
		manager.Keyboard.Pending = nil
	} else if *manager.Keyboard.Pending == sdl.K_t {
		StartKurinEventLayerActionsInput(manager, layer, actions.KurinActionModeSay)
		manager.Keyboard.Pending = nil
	} else if *manager.Keyboard.Pending == sdl.K_b {
		StartKurinEventLayerActionsInput(manager, layer, actions.KurinActionModeBuild)
		manager.Keyboard.Pending = nil
	}

	return nil
}

func StartKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, mode actions.KurinActionMode) {
	manager.Keyboard.InputMode = true
	manager.Renderer.WindowContext.State = gfx.KurinRendererContextStateActions
	SetKurinEventLayerActionsMode(manager, layer, mode)
}

func ProcessKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) {
	input := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData).Input
	switch *manager.Keyboard.Pending {
	case sdl.K_ESCAPE:
		EndKurinEventLayerActionsInput(manager, layer, game)
		return
	case sdl.K_RETURN:
		ExecuteKurinEventLayerActionsInput(manager, layer, game)
		return
	case sdl.K_UP:
		MoveKurinEventLayerActionsIndex(manager, layer, -1)
		return
	case sdl.K_DOWN:
		MoveKurinEventLayerActionsIndex(manager, layer, 1)
		return
	case sdl.K_BACKSPACE:
		if len(input) > 0 {
			input = input[:len(input)-1]
		}
	default:
		input += manager.Keyboard.Input
	}
	SetKurinEventLayerActionsInput(manager, layer, input)
}

func SetKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, input string) {
	actionsData := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData)
	currentStructures := actions.GetMenuStructureGraphics(&actionsData)
	if actionsData.Index < len(currentStructures) {
		currentStructure := currentStructures[actionsData.Index]
		actionsData.Input = input
		newStructures := actions.GetMenuStructureGraphics(&actionsData)
		actionsData.Index = ix.Max(slices.Index(newStructures, currentStructure), 0)
	} else {
		actionsData.Input = input
		actionsData.Index = 0
	}
	layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data = actionsData
}

func SetKurinEventLayerActionsMode(manager *event.KurinEventManager, layer *event.KurinEventLayer, mode actions.KurinActionMode) {
	actionsData := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData)
	actionsData.Mode = mode
	layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data = actionsData
}

func MoveKurinEventLayerActionsIndex(manager *event.KurinEventManager, layer *event.KurinEventLayer, offset int) {
	actionsData := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData)
	actionsData.Index = ix.Max(ix.Min(actionsData.Index+offset, len(actions.GetMenuStructureGraphics(&actionsData))-1), 0)
	layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data = actionsData
}

func ExecuteKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) {
	actionsData := layer.Data.(KurinEventLayerActionsData).ActionsLayer.Data.(actions.KurinRendererLayerActionsData)
	input := actionsData.Input
	EndKurinEventLayerActionsInput(manager, layer, game)

	switch actionsData.Mode {
	case actions.KurinActionModeBuild:
		manager.Renderer.WindowContext.State = gfx.KurinRendererContextStateTool
		toolData := layer.Data.(KurinEventLayerActionsData).ToolLayer.Data.(tool.KurinRendererLayerToolData)
		toolData.Mode = tool.KurinToolModeBuild
		toolData.Prefab = gameplay.NewKurinObject(actions.GetMenuStructureGraphics(&actionsData)[actionsData.Index].Template.Id)
		layer.Data.(KurinEventLayerActionsData).ToolLayer.Data = toolData
	case actions.KurinActionModeSay:
		if len(input) == 0 {
			return
		}
		gameplay.CreateKurinRunechatMessage(&game.RunechatController, gameplay.NewKurinRunechatCharacter(game.SelectedCharacter, input))
	}
}

func EndKurinEventLayerActionsInput(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) {
	manager.Renderer.WindowContext.State = gfx.KurinRendererContextStateNone
	manager.Keyboard.InputMode = false
	SetKurinEventLayerActionsInput(manager, layer, "")
}

package actions

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/actions"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/LamkasDev/kurin/cmd/gfx/tool"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"github.com/adam-lavrik/go-imath/ix"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type EventLayerActionsData struct {
	ActionsLayer *gfx.RendererLayer
	ToolLayer    *gfx.RendererLayer
}

func NewEventLayerActions(actionsLayer *gfx.RendererLayer, toolLayer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerActions,
		Process: ProcessEventLayerActions,
		Data: &EventLayerActionsData{
			ActionsLayer: actionsLayer,
			ToolLayer:    toolLayer,
		},
	}
}

func LoadEventLayerActions(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerActions(layer *event.EventLayer) error {
	data := layer.Data.(*EventLayerActionsData)
	toolData := data.ToolLayer.Data.(*tool.RendererLayerToolData)
	if event.EventManagerInstance.Keyboard.Pending == nil {
		return nil
	}
	if gfx.RendererInstance.Context.State == gfx.RendererContextStateActions {
		ProcessEventLayerActionsInput(layer)
		event.EventManagerInstance.Keyboard.Pending = nil
		return nil
	}
	if gfx.RendererInstance.Context.State == gfx.RendererContextStateNone && !event.EventManagerInstance.Keyboard.InputMode {
		switch *event.EventManagerInstance.Keyboard.Pending {
		case sdl.K_t:
			StartEventLayerActionsInput(layer, actions.ActionModeSay)
			event.EventManagerInstance.Keyboard.Pending = nil
		case sdl.K_b:
			StartEventLayerActionsInput(layer, actions.ActionModeBuild)
			event.EventManagerInstance.Keyboard.Pending = nil
		case sdl.K_n:
			gfx.RendererInstance.Context.State = gfx.RendererContextStateTool
			toolData.Mode = tool.ToolModeDestroy
			event.EventManagerInstance.Keyboard.Pending = nil
		}
	}

	return nil
}

func StartEventLayerActionsInput(layer *event.EventLayer, mode actions.ActionMode) {
	actionsData := layer.Data.(*EventLayerActionsData).ActionsLayer.Data.(*actions.RendererLayerActionsData)
	actionsData.Input = ""
	actionsData.Index = 0
	actionsData.Mode = mode
	event.EventManagerInstance.Keyboard.InputMode = true
	gfx.RendererInstance.Context.State = gfx.RendererContextStateActions
}

func ProcessEventLayerActionsInput(layer *event.EventLayer) {
	data := layer.Data.(*EventLayerActionsData)
	actionsData := data.ActionsLayer.Data.(*actions.RendererLayerActionsData)
	input := actionsData.Input
	switch *event.EventManagerInstance.Keyboard.Pending {
	case sdl.K_ESCAPE:
		EndEventLayerActionsInput(layer)
		return
	case sdl.K_RETURN:
		ExecuteEventLayerActionsInput(layer)
		return
	case sdl.K_UP:
		actionsData.Index = ix.Max(ix.Min(actionsData.Index-1, len(actions.GetMenuGraphics(actionsData))-1), 0)
		return
	case sdl.K_DOWN:
		actionsData.Index = ix.Max(ix.Min(actionsData.Index+1, len(actions.GetMenuGraphics(actionsData))-1), 0)
		return
	case sdl.K_BACKSPACE:
		if len(input) > 0 {
			input = input[:len(input)-1]
		}
	default:
		input += event.EventManagerInstance.Keyboard.Input
	}

	currentStructures := actions.GetMenuGraphics(actionsData)
	if actionsData.Index < len(currentStructures) {
		currentStructure := currentStructures[actionsData.Index]
		actionsData.Input = input
		newStructures := actions.GetMenuGraphics(actionsData)
		actionsData.Index = ix.Max(slices.Index(newStructures, currentStructure), 0)
	} else {
		actionsData.Input = input
		actionsData.Index = 0
	}
}

func ExecuteEventLayerActionsInput(layer *event.EventLayer) {
	data := layer.Data.(*EventLayerActionsData)
	actionsData := data.ActionsLayer.Data.(*actions.RendererLayerActionsData)
	toolData := data.ToolLayer.Data.(*tool.RendererLayerToolData)
	EndEventLayerActionsInput(layer)
	switch actionsData.Mode {
	case actions.ActionModeBuild:
		gfx.RendererInstance.Context.State = gfx.RendererContextStateTool
		graphic := actions.GetMenuGraphics(actionsData)[actionsData.Index]
		switch realGraphic := graphic.(type) {
		case *structure.StructureGraphic:
			toolData.Mode = tool.ToolModeBuild
			toolData.Prefab = gameplay.NewObject(&gameplay.Tile{}, realGraphic.Template.Id)
		case *turf.TurfGraphic:
			toolData.Mode = tool.ToolModeBuild
			toolData.Prefab = gameplay.NewTile(realGraphic.Template.Id, sdlutils.Vector3{})
		}
	case actions.ActionModeSay:
		if len(actionsData.Input) == 0 {
			return
		}
		gameplay.CreateRunechatMessage(&gameplay.GameInstance.RunechatController, gameplay.NewRunechatMob(gameplay.GameInstance.SelectedCharacter, actionsData.Input))
	}
}

func EndEventLayerActionsInput(layer *event.EventLayer) {
	gfx.RendererInstance.Context.State = gfx.RendererContextStateNone
	event.EventManagerInstance.Keyboard.InputMode = false
}

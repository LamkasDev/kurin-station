package tool

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/LamkasDev/kurin/cmd/gfx/tool"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerToolData struct {
	Layer *gfx.KurinRendererLayer
}

func NewKurinEventLayerTool(layer *gfx.KurinRendererLayer) *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerTool,
		Process: ProcessKurinEventLayerTool,
		Data: KurinEventLayerToolData{
			Layer: layer,
		},
	}
}

func LoadKurinEventLayerTool(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	return nil
}

func ProcessKurinEventLayerTool(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	if manager.Renderer.Context.State == gfx.KurinRendererContextStateTool {
		ProcessKurinEventLayerToolInput(manager, layer)
	}

	return nil
}

func ProcessKurinEventLayerToolInput(manager *event.KurinEventManager, layer *event.KurinEventLayer) {
	data := layer.Data.(KurinEventLayerToolData)
	toolData := data.Layer.Data.(tool.KurinRendererLayerToolData)
	if manager.Keyboard.Pending != nil {
		switch *manager.Keyboard.Pending {
		case sdl.K_ESCAPE:
			manager.Renderer.Context.State = gfx.KurinRendererContextStateNone
			manager.Keyboard.Pending = nil
			return
		}
	}
	if manager.Mouse.PendingRight != nil {
		manager.Renderer.Context.State = gfx.KurinRendererContextStateNone
		manager.Mouse.PendingRight = nil
	}
	wrect := render.ScreenToWorldRect(manager.Renderer, sdl.Rect{X: manager.Renderer.Context.MousePosition.X, Y: manager.Renderer.Context.MousePosition.Y, W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y})
	tile := gameplay.GetTileAt(&gameplay.KurinGameInstance.Map, sdlutils.Vector3{Base: sdl.Point{X: wrect.X, Y: wrect.Y}, Z: 0})
	toolData.Prefab.Tile = tile
	if toolData.Prefab.Tile == nil {
		return
	}
	if manager.Mouse.PendingLeft != nil {
		prefabCopy := *toolData.Prefab
		gameplay.PushKurinJobToController(&gameplay.KurinGameInstance.JobController, gameplay.NewKurinJobDriverBuild(gameplay.KurinJobDriverBuildData{
			Prefab: &prefabCopy,
		}))
	}
}

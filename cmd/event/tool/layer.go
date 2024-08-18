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
	Layer *gfx.RendererLayer
}

func NewKurinEventLayerTool(layer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadKurinEventLayerTool,
		Process: ProcessKurinEventLayerTool,
		Data: &KurinEventLayerToolData{
			Layer: layer,
		},
	}
}

func LoadKurinEventLayerTool(layer *event.EventLayer) error {
	return nil
}

func ProcessKurinEventLayerTool(layer *event.EventLayer) error {
	if gfx.RendererInstance.Context.State != gfx.KurinRendererContextStateTool {
		return nil
	}

	ProcessKurinEventLayerToolInput(layer)
	return nil
}

func ProcessKurinEventLayerToolInput(layer *event.EventLayer) {
	data := layer.Data.(*KurinEventLayerToolData)
	toolData := data.Layer.Data.(*tool.KurinRendererLayerToolData)
	if event.EventManagerInstance.Keyboard.Pending != nil {
		switch *event.EventManagerInstance.Keyboard.Pending {
		case sdl.K_ESCAPE:
			gfx.RendererInstance.Context.State = gfx.RendererContextStateNone
			event.EventManagerInstance.Keyboard.Pending = nil
			return
		}
	}
	if event.EventManagerInstance.Mouse.PendingRight != nil {
		gfx.RendererInstance.Context.State = gfx.RendererContextStateNone
		event.EventManagerInstance.Mouse.PendingRight = nil
	}

	mouseRect := render.ScreenToWorldRect(sdl.Rect{X: gfx.RendererInstance.Context.MousePosition.X, Y: gfx.RendererInstance.Context.MousePosition.Y, W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y})
	mousePosition := sdlutils.Vector3{Base: sdl.Point{X: mouseRect.X, Y: mouseRect.Y}, Z: 0}
	tile := gameplay.GetKurinTileAt(&gameplay.GameInstance.Map, mousePosition)
	switch realPrefab := toolData.Prefab.(type) {
	case *gameplay.KurinObject:
		realPrefab.Tile = tile
		if realPrefab.Tile == nil || !gameplay.CanBuildKurinObjectAtMapPosition(&gameplay.GameInstance.Map, tile.Position) {
			return
		}
		if event.EventManagerInstance.Mouse.PendingLeft != nil {
			job := gameplay.NewKurinJobDriverBuild()
			job.Tile = realPrefab.Tile
			job.Initialize(job, &gameplay.KurinJobDriverBuildData{
				Prefab: realPrefab.Type,
			})
			gameplay.PushKurinJobToController(gameplay.GameInstance.JobController, job)
		}
	case *gameplay.KurinTile:
		realPrefab.Position = mousePosition
		if !gameplay.CanBuildKurinTileAtMapPosition(&gameplay.GameInstance.Map, mousePosition) {
			return
		}
		if event.EventManagerInstance.Mouse.PendingLeft != nil {
			gameplay.CreateKurinTile(realPrefab.Position, realPrefab.Type)
		}
	}
}

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

type EventLayerToolData struct {
	Layer *gfx.RendererLayer
}

func NewEventLayerTool(layer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerTool,
		Process: ProcessEventLayerTool,
		Data: &EventLayerToolData{
			Layer: layer,
		},
	}
}

func LoadEventLayerTool(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerTool(layer *event.EventLayer) error {
	if gfx.RendererInstance.Context.State != gfx.RendererContextStateTool {
		return nil
	}

	ProcessEventLayerToolInput(layer)
	return nil
}

func ProcessEventLayerToolInput(layer *event.EventLayer) {
	data := layer.Data.(*EventLayerToolData)
	toolData := data.Layer.Data.(*tool.RendererLayerToolData)
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

	switch toolData.Mode {
	case tool.ToolModeBuild:
		switch realPrefab := toolData.Prefab.(type) {
		case *gameplay.Object:
			realPrefab.Tile = gameplay.GameInstance.HoveredTile
			if realPrefab.Tile == nil {
				return
			}
			if !gameplay.CanBuildObjectAtMapPosition(gameplay.GameInstance.Map, realPrefab.Tile.Position) && !gameplay.GameInstance.Godmode {
				return
			}
			if event.EventManagerInstance.Mouse.PendingLeft != nil {
				if gameplay.GameInstance.Godmode {
					gameplay.ReplaceObjectRaw(gameplay.GameInstance.Map, gameplay.GetTileAt(gameplay.GameInstance.Map, realPrefab.Tile.Position), realPrefab.Type)
					return
				}

				job := gameplay.NewJobDriver("build", realPrefab.Tile)
				job.Data = &gameplay.JobDriverBuildData{
					ObjectType: realPrefab.Type,
				}
				job.Template.Initialize(job)
				gameplay.PushJobToController(gameplay.GameInstance.JobController[gameplay.FactionPlayer], job)
			}
		case *gameplay.Tile:
			realPrefab.Position = sdlutils.Vector3{Base: render.ScreenToWorldPosition(gfx.RendererInstance.Context.MousePosition), Z: gameplay.GameInstance.SelectedZ}
			if event.EventManagerInstance.Mouse.PendingLeft == nil {
				return
			}
			if gameplay.IsMapPositionOutOfBounds(gameplay.GameInstance.Map, realPrefab.Position) {
				return
			}
			if !gameplay.CanBuildTileAtMapPosition(gameplay.GameInstance.Map, realPrefab.Position) && !gameplay.GameInstance.Godmode {
				return
			}
			if gameplay.DoesBuildFloorJobExistAtPosition(realPrefab.Position) {
				return
			}
			if gameplay.GameInstance.Godmode {
				tile := gameplay.GetTileAt(gameplay.GameInstance.Map, realPrefab.Position)
				if tile != nil {
					tile.Type = realPrefab.Type
				} else {
					gameplay.CreateTileRaw(gameplay.GameInstance.Map, realPrefab.Position, realPrefab.Type)
				}

				return
			}
			job := gameplay.NewJobDriver("build_floor", nil)
			job.Data = &gameplay.JobDriverBuildFloorData{
				Position: realPrefab.Position,
				TileType: realPrefab.Type,
			}
			job.Template.Initialize(job)
			gameplay.PushJobToController(gameplay.GameInstance.JobController[gameplay.FactionPlayer], job)
		}
	case tool.ToolModeDestroy:
		if event.EventManagerInstance.Mouse.PendingLeft == nil {
			return
		}
		if gameplay.GameInstance.HoveredObject != nil {
			if gameplay.GameInstance.Godmode {
				gameplay.DestroyObjectRaw(gameplay.GameInstance.Map, gameplay.GameInstance.HoveredObject)
				return
			}
			job := gameplay.NewJobDriver("destroy", gameplay.GameInstance.HoveredObject.Tile)
			job.Template.Initialize(job)
			gameplay.PushJobToController(gameplay.GameInstance.JobController[gameplay.FactionPlayer], job)
			return
		}
		if gameplay.GameInstance.HoveredTile != nil {
			if !gameplay.CanDestroyTileAtMapPosition(gameplay.GameInstance.Map, gameplay.GameInstance.HoveredTile.Position) {
				return
			}
			if gameplay.GameInstance.Godmode {
				gameplay.DestroyTileRaw(gameplay.GameInstance.Map, gameplay.GameInstance.HoveredTile)
				return
			}
			job := gameplay.NewJobDriver("destroy_floor", gameplay.GameInstance.HoveredTile)
			job.Template.Initialize(job)
			gameplay.PushJobToController(gameplay.GameInstance.JobController[gameplay.FactionPlayer], job)
			return
		}
	}
}

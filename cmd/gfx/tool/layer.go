package tool

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"github.com/veandco/go-sdl2/sdl"
)

type RendererLayerToolData struct {
	Mode   ToolMode
	Prefab interface{}

	TurfLayer   *gfx.RendererLayer
	ObjectLayer *gfx.RendererLayer
}

type ToolMode uint8

const (
	ToolModeBuild   = ToolMode(0)
	ToolModeDestroy = ToolMode(1)
)

func NewRendererLayerTool(turfLayer *gfx.RendererLayer, objectLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerTool,
		Render: RenderRendererLayerTool,
		Data: &RendererLayerToolData{
			Mode:        ToolModeBuild,
			TurfLayer:   turfLayer,
			ObjectLayer: objectLayer,
		},
	}
}

func LoadRendererLayerTool(layer *gfx.RendererLayer) error {
	return nil
}

func RenderRendererLayerTool(layer *gfx.RendererLayer) error {
	if gfx.RendererInstance.Context.State != gfx.RendererContextStateTool {
		return nil
	}
	data := layer.Data.(*RendererLayerToolData)
	switch data.Mode {
	case ToolModeBuild:
		switch realPrefab := data.Prefab.(type) {
		case *gameplay.Object:
			if realPrefab.Tile == nil {
				break
			}
			color := sdlutils.White
			if !gameplay.CanBuildObjectAtMapPosition(gameplay.GameInstance.Map, realPrefab.Tile.Position) {
				color = sdl.Color{R: 128, G: 0, B: 0}
			}
			if err := structure.RenderObjectBlueprint(data.ObjectLayer, realPrefab, color); err != nil {
				return err
			}
		case *gameplay.Tile:
			color := sdlutils.White
			if !gameplay.CanBuildTileAtMapPosition(gameplay.GameInstance.Map, realPrefab.Position) {
				color = sdl.Color{R: 128, G: 0, B: 0}
			}
			if err := turf.RenderTileBlueprint(data.TurfLayer, realPrefab, color); err != nil {
				return err
			}
		}
	case ToolModeDestroy:
		if gameplay.GameInstance.HoveredObject != nil {
			rect := sdlutils.ScaleRectCentered(structure.GetObjectRect(data.ObjectLayer, gameplay.GameInstance.HoveredObject), 0.8)
			sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, sdlutils.GetTextureFromContainer(gfx.RendererInstance.IconTextures, gfx.RendererInstance.Renderer, "delete"), rect)
		} else if gameplay.GameInstance.HoveredTile != nil {
			if !gameplay.CanDestroyTileAtMapPosition(gameplay.GameInstance.Map, gameplay.GameInstance.HoveredTile.Position) {
				break
			}
			rect := sdlutils.ScaleRectCentered(turf.GetTileRect(gameplay.GameInstance.HoveredTile), 0.8)
			sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, sdlutils.GetTextureFromContainer(gfx.RendererInstance.IconTextures, gfx.RendererInstance.Renderer, "delete_floor"), rect)
		}
	}

	return nil
}

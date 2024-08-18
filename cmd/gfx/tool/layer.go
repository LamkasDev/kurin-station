package tool

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRendererLayerToolData struct {
	Mode   KurinToolMode
	Prefab interface{}

	TurfLayer   *gfx.RendererLayer
	ObjectLayer *gfx.RendererLayer
}

type KurinToolMode uint8

const KurinToolModeBuild = KurinToolMode(0)

func NewKurinRendererLayerTool(turfLayer *gfx.RendererLayer, objectLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerTool,
		Render: RenderKurinRendererLayerTool,
		Data: &KurinRendererLayerToolData{
			Mode:        KurinToolModeBuild,
			TurfLayer:   turfLayer,
			ObjectLayer: objectLayer,
		},
	}
}

func LoadKurinRendererLayerTool(layer *gfx.RendererLayer) error {
	return nil
}

func RenderKurinRendererLayerTool(layer *gfx.RendererLayer) error {
	if gfx.RendererInstance.Context.State != gfx.KurinRendererContextStateTool {
		return nil
	}
	data := layer.Data.(*KurinRendererLayerToolData)
	switch data.Mode {
	case KurinToolModeBuild:
		switch realPrefab := data.Prefab.(type) {
		case *gameplay.KurinObject:
			if realPrefab.Tile == nil {
				return nil
			}
			color := sdlutils.White
			if !gameplay.CanEnterKurinTile(realPrefab.Tile) {
				color = sdl.Color{R: 128, G: 0, B: 0}
			}
			if err := structure.RenderKurinObjectBlueprint(data.ObjectLayer, realPrefab, color); err != nil {
				return err
			}
		case *gameplay.KurinTile:
			color := sdlutils.White
			if !gameplay.CanBuildKurinTileAtMapPosition(&gameplay.GameInstance.Map, realPrefab.Position) {
				color = sdl.Color{R: 128, G: 0, B: 0}
			}
			if err := turf.RenderKurinTileBlueprint(data.TurfLayer, realPrefab, color); err != nil {
				return err
			}
		}
	}

	return nil
}

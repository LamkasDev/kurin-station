package tool

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRendererLayerToolData struct {
	Mode   KurinToolMode
	Prefab *gameplay.KurinObject

	ObjectLayer *gfx.KurinRendererLayer
}

type KurinToolMode uint8

const KurinToolModeBuild = KurinToolMode(0)

func NewKurinRendererLayerTool(objectLayer *gfx.KurinRendererLayer) *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerTool,
		Render: RenderKurinRendererLayerTool,
		Data: KurinRendererLayerToolData{
			Mode:        KurinToolModeBuild,
			ObjectLayer: objectLayer,
		},
	}
}

func LoadKurinRendererLayerTool(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	return nil
}

func RenderKurinRendererLayerTool(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	if renderer.Context.State != gfx.KurinRendererContextStateTool {
		return nil
	}
	data := layer.Data.(KurinRendererLayerToolData)
	if data.Mode == KurinToolModeBuild {
		if data.Prefab.Tile == nil {
			return nil
		}
		color := sdlutils.White
		if !gameplay.CanEnterKurinTile(data.Prefab.Tile) {
			color = sdl.Color{R: 128, G: 0, B: 0}
		}
		if err := structure.RenderKurinObjectBlueprint(renderer, data.ObjectLayer, data.Prefab, color); err != nil {
			return err
		}
	}

	return nil
}

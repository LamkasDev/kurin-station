package tool

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
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

func LoadKurinRendererLayerTool(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	return nil
}

func RenderKurinRendererLayerTool(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	if renderer.RendererContext.State != gfx.KurinRendererContextStateTool {
		return nil
	}
	data := layer.Data.(KurinRendererLayerToolData)

	if data.Mode == KurinToolModeBuild {
		// sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Container[gfx.KurinRendererFontDefault], "Build", sdl.Color{R: 255, G: 255, B: 255}, sdl.Point{X: renderer.WindowContext.MousePosition.X - 16, Y: renderer.WindowContext.MousePosition.Y - 16})
		wrect := render.ScreenToWorldRect(renderer, sdl.Rect{X: renderer.RendererContext.MousePosition.X, Y: renderer.RendererContext.MousePosition.Y, W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y})
		tile := gameplay.GetTileAt(&game.Map, sdlutils.Vector3{Base: sdl.Point{X: wrect.X, Y: wrect.Y}, Z: 0})
		if tile == nil {
			return nil
		}

		// color := sdl.Color{R: 64, G: 64, B: 255}
		color := sdl.Color{R: 255, G: 255, B: 255}
		if !gameplay.CanEnterKurinTile(tile) {
			color = sdl.Color{R: 128, G: 0, B: 0}
		}

		if err := structure.RenderKurinObjectBlueprint(renderer, data.ObjectLayer, tile, data.Prefab, color); err != nil {
			return err
		}
	}

	return nil
}

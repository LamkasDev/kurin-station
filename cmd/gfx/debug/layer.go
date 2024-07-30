package debug

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRendererLayerDebugData struct {
	Path *gameplay.KurinPath
}

func NewKurinRendererLayerDebug() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerDebug,
		Render: RenderKurinRendererLayerDebug,
		Data:   KurinRendererLayerDebugData{},
	}
}

func LoadKurinRendererLayerDebug(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	return nil
}

func RenderKurinRendererLayerDebug(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	data := layer.Data.(KurinRendererLayerDebugData)
	if err := renderer.Renderer.SetDrawColor(0, 255, 0, 0); err != nil {
		return &err
	}
	sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Container[gfx.KurinRendererFontDefault], fmt.Sprintf("Camera Mode: %v", renderer.RendererContext.CameraMode), sdl.Color{R: 255, G: 255, B: 255}, sdl.Point{X: 10, Y: 10}, sdl.FPoint{X: 1, Y: 1})
	/* if err := renderer.Renderer.DrawLine(int32(renderer.WindowContext.WindowSize.Float.X/2), 0, int32(renderer.WindowContext.WindowSize.Float.X/2), renderer.WindowContext.WindowSize.Integer.Y); err != nil {
		return &err
	}
	if err := renderer.Renderer.DrawLine(0, int32(renderer.WindowContext.WindowSize.Float.Y/2), renderer.WindowContext.WindowSize.Integer.X, int32(renderer.WindowContext.WindowSize.Float.Y/2)); err != nil {
		return &err
	} */
	wrect := render.ScreenToWorldRect(renderer, sdl.Rect{X: renderer.RendererContext.MousePosition.X, Y: renderer.RendererContext.MousePosition.Y, W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y})
	/* srect := camera.WorldToScreenRect(renderer, sdlutils.RectToFRect(wrect))
	if err := renderer.Renderer.DrawRect(&srect); err != nil {
		return &err
	} */
	tile := gameplay.GetTileAt(&game.Map, sdlutils.Vector3{Base: sdl.Point{X: wrect.X, Y: wrect.Y}, Z: 0})
	if tile != nil {
		sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Container[gfx.KurinRendererFontDefault], gameplay.GetKurinTileDescription(tile), sdl.Color{R: 255, G: 255, B: 255}, sdl.Point{X: 10, Y: 40}, sdl.FPoint{X: 1, Y: 1})
	}
	if data.Path != nil {
		var prevNode *gameplay.KurinPathfindingNode = nil
		for _, currentNode := range data.Path.Nodes {
			if prevNode != nil {
				prevRect := turf.GetKurinTileRect(renderer, prevNode.Tile, sdl.FPoint{X: 0.5, Y: 0.5})
				currentRect := turf.GetKurinTileRect(renderer, currentNode.Tile, sdl.FPoint{X: 0.5, Y: 0.5})
				renderer.Renderer.DrawLine(prevRect.X, prevRect.Y, currentRect.X, currentRect.Y)
			}
			prevNode = currentNode
		}
	}

	return nil
}

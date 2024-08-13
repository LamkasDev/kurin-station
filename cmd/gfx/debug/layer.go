package debug

import (
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

func LoadKurinRendererLayerDebug(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	return nil
}

func RenderKurinRendererLayerDebug(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	data := layer.Data.(KurinRendererLayerDebugData)
	if err := renderer.Renderer.SetDrawColor(0, 255, 0, 0); err != nil {
		return err
	}
	/* if err := renderer.Renderer.DrawLine(int32(renderer.WindowContext.WindowSize.Float.X/2), 0, int32(renderer.WindowContext.WindowSize.Float.X/2), renderer.WindowContext.WindowSize.Integer.Y); err != nil {
		return err
	}
	if err := renderer.Renderer.DrawLine(0, int32(renderer.WindowContext.WindowSize.Float.Y/2), renderer.WindowContext.WindowSize.Integer.X, int32(renderer.WindowContext.WindowSize.Float.Y/2)); err != nil {
		return err
	} */
	wrect := render.ScreenToWorldRect(renderer, sdl.Rect{X: renderer.Context.MousePosition.X, Y: renderer.Context.MousePosition.Y, W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y})
	/* srect := camera.WorldToScreenRect(renderer, sdlutils.RectToFRect(wrect))
	if err := renderer.Renderer.DrawRect(&srect); err != nil {
		return err
	} */
	tile := gameplay.GetTileAt(&gameplay.KurinGameInstance.Map, sdlutils.Vector3{Base: sdl.Point{X: wrect.X, Y: wrect.Y}, Z: 0})
	if tile != nil {
		sdlutils.RenderLabel(renderer.Renderer, "debug", renderer.Fonts.Default, sdlutils.White, gameplay.GetKurinTileDescription(tile), sdl.Point{X: 10, Y: renderer.Context.WindowSize.Y - 24}, sdl.FPoint{X: 1, Y: 1})
	}
	if data.Path != nil {
		var prevNode *gameplay.KurinPathfindingNode = nil
		for _, currentNode := range data.Path.Nodes {
			if prevNode != nil {
				prevRect := turf.GetKurinTileRectDebug(renderer, prevNode.Tile, sdl.FPoint{X: 0.5, Y: 0.5})
				currentRect := turf.GetKurinTileRectDebug(renderer, currentNode.Tile, sdl.FPoint{X: 0.5, Y: 0.5})
				renderer.Renderer.DrawLine(prevRect.X, prevRect.Y, currentRect.X, currentRect.Y)
			}
			prevNode = currentNode
		}
	}

	return nil
}

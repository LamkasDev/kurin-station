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

func NewKurinRendererLayerDebug() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerDebug,
		Render: RenderKurinRendererLayerDebug,
		Data:   &KurinRendererLayerDebugData{},
	}
}

func LoadKurinRendererLayerDebug(layer *gfx.RendererLayer) error {
	return nil
}

func RenderKurinRendererLayerDebug(layer *gfx.RendererLayer) error {
	data := layer.Data.(*KurinRendererLayerDebugData)
	if err := gfx.RendererInstance.Renderer.SetDrawColor(0, 255, 0, 0); err != nil {
		return err
	}
	/* if err := gfx.KurinRendererInstance.Renderer.DrawLine(int32(gfx.KurinRendererInstance.WindowContext.WindowSize.Float.X/2), 0, int32(gfx.KurinRendererInstance.WindowContext.WindowSize.Float.X/2), gfx.KurinRendererInstance.WindowContext.WindowSize.Integer.Y); err != nil {
		return err
	}
	if err := gfx.KurinRendererInstance.Renderer.DrawLine(0, int32(gfx.KurinRendererInstance.WindowContext.WindowSize.Float.Y/2), gfx.KurinRendererInstance.WindowContext.WindowSize.Integer.X, int32(gfx.KurinRendererInstance.WindowContext.WindowSize.Float.Y/2)); err != nil {
		return err
	} */
	wrect := render.ScreenToWorldRect(sdl.Rect{X: gfx.RendererInstance.Context.MousePosition.X, Y: gfx.RendererInstance.Context.MousePosition.Y, W: gameplay.KurinTileSize.X, H: gameplay.KurinTileSize.Y})
	/* srect := camera.WorldToScreenRect(renderer, sdlutils.RectToFRect(wrect))
	if err := gfx.KurinRendererInstance.Renderer.DrawRect(&srect); err != nil {
		return err
	} */
	tile := gameplay.GetKurinTileAt(&gameplay.GameInstance.Map, sdlutils.Vector3{Base: sdl.Point{X: wrect.X, Y: wrect.Y}, Z: 0})
	if tile != nil {
		sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "debug", gfx.RendererInstance.Fonts.Default, sdlutils.White, gameplay.GetKurinTileDescription(tile), sdl.Point{X: 10, Y: gfx.RendererInstance.Context.WindowSize.Y - 24}, sdl.FPoint{X: 1, Y: 1})
	}
	if data.Path != nil {
		var prevNode *gameplay.KurinPathfindingNode = nil
		for _, currentNode := range data.Path.Nodes {
			if prevNode != nil {
				prevTile := gameplay.GetKurinTileAt(&gameplay.GameInstance.Map, prevNode.Position)
				prevRect := turf.GetKurinTileRectDebug(prevTile, sdl.FPoint{X: 0.5, Y: 0.5})
				currentTile := gameplay.GetKurinTileAt(&gameplay.GameInstance.Map, currentNode.Position)
				currentRect := turf.GetKurinTileRectDebug(currentTile, sdl.FPoint{X: 0.5, Y: 0.5})
				gfx.RendererInstance.Renderer.DrawLine(prevRect.X, prevRect.Y, currentRect.X, currentRect.Y)
			}
			prevNode = currentNode
		}
	}

	return nil
}

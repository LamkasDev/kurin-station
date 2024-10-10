package debug

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"github.com/veandco/go-sdl2/sdl"
)

type RendererLayerDebugData struct {
	Path *gameplay.Path
}

func NewRendererLayerDebug() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerDebug,
		Render: RenderRendererLayerDebug,
		Data:   &RendererLayerDebugData{},
	}
}

func LoadRendererLayerDebug(layer *gfx.RendererLayer) error {
	return nil
}

func RenderRendererLayerDebug(layer *gfx.RendererLayer) error {
	data := layer.Data.(*RendererLayerDebugData)
	if err := gfx.RendererInstance.Renderer.SetDrawColor(0, 255, 0, 0); err != nil {
		return err
	}
	text := ""
	if gameplay.GameInstance.Godmode {
		text += "GODMODE "
	}
	if gameplay.GameInstance.HoveredTile != nil {
		text += gameplay.GetTileDescription(gameplay.GameInstance.HoveredTile)
	}
	if gameplay.GameInstance.HoveredItem != nil {
		text = fmt.Sprintf("%s %s", text, gameplay.GetItemDescription(gameplay.GameInstance.HoveredItem))
	}
	if gameplay.GameInstance.HoveredMob != nil {
		text = fmt.Sprintf("%s %s", text, gameplay.GetMobDescription(gameplay.GameInstance.HoveredMob))
	}
	if text != "" {
		sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "debug", gfx.RendererInstance.Fonts.Default, sdlutils.White, text, sdl.Point{X: 10, Y: gfx.RendererInstance.Context.WindowSize.Y - 24}, sdl.FPoint{X: 1, Y: 1})
	}
	if data.Path != nil {
		var prevNode *gameplay.PathfindingNode = nil
		for _, currentNode := range data.Path.Nodes {
			if prevNode != nil {
				prevTile := gameplay.GetTileAt(gameplay.GameInstance.Map, prevNode.Position)
				prevRect := turf.GetTileRectDebug(prevTile, sdl.FPoint{X: 0.5, Y: 0.5})
				currentTile := gameplay.GetTileAt(gameplay.GameInstance.Map, currentNode.Position)
				currentRect := turf.GetTileRectDebug(currentTile, sdl.FPoint{X: 0.5, Y: 0.5})
				gfx.RendererInstance.Renderer.DrawLine(prevRect.X, prevRect.Y, currentRect.X, currentRect.Y)
			}
			prevNode = currentNode
		}
	}

	return nil
}

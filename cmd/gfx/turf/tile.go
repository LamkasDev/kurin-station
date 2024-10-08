package turf

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetTileRectDebug(tile *gameplay.Tile, offset sdl.FPoint) sdl.Rect {
	return render.WorldToScreenRectF(sdl.FRect{
		X: float32(tile.Position.Base.X) + offset.X, Y: float32(tile.Position.Base.Y) + offset.Y,
		W: gameplay.TileSizeF.X, H: gameplay.TileSizeF.Y,
	})
}

func GetTileRect(tile *gameplay.Tile) sdl.Rect {
	return render.WorldToScreenRectF(sdl.FRect{
		X: float32(tile.Position.Base.X), Y: float32(tile.Position.Base.Y),
		W: gameplay.TileSizeF.X, H: gameplay.TileSizeF.Y,
	})
}

func RenderTile(layer *gfx.RendererLayer, tile *gameplay.Tile, blur bool) error {
	rect := GetTileRect(tile)
	if gfx.ShouldOcclude(rect) {
		return nil
	}
	graphic := layer.Data.(*RendererLayerTileData).Turfs[tile.Type]
	if blur {
		if err := gfx.RendererInstance.Renderer.Copy(graphic.BlurTextures[0].Texture, nil, &rect); err != nil {
			return err
		}

		return nil
	}
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Textures[tile.Template.GetTexture(tile)].Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

func RenderTileBlueprint(layer *gfx.RendererLayer, tile *gameplay.Tile, color sdl.Color) error {
	rect := GetTileRect(tile)
	if gfx.ShouldOcclude(rect) {
		return nil
	}
	graphic := layer.Data.(*RendererLayerTileData).Turfs[tile.Type]
	if graphic.Blueprint == nil {
		return nil
	}
	graphic.Blueprint.Texture.SetColorMod(color.R, color.G, color.B)
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Blueprint.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

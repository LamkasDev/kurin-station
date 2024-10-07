package turf

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetTileRectDebug(tile *gameplay.Tile, offset sdl.FPoint) sdl.Rect {
	return render.WorldToScreenRect(sdl.FRect{
		X: float32(tile.Position.Base.X) + offset.X, Y: float32(tile.Position.Base.Y) + offset.Y,
		W: gameplay.TileSizeF.X, H: gameplay.TileSizeF.Y,
	})
}

func GetTileRect(tile *gameplay.Tile) sdl.Rect {
	return render.WorldToScreenRect(sdl.FRect{
		X: float32(tile.Position.Base.X), Y: float32(tile.Position.Base.Y),
		W: gameplay.TileSizeF.X, H: gameplay.TileSizeF.Y,
	})
}

func RenderTile(layer *gfx.RendererLayer, tile *gameplay.Tile, blur bool) error {
	graphic := layer.Data.(*RendererLayerTileData).Turfs[tile.Type]
	texture := graphic.Textures[0].Texture
	if blur {
		texture = graphic.BlurTextures[0].Texture
	}
	rect := GetTileRect(tile)
	if err := gfx.RendererInstance.Renderer.Copy(texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

func RenderTileBlueprint(layer *gfx.RendererLayer, tile *gameplay.Tile, color sdl.Color) error {
	graphic := layer.Data.(*RendererLayerTileData).Turfs[tile.Type]
	if graphic.Blueprint == nil {
		return nil
	}
	rect := GetTileRect(tile)
	graphic.Blueprint.Texture.SetColorMod(color.R, color.G, color.B)
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Blueprint.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

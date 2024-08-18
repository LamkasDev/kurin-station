package turf

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinTileRectDebug(tile *gameplay.KurinTile, offset sdl.FPoint) sdl.Rect {
	return render.WorldToScreenRect(sdl.FRect{
		X: float32(tile.Position.Base.X) + offset.X, Y: float32(tile.Position.Base.Y) + offset.Y,
		W: gameplay.KurinTileSizeF.X, H: gameplay.KurinTileSizeF.Y,
	})
}

func GetKurinTileRect(tile *gameplay.KurinTile) sdl.Rect {
	return render.WorldToScreenRect(sdl.FRect{
		X: float32(tile.Position.Base.X), Y: float32(tile.Position.Base.Y),
		W: gameplay.KurinTileSizeF.X, H: gameplay.KurinTileSizeF.Y,
	})
}

func RenderKurinTile(layer *gfx.RendererLayer, tile *gameplay.KurinTile) error {
	graphic := layer.Data.(*KurinRendererLayerTileData).Turfs[tile.Type]
	rect := GetKurinTileRect(tile)
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Textures[0].Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

func RenderKurinTileBlueprint(layer *gfx.RendererLayer, tile *gameplay.KurinTile, color sdl.Color) error {
	graphic := layer.Data.(*KurinRendererLayerTileData).Turfs[tile.Type]
	rect := GetKurinTileRect(tile)
	graphic.Blueprint.Texture.SetColorMod(color.R, color.G, color.B)
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Blueprint.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

package turf

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinTileRect(renderer *gfx.KurinRenderer, tile *gameplay.KurinTile, offset sdl.FPoint) sdl.Rect {
	return render.WorldToScreenRect(renderer, sdl.FRect{
		X: float32(tile.Position.Base.X) + offset.X, Y: float32(tile.Position.Base.Y) + offset.Y,
		W: float32(gameplay.KurinTileSize.X), H: float32(gameplay.KurinTileSize.Y),
	})
}

func RenderKurinTile(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, tile *gameplay.KurinTile) *error {
	graphic := layer.Data.(KurinRendererLayerTileData).Turfs[tile.Type]
	texture := graphic.Textures[tile.Direction]
	rect := GetKurinTileRect(renderer, tile, sdl.FPoint{})
	if err := renderer.Renderer.Copy(texture.Texture, nil, &rect); err != nil {
		return &err
	}

	return nil
}

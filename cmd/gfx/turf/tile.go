package turf

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

var TileRect sdl.Rect

func GetTileRectDebug(tile *gameplay.Tile, offset sdl.FPoint) sdl.Rect {
	pos := render.WorldToScreenPositionF(sdl.FPoint{
		X: float32(tile.Position.Base.X) + offset.X,
		Y: float32(tile.Position.Base.Y) + offset.Y,
	})

	return sdl.Rect{
		X: pos.X, Y: pos.Y,
		W: gfx.RendererInstance.Context.CameraTileSize.X, H: gfx.RendererInstance.Context.CameraTileSize.Y,
	}
}

func GetTileRect(tile *gameplay.Tile) sdl.Rect {
	pos := render.WorldToScreenPosition(tile.Position.Base)

	return sdl.Rect{
		X: pos.X, Y: pos.Y,
		W: gfx.RendererInstance.Context.CameraTileSize.X, H: gfx.RendererInstance.Context.CameraTileSize.Y,
	}
}

func RenderTile(layer *gfx.RendererLayer, tile *gameplay.Tile, blur bool) error {
	TileRect = GetTileRect(tile)
	if gfx.ShouldOcclude(TileRect) {
		return nil
	}
	graphic := layer.Data.(*RendererLayerTileData).Turfs[tile.Type]
	if blur {
		if err := gfx.RendererInstance.Renderer.Copy(graphic.BlurTextures[0].Texture, nil, &TileRect); err != nil {
			return err
		}

		return nil
	}
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Textures[tile.Template.GetTexture(tile)].Texture, nil, &TileRect); err != nil {
		return err
	}

	return nil
}

func RenderTileBlueprint(layer *gfx.RendererLayer, tile *gameplay.Tile, color sdl.Color) error {
	TileRect = GetTileRect(tile)
	if gfx.ShouldOcclude(TileRect) {
		return nil
	}
	graphic := layer.Data.(*RendererLayerTileData).Turfs[tile.Type]
	if graphic.Blueprint == nil {
		return nil
	}
	graphic.Blueprint.Texture.SetColorMod(color.R, color.G, color.B)
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Blueprint.Texture, nil, &TileRect); err != nil {
		return err
	}

	return nil
}

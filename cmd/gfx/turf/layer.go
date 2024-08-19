package turf

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerTileData struct {
	Turfs map[string]*TurfGraphic
}

func NewRendererLayerTile() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerTile,
		Render: RenderRendererLayerTile,
		Data: &RendererLayerTileData{
			Turfs: map[string]*TurfGraphic{},
		},
	}
}

func LoadRendererLayerTile(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*RendererLayerTileData).Turfs["floor"], err = NewTurfGraphic("floor"); err != nil {
		return err
	}
	if layer.Data.(*RendererLayerTileData).Turfs["blank"], err = NewTurfGraphic("blank"); err != nil {
		return err
	}
	if layer.Data.(*RendererLayerTileData).Turfs["catwalk"], err = NewTurfGraphic("catwalk"); err != nil {
		return err
	}

	return nil
}

func RenderRendererLayerTile(layer *gfx.RendererLayer) error {
	for x := range gameplay.GameInstance.Map.Size.Base.X {
		for y := range gameplay.GameInstance.Map.Size.Base.Y {
			tile := gameplay.GameInstance.Map.Tiles[x][y][0]
			if tile == nil {
				continue
			}
			RenderTile(layer, tile)
		}
	}

	return nil
}

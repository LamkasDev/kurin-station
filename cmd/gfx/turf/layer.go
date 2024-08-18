package turf

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerTileData struct {
	Turfs map[string]*KurinTurfGraphic
}

func NewKurinRendererLayerTile() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerTile,
		Render: RenderKurinRendererLayerTile,
		Data: &KurinRendererLayerTileData{
			Turfs: map[string]*KurinTurfGraphic{},
		},
	}
}

func LoadKurinRendererLayerTile(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*KurinRendererLayerTileData).Turfs["floor"], err = NewKurinTurfGraphic("floor"); err != nil {
		return err
	}
	if layer.Data.(*KurinRendererLayerTileData).Turfs["blank"], err = NewKurinTurfGraphic("blank"); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerTile(layer *gfx.RendererLayer) error {
	for x := range gameplay.GameInstance.Map.Size.Base.X {
		for y := range gameplay.GameInstance.Map.Size.Base.Y {
			tile := gameplay.GameInstance.Map.Tiles[x][y][0]
			if tile == nil {
				continue
			}
			RenderKurinTile(layer, tile)
		}
	}

	return nil
}

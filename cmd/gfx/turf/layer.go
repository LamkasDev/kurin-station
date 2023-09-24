package turf

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerTileData struct {
	Turfs map[string]*KurinTurfGraphic
}

func NewKurinRendererLayerTile() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerTile,
		Render: RenderKurinRendererLayerTile,
		Data: KurinRendererLayerTileData{
			Turfs: map[string]*KurinTurfGraphic{},
		},
	}
}

func LoadKurinRendererLayerTile(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	var err *error
	if layer.Data.(KurinRendererLayerTileData).Turfs["floor"], err = NewKurinTurfGraphic(renderer, "floor"); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerTile(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	for x := int32(0); x < game.Map.Size.Base.X; x++ {
		for y := int32(0); y < game.Map.Size.Base.Y; y++ {
			RenderKurinTile(renderer, layer, game.Map.Tiles[x][y][0])
		}
	}

	return nil
}

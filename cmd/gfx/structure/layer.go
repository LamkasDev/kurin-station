package structure

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerObjectData struct {
	Structures map[string]*KurinStructureGraphic
}

func NewKurinRendererLayerObject() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerObject,
		Render: RenderKurinRendererLayerObject,
		Data: KurinRendererLayerObjectData{
			Structures: map[string]*KurinStructureGraphic{},
		},
	}
}

func LoadKurinRendererLayerObject(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	var err *error
	if layer.Data.(KurinRendererLayerObjectData).Structures["grille"], err = NewKurinStructureGraphic(renderer, "grille"); err != nil {
		return err
	}
	if layer.Data.(KurinRendererLayerObjectData).Structures["displaced"], err = NewKurinStructureGraphic(renderer, "displaced"); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerObject(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	for x := int32(0); x < game.Map.Size.Base.X; x++ {
		for y := int32(0); y < game.Map.Size.Base.Y; y++ {
			tile := game.Map.Tiles[x][y][0]
			for _, obj := range tile.Objects {
				RenderKurinObject(renderer, layer, tile, obj)
			}
		}
	}

	return nil
}

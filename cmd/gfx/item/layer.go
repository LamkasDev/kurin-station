package item

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerItemData struct {
	Items map[string]*KurinItemGraphic
}

func NewKurinRendererLayerItem() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerItem,
		Render: RenderKurinRendererLayerItem,
		Data: KurinRendererLayerItemData{
			Items: map[string]*KurinItemGraphic{},
		},
	}
}

func LoadKurinRendererLayerItem(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	var err *error
	if layer.Data.(KurinRendererLayerItemData).Items["survivalknife"], err = NewKurinItemGraphic(renderer, "survivalknife"); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerItem(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	for _, item := range game.Map.Items {
		RenderKurinItem(renderer, layer, game, item)
	}

	return nil
}

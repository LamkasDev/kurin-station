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

func LoadKurinRendererLayerObject(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	var err error
	if layer.Data.(KurinRendererLayerObjectData).Structures["grille"], err = NewKurinStructureGraphic(renderer, "grille"); err != nil {
		return err
	}
	if layer.Data.(KurinRendererLayerObjectData).Structures["displaced"], err = NewKurinStructureGraphic(renderer, "displaced"); err != nil {
		return err
	}
	if layer.Data.(KurinRendererLayerObjectData).Structures["pod"], err = NewKurinStructureGraphic(renderer, "pod"); err != nil {
		return err
	}
	if layer.Data.(KurinRendererLayerObjectData).Structures["broken_grille"], err = NewKurinStructureGraphic(renderer, "broken_grille"); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerObject(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	for _, obj := range gameplay.KurinGameInstance.Map.Objects {
		RenderKurinObject(renderer, layer, obj)
	}

	return nil
}

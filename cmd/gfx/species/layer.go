package species

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerCharacterData struct {
	Species   map[string]*KurinSpeciesGraphicContainer
	ItemLayer *gfx.RendererLayer
}

func NewKurinRendererLayerCharacter(itemLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerCharacter,
		Render: RenderKurinRendererLayerCharacter,
		Data: &KurinRendererLayerCharacterData{
			Species:   map[string]*KurinSpeciesGraphicContainer{},
			ItemLayer: itemLayer,
		},
	}
}

func LoadKurinRendererLayerCharacter(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*KurinRendererLayerCharacterData).Species[gameplay.KurinDefaultSpecies], err = NewKurinSpeciesGraphicContainer(gameplay.KurinDefaultSpecies); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerCharacter(layer *gfx.RendererLayer) error {
	for _, character := range gameplay.GameInstance.Characters {
		gameplay.ProcessKurinCharacter(character)
		RenderKurinCharacter(layer, character)
	}

	return nil
}

package species

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerCharacterData struct {
	Species   map[string]*SpeciesGraphicContainer
	ItemLayer *gfx.RendererLayer
}

func NewRendererLayerCharacter(itemLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerCharacter,
		Render: RenderRendererLayerCharacter,
		Data: &RendererLayerCharacterData{
			Species:   map[string]*SpeciesGraphicContainer{},
			ItemLayer: itemLayer,
		},
	}
}

func LoadRendererLayerCharacter(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*RendererLayerCharacterData).Species[gameplay.DefaultSpecies], err = NewSpeciesGraphicContainer(gameplay.DefaultSpecies); err != nil {
		return err
	}

	return nil
}

func RenderRendererLayerCharacter(layer *gfx.RendererLayer) error {
	for _, character := range gameplay.GameInstance.Characters {
		RenderCharacter(layer, character)
	}

	return nil
}

package species

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerCharacterData struct {
	Species   map[string]*KurinSpeciesGraphicContainer
	ItemLayer *gfx.KurinRendererLayer
}

func NewKurinRendererLayerCharacter(itemLayer *gfx.KurinRendererLayer) *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerCharacter,
		Render: RenderKurinRendererLayerCharacter,
		Data: KurinRendererLayerCharacterData{
			Species:   map[string]*KurinSpeciesGraphicContainer{},
			ItemLayer: itemLayer,
		},
	}
}

func LoadKurinRendererLayerCharacter(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	var err *error
	if layer.Data.(KurinRendererLayerCharacterData).Species[gameplay.KurinDefaultSpecies], err = NewKurinSpeciesGraphicContainer(renderer, gameplay.KurinDefaultSpecies); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerCharacter(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	for _, character := range game.Characters {
		gameplay.ProcessKurinCharacter(game, character)
		RenderKurinCharacter(renderer, layer, character)
	}

	return nil
}

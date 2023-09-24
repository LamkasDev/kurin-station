package animation

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerAnimationData struct {
	Animations map[string]*KurinAnimationGraphic
}

func NewKurinRendererLayerAnimation() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerAnimation,
		Render: RenderKurinRendererLayerAnimation,
		Data: KurinRendererLayerAnimationData{
			Animations: map[string]*KurinAnimationGraphic{},
		},
	}
}

func LoadKurinRendererLayerAnimation(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	var err *error
	if layer.Data.(KurinRendererLayerAnimationData).Animations["hit"], err = NewKurinAnimationGraphic(renderer, "hit"); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerAnimation(_ *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	for _, character := range game.Characters {
		if character.AnimationController.Animation != nil {
			graphic := layer.Data.(KurinRendererLayerAnimationData).Animations[character.AnimationController.Animation.Type]
			if character.AnimationController.Animation.Step == -1 {
				AdvanceKurinCharacterAnimation(character, &graphic.Template)
			}
			ProcessKurinCharacterAnimation(character, &graphic.Template)
		}
	}

	return nil
}

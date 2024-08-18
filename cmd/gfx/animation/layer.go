package animation

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerAnimationData struct {
	Animations map[string]*KurinAnimationGraphic
}

func NewKurinRendererLayerAnimation() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerAnimation,
		Render: RenderKurinRendererLayerAnimation,
		Data: &KurinRendererLayerAnimationData{
			Animations: map[string]*KurinAnimationGraphic{},
		},
	}
}

func LoadKurinRendererLayerAnimation(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*KurinRendererLayerAnimationData).Animations["hit"], err = NewKurinAnimationGraphic("hit"); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerAnimation(layer *gfx.RendererLayer) error {
	for _, character := range gameplay.GameInstance.Characters {
		if character.AnimationController.Animation != nil {
			graphic := layer.Data.(*KurinRendererLayerAnimationData).Animations[character.AnimationController.Animation.Type]
			if character.AnimationController.Animation.Step == -1 {
				AdvanceKurinCharacterAnimation(character, &graphic.Template)
			}
			ProcessKurinCharacterAnimation(character, &graphic.Template)
		}
	}

	return nil
}

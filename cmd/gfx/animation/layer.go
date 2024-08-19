package animation

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerAnimationData struct {
	Animations map[string]*AnimationGraphic
}

func NewRendererLayerAnimation() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerAnimation,
		Render: RenderRendererLayerAnimation,
		Data: &RendererLayerAnimationData{
			Animations: map[string]*AnimationGraphic{},
		},
	}
}

func LoadRendererLayerAnimation(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*RendererLayerAnimationData).Animations["hit"], err = NewAnimationGraphic("hit"); err != nil {
		return err
	}

	return nil
}

func RenderRendererLayerAnimation(layer *gfx.RendererLayer) error {
	for _, character := range gameplay.GameInstance.Characters {
		if character.AnimationController.Animation != nil {
			graphic := layer.Data.(*RendererLayerAnimationData).Animations[character.AnimationController.Animation.Type]
			if character.AnimationController.Animation.Step == -1 {
				AdvanceCharacterAnimation(character, &graphic.Template)
			}
			ProcessCharacterAnimation(character, &graphic.Template)
		}
	}

	return nil
}

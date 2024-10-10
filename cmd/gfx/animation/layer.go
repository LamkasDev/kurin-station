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
	for _, mob := range gameplay.GameInstance.Map.Mobs {
		if mob.AnimationController.Animation != nil {
			graphic := layer.Data.(*RendererLayerAnimationData).Animations[mob.AnimationController.Animation.Type]
			if mob.AnimationController.Animation.Step == -1 {
				AdvanceMobAnimation(mob, &graphic.Template)
			}
			ProcessMobAnimation(mob, &graphic.Template)
		}
	}

	return nil
}

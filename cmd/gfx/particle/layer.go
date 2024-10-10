package particle

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"golang.org/x/exp/slices"
)

type RendererLayerParticleData struct {
	Particles map[string]*ParticleGraphic
}

func NewRendererLayerParticle() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerParticle,
		Render: RenderRendererLayerParticle,
		Data: &RendererLayerParticleData{
			Particles: map[string]*ParticleGraphic{},
		},
	}
}

func LoadRendererLayerParticle(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*RendererLayerParticleData).Particles["cross"], err = NewParticleGraphic("cross", 1); err != nil {
		return err
	}
	if layer.Data.(*RendererLayerParticleData).Particles["ion"], err = NewParticleGraphic("ion", 3); err != nil {
		return err
	}

	return nil
}

func RenderRendererLayerParticle(layer *gfx.RendererLayer) error {
	if len(gameplay.GameInstance.ParticleController.Pending) > 0 {
		for i := len(gameplay.GameInstance.ParticleController.Pending) - 1; i >= 0; i-- {
			particle := gameplay.GameInstance.ParticleController.Pending[i]
			if particle.Position.Z == gameplay.GameInstance.SelectedZ {
				if err := RenderParticle(layer, particle); err != nil {
					return err
				}
			}

			gameplay.ProcessParticle(particle)
			if particle.Ticks == 0 {
				gameplay.GameInstance.ParticleController.Pending = slices.Delete(gameplay.GameInstance.ParticleController.Pending, i, i+1)
			}
		}
	}

	return nil
}

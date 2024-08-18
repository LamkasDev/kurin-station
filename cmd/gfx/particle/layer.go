package particle

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"golang.org/x/exp/slices"
)

type KurinRendererLayerParticleData struct {
	Particles map[string]*KurinParticleGraphic
}

func NewKurinRendererLayerParticle() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerParticle,
		Render: RenderKurinRendererLayerParticle,
		Data: &KurinRendererLayerParticleData{
			Particles: map[string]*KurinParticleGraphic{},
		},
	}
}

func LoadKurinRendererLayerParticle(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*KurinRendererLayerParticleData).Particles["cross"], err = NewKurinParticleGraphic("cross", 1); err != nil {
		return err
	}
	if layer.Data.(*KurinRendererLayerParticleData).Particles["ion"], err = NewKurinParticleGraphic("ion", 3); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerParticle(layer *gfx.RendererLayer) error {
	if len(gameplay.GameInstance.ParticleController.Pending) > 0 {
		for i := len(gameplay.GameInstance.ParticleController.Pending) - 1; i >= 0; i-- {
			particle := gameplay.GameInstance.ParticleController.Pending[i]
			if err := RenderKurinParticle(layer, particle); err != nil {
				return err
			}

			gameplay.ProcessKurinParticle(particle)
			if particle.Ticks == 0 {
				gameplay.GameInstance.ParticleController.Pending = slices.Delete(gameplay.GameInstance.ParticleController.Pending, i, i+1)
			}
		}
	}

	return nil
}

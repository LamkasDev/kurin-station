package particle

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"golang.org/x/exp/slices"
)

type KurinRendererLayerParticleData struct {
	Particles map[string]*KurinParticleGraphic
}

func NewKurinRendererLayerParticle() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerParticle,
		Render: RenderKurinRendererLayerParticle,
		Data: KurinRendererLayerParticleData{
			Particles: map[string]*KurinParticleGraphic{},
		},
	}
}

func LoadKurinRendererLayerParticle(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	var err error
	if layer.Data.(KurinRendererLayerParticleData).Particles["cross"], err = NewKurinParticleGraphic(renderer, "cross"); err != nil {
		return err
	}

	return nil
}

func RenderKurinRendererLayerParticle(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	if len(gameplay.KurinGameInstance.ParticleController.Pending) > 0 {
		for i := len(gameplay.KurinGameInstance.ParticleController.Pending) - 1; i >= 0; i-- {
			particle := gameplay.KurinGameInstance.ParticleController.Pending[i]
			if err := RenderKurinParticle(renderer, layer, particle); err != nil {
				return err
			}

			gameplay.ProcessKurinParticle(particle)
			if particle.Ticks == 0 {
				gameplay.KurinGameInstance.ParticleController.Pending = slices.Delete(gameplay.KurinGameInstance.ParticleController.Pending, i, i+1)
			}
		}
	}

	return nil
}

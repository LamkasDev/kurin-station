package particle

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinParticleRect(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, particle *gameplay.KurinParticle) sdl.Rect {
	return render.WorldToScreenRect(renderer, sdl.FRect{
		X: particle.Position.Base.X - 0.5 * particle.Scale, Y: particle.Position.Base.Y - 0.5 * particle.Scale,
		W: gameplay.KurinTileSizeF.X * particle.Scale, H: gameplay.KurinTileSizeF.Y * particle.Scale,
	})
}

func RenderKurinParticle(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, particle *gameplay.KurinParticle) error {
	graphic := layer.Data.(KurinRendererLayerParticleData).Particles[particle.Type]
	rect := GetKurinParticleRect(renderer, layer, particle)
	graphic.Texture.Texture.SetColorMod(particle.Color.R, particle.Color.G, particle.Color.B)
	if err := renderer.Renderer.Copy(graphic.Texture.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

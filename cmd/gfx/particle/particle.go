package particle

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinParticleRect(layer *gfx.RendererLayer, particle *gameplay.KurinParticle) sdl.Rect {
	return render.WorldToScreenRect(sdl.FRect{
		X: particle.Position.Base.X - 0.5*particle.Scale, Y: particle.Position.Base.Y - 0.5*particle.Scale,
		W: gameplay.KurinTileSizeF.X * particle.Scale, H: gameplay.KurinTileSizeF.Y * particle.Scale,
	})
}

func RenderKurinParticle(layer *gfx.RendererLayer, particle *gameplay.KurinParticle) error {
	graphic := layer.Data.(*KurinRendererLayerParticleData).Particles[particle.Type]
	texture := graphic.Textures[particle.Index]
	texture.Texture.SetColorMod(particle.Color.R, particle.Color.G, particle.Color.B)
	rect := GetKurinParticleRect(layer, particle)
	if err := gfx.RendererInstance.Renderer.CopyEx(texture.Texture, nil, &rect, particle.Rotation, nil, sdl.FLIP_NONE); err != nil {
		return err
	}

	return nil
}

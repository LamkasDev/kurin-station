package particle

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetParticleRect(layer *gfx.RendererLayer, particle *gameplay.Particle) sdl.Rect {
	return render.WorldToScreenRectF(sdl.FRect{
		X: particle.Position.Base.X - 0.5*particle.Scale, Y: particle.Position.Base.Y - 0.5*particle.Scale,
		W: gameplay.TileSizeF.X * particle.Scale, H: gameplay.TileSizeF.Y * particle.Scale,
	})
}

func RenderParticle(layer *gfx.RendererLayer, particle *gameplay.Particle) error {
	graphic := layer.Data.(*RendererLayerParticleData).Particles[particle.Type]
	texture := graphic.Textures[particle.Index]
	texture.Texture.SetColorMod(particle.Color.R, particle.Color.G, particle.Color.B)
	rect := GetParticleRect(layer, particle)
	if err := gfx.RendererInstance.Renderer.CopyEx(texture.Texture, nil, &rect, particle.Rotation, nil, sdl.FLIP_NONE); err != nil {
		return err
	}

	return nil
}

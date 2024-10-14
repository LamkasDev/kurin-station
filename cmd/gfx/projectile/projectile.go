package projectile

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetProjectileRect(layer *gfx.RendererLayer, projectile *gameplay.Projectile) sdl.Rect {
	return render.WorldToScreenRectF(sdl.FRect{
		X: float32(projectile.Position.Base.X) - 0.5, Y: float32(projectile.Position.Base.Y) - 0.5,
		W: 32, H: 32,
	})
}

func RenderProjectile(layer *gfx.RendererLayer, projectile *gameplay.Projectile) error {
	texture := layer.Data.(*RendererLayerProjectileData).Projectiles[projectile.Type]
	rect := GetProjectileRect(layer, projectile)
	if err := gfx.RendererInstance.Renderer.Copy(texture.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

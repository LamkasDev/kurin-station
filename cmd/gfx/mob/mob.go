package mob

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetMobRect(mob *gameplay.Mob) sdl.Rect {
	position := sdlutils.AddFPoints(mob.PositionRender, gameplay.GetAnimationOffset(mob))
	return render.WorldToScreenRectF(sdl.FRect{
		X: position.X, Y: position.Y,
		W: gameplay.TileSizeF.X, H: gameplay.TileSizeF.Y,
	})
}

func RenderMob(layer *gfx.RendererLayer, mob *gameplay.Mob) error {
	graphic := layer.Data.(*RendererLayerMobData).Mobs[mob.Type]
	rect := GetMobRect(mob)
	texture := graphic.Textures[mob.Direction]
	if mob.Health.Dead && graphic.Dead != nil {
		texture = graphic.Dead
	}
	if err := gfx.RendererInstance.Renderer.Copy(texture.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

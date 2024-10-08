package runechat

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetRunechatCharacterRect(layer *gfx.RendererLayer, runechat *gameplay.Runechat, offset int32) sdl.Rect {
	w, h, _ := gfx.RendererInstance.Fonts.Container[gfx.RendererFontPixeled].SizeUTF8(runechat.Message)
	rect := render.WorldToScreenRectF(sdl.FRect{
		X: runechat.Data.(gameplay.RunechatMobData).Mob.PositionRender.X + 0.5, Y: runechat.Data.(gameplay.RunechatMobData).Mob.PositionRender.Y - 0.35,
		W: float32(w) / 3, H: float32(h) / 3,
	})

	return sdl.Rect{
		X: rect.X - int32(float32(rect.W)/2),
		Y: rect.Y - int32(float32(rect.H)/2) - (offset * rect.H),
		W: rect.W,
		H: rect.H,
	}
}

func RenderRunechatCharacter(layer *gfx.RendererLayer, runechat *gameplay.Runechat, offset int32) error {
	rect := GetRunechatCharacterRect(layer, runechat, offset)
	if runechat.Texture == nil {
		runechat.Texture, _ = sdlutils.CreateUTF8SolidTexture(gfx.RendererInstance.Renderer, gfx.RendererInstance.Fonts.Container[gfx.RendererFontPixeled], sdlutils.White, runechat.Message)
	}
	sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, runechat.Texture, rect)

	return nil
}

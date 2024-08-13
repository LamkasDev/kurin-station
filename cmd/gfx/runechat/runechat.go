package runechat

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinRunechatCharacterRect(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, runechat *gameplay.KurinRunechat, offset int32) sdl.Rect {
	w, h, _ := renderer.Fonts.Container[gfx.KurinRendererFontPixeled].SizeUTF8(runechat.Message)
	rect := render.WorldToScreenRect(renderer, sdl.FRect{
		X: runechat.Data.(gameplay.KurinRunechatCharacterData).Character.PositionRender.X + 0.5, Y: runechat.Data.(gameplay.KurinRunechatCharacterData).Character.PositionRender.Y - 0.35,
		W: float32(w) / 3, H: float32(h) / 3,
	})

	return sdl.Rect{
		X: rect.X - int32(float32(rect.W)/2),
		Y: rect.Y - int32(float32(rect.H)/2) - (offset * rect.H),
		W: rect.W,
		H: rect.H,
	}
}

func RenderKurinRunechatCharacter(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, runechat *gameplay.KurinRunechat, offset int32) error {
	rect := GetKurinRunechatCharacterRect(renderer, layer, runechat, offset)
	if runechat.Texture == nil {
		runechat.Texture, _ = sdlutils.CreateUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Container[gfx.KurinRendererFontPixeled], sdlutils.White, runechat.Message)
	}
	sdlutils.RenderUTF8SolidTextureRect(renderer.Renderer, runechat.Texture, rect)

	return nil
}

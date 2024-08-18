package dialog

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func NewKurinDialogElementLabel(position sdl.Point, text string, icon *sdlutils.TextureWithSize) *KurinDialogElement {
	w, h, _ := gfx.RendererInstance.Fonts.Default.SizeUTF8(text)
	return &KurinDialogElement{
		GetRect: func(dialogRect *sdl.Rect) *sdl.Rect {
			return &sdl.Rect{X: position.X, Y: position.Y, W: int32(w), H: int32(h)}
		},
		Render: func(element *KurinDialogElement, rect *sdl.Rect) {
			if icon != nil {
				sdlutils.RenderTexture(gfx.RendererInstance.Renderer, icon, sdl.Point{X: rect.X, Y: rect.Y}, sdl.FPoint{X: 1, Y: 1})
				rect.X += 24
			}
			sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("dialog.text.%s", text), gfx.RendererInstance.Fonts.Default, sdlutils.White, text, sdl.Point{X: rect.X, Y: rect.Y}, sdl.FPoint{X: 1, Y: 1})
		},
	}
}

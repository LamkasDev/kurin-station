package dialog

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func NewKurinDialogElementButton(position sdl.Point, text string, onClick KurinDialogElementOnClick) *KurinDialogElement {
	w, h, _ := gfx.RendererInstance.Fonts.Default.SizeUTF8(text)
	return &KurinDialogElement{
		GetRect: func(dialogRect *sdl.Rect) *sdl.Rect {
			return &sdl.Rect{X: position.X, Y: position.Y, W: int32(w) + 28, H: int32(h) + 12}
		},
		Render: func(element *KurinDialogElement, rect *sdl.Rect) {
			if element.Hovered {
				sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Blue)
			} else {
				sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Gray)
			}
			gfx.RendererInstance.Renderer.FillRect(rect)
			sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("dialog.button.%s", text), gfx.RendererInstance.Fonts.Default, sdlutils.White, text, sdl.Point{X: rect.X + 14, Y: rect.Y + 6}, sdl.FPoint{X: 1, Y: 1})
		},
		OnClick:   onClick,
		Clickable: true,
	}
}

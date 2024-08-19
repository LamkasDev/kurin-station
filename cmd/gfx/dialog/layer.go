package dialog

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type RendererLayerDialogData struct {
	Dialog    *Dialog
	ItemLayer *gfx.RendererLayer
}

func NewRendererLayerDialog(itemLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerDialog,
		Render: RenderRendererLayerDialog,
		Data: &RendererLayerDialogData{
			Dialog:    nil,
			ItemLayer: itemLayer,
		},
	}
}

func LoadRendererLayerDialog(layer *gfx.RendererLayer) error {
	return nil
}

func RenderRendererLayerDialog(layer *gfx.RendererLayer) error {
	data := layer.Data.(*RendererLayerDialogData)
	if data.Dialog == nil {
		return nil
	}

	dialogSize := data.Dialog.GetSize(gfx.RendererInstance.Context.WindowSize)
	dialogRect := &sdl.Rect{
		X: data.Dialog.Position.X,
		Y: data.Dialog.Position.Y,
		W: dialogSize.X,
		H: dialogSize.Y,
	}

	dialogRect.H += 32
	sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.DarkGray)
	gfx.RendererInstance.Renderer.FillRect(dialogRect)
	sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.LightBlack)
	gfx.RendererInstance.Renderer.DrawRect(dialogRect)

	sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Gray)
	gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: dialogRect.X, Y: dialogRect.Y, W: dialogRect.W, H: 32})
	sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.LightBlack)
	gfx.RendererInstance.Renderer.DrawRect(&sdl.Rect{X: dialogRect.X, Y: dialogRect.Y, W: dialogRect.W, H: 32})

	closeIcon := sdlutils.GetTextureFromContainer(gfx.RendererInstance.IconTextures, gfx.RendererInstance.Renderer, "close")
	dialogIcon := sdlutils.GetTextureFromContainer(gfx.RendererInstance.IconTextures, gfx.RendererInstance.Renderer, data.Dialog.Icon)
	sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, dialogIcon, sdl.Rect{X: dialogRect.X + 8, Y: dialogRect.Y + 8, W: 16, H: 16})
	sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("dialog.%s.title", data.Dialog.Type), gfx.RendererInstance.Fonts.Default, sdlutils.White, data.Dialog.Title, sdl.Point{X: dialogRect.X + 32, Y: dialogRect.Y + 8}, sdl.FPoint{X: 1, Y: 1})
	sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, closeIcon, sdl.Rect{X: dialogRect.X + dialogRect.W - 32 + 8, Y: dialogRect.Y + 8, W: 16, H: 16})
	dialogRect.Y += 32
	dialogRect.H -= 32

	for _, element := range data.Dialog.Elements {
		rect := element.GetRect(dialogRect)
		rect.X += dialogRect.X
		rect.Y += dialogRect.Y
		element.Render(element, rect)
	}

	return nil
}

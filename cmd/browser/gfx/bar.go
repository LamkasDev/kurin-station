package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/context"
	"github.com/LamkasDev/kitsune/cmd/common/constants"
	"github.com/LamkasDev/kitsune/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/net/html/atom"
)

func RenderKitsuneRendererBar(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender) *error {
	sdlutils.SetDrawColor(renderer.Renderer, constants.ApplicationBarColor)
	renderer.Renderer.FillRect(&sdl.Rect{W: renderer.WindowContext.WindowSize.W, H: constants.ApplicationTabHeight})
	sdlutils.SetDrawColor(renderer.Renderer, constants.ApplicationTabColor)
	renderer.Renderer.FillRect(&sdl.Rect{X: 6, Y: 6, W: 384, H: constants.ApplicationTabHeight - 12})

	RenderKitsuneRendererBarIcon(renderer, rcontext)
	if err := RenderKitsuneRendererBarTitle(renderer, rcontext); err != nil {
		return err
	}
	RenderKitsuneRendererBarControls(renderer, rcontext)
	RenderKitsuneRendererScrollbar(renderer, rcontext)

	return nil
}

func RenderKitsuneRendererBarIcon(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender) {
	renderer.Renderer.Copy(rcontext.Icon, nil, &sdl.Rect{X: 12, Y: 10, W: constants.ApplicationTabHeight - 22, H: constants.ApplicationTabHeight - 22})
}

func RenderKitsuneRendererBarTitle(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender) *error {
	titleSurface, titleTexture, _ := sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Elements.Regular[atom.P], rcontext.Title, sdl.Color{R: 255, G: 255, B: 255})
	renderer.Renderer.Copy(titleTexture, nil, &sdl.Rect{X: 44, Y: 12, W: titleSurface.W, H: titleSurface.H})

	return nil
}

func RenderKitsuneRendererBarControls(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender) {
	sdlutils.SetDrawColor(renderer.Renderer, constants.ApplicationTabColor)
	renderer.Renderer.FillRect(&sdl.Rect{X: renderer.WindowContext.WindowSize.W - 40, Y: 6, W: 34, H: constants.ApplicationTabHeight - 12})
	renderer.Renderer.Copy(renderer.Icons.Close, nil, &sdl.Rect{X: renderer.WindowContext.WindowSize.W - 31, Y: 16, W: 16, H: 16})
}

func RenderKitsuneRendererScrollbar(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender) {
	sdlutils.SetDrawColor(renderer.Renderer, constants.ApplicationBarColor)
	// renderer.Renderer.FillRect(&sdl.Rect{X: renderer.WindowContext.WindowSize.W - 12, Y: constants.ApplicationTabHeight, W: 12, H: renderer.WindowContext.WindowSize.H - constants.ApplicationTabHeight})

	sdlutils.SetDrawColor(renderer.Renderer, constants.ApplicationTabColor)
	scrollbarHeight := renderer.WindowContext.WindowSize.H - constants.ApplicationTabHeight
	scrollbarHandleHeight := int32((float32(scrollbarHeight) / float32(renderer.FrameContext.PageSize.H)) * float32(scrollbarHeight))
	scrollbarProgress := (rcontext.ScrollPosition.Y / GetScrollbarLimit(renderer))
	scrollY := int32(scrollbarProgress * float32(scrollbarHeight-scrollbarHandleHeight))
	renderer.Renderer.FillRect(&sdl.Rect{X: renderer.WindowContext.WindowSize.W - 12, Y: constants.ApplicationTabHeight + scrollY, W: 12, H: scrollbarHandleHeight})
}

func GetScrollbarLimit(renderer *KitsuneRenderer) float32 {
	return float32(-renderer.FrameContext.PageSize.H + renderer.WindowContext.WindowSize.H - constants.ApplicationTabHeight)
}

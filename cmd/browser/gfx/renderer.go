package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/context"
	"github.com/LamkasDev/kitsune/cmd/common/constants"
	"github.com/LamkasDev/kitsune/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type KitsuneRenderer struct {
	Window        *sdl.Window
	Renderer      *sdl.Renderer
	Icons         *KitsuneRendererIcons
	Fonts         *KitsuneRendererFonts
	WindowContext *KitsuneRendererWindowContext
	FrameContext  *KitsuneRendererFrameContext
}

type KitsuneRendererIcons struct {
	Icon  *sdl.Texture
	Close *sdl.Texture
}

func NewKitsuneRenderer() (*KitsuneRenderer, *error) {
	renderer := &KitsuneRenderer{
		Icons:         &KitsuneRendererIcons{},
		WindowContext: NewKitsuneWindowContext(),
	}
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, &err
	}
	if err = ttf.Init(); err != nil {
		return nil, &err
	}

	if renderer.Window, err = sdl.CreateWindow(constants.ApplicationName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE); err != nil {
		return nil, &err
	}
	renderer.WindowContext.WindowSize = sdl.Rect{
		W: 800,
		H: 600,
	}

	if renderer.Renderer, err = sdl.CreateRenderer(renderer.Window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return nil, &err
	}
	renderer.Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)

	icon, err := img.Load(constants.ApplicationIcon)
	if err != nil {
		return nil, &err
	}
	renderer.Icons.Icon, err = renderer.Renderer.CreateTextureFromSurface(icon)
	if err != nil {
		return nil, &err
	}
	renderer.Window.SetIcon(icon)

	iconClose, err := img.Load(constants.ApplicationIconClose)
	if err != nil {
		return nil, &err
	}
	renderer.Icons.Close, err = renderer.Renderer.CreateTextureFromSurface(iconClose)
	if err != nil {
		return nil, &err
	}

	var fontErr *error
	renderer.Fonts, fontErr = NewKitsuneRendererFonts()
	if fontErr != nil {
		return nil, fontErr
	}

	return renderer, nil
}

func ClearKitsuneRenderer(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender) {
	renderer.Renderer.Clear()
	sdlutils.SetDrawColor(renderer.Renderer, constants.ApplicationBackgroundColor)
	renderer.Renderer.FillRect(nil)
	renderer.FrameContext = NewKitsuneRendererContext()
}

func RenderKitsuneRenderer(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender) *error {
	renderer.FrameContext.LayoutPosition.Y += constants.ApplicationTabHeight + int32(rcontext.ScrollPosition.Y)
	if err := RenderKitsuneRendererElementTree(renderer, rcontext, rcontext.Document); err != nil {
		return err
	}
	if err := RenderKitsuneRendererBar(renderer, rcontext); err != nil {
		return err
	}
	if renderer.FrameContext.InspectedElement != nil {
		elementRect := renderer.FrameContext.InspectedElement.CachedRect
		marginRect := GetKitsuneElementMarginRect(renderer, renderer.FrameContext.InspectedElement)

		sdlutils.SetDrawColor(renderer.Renderer, constants.MarginColor)
		renderer.Renderer.FillRect(&marginRect)
		sdlutils.SetDrawColor(renderer.Renderer, constants.BoundingBoxFillColor)
		renderer.Renderer.FillRect(&elementRect)
		sdlutils.SetDrawColor(renderer.Renderer, constants.BoundingBoxBorderColor)
		renderer.Renderer.DrawRect(&elementRect)
	}

	return nil
}

func PresentKitsuneRenderer(renderer *KitsuneRenderer, rcontext *context.KitsuneContextRender) {
	renderer.Renderer.Present()
}

func FreeKitsuneRenderer(renderer *KitsuneRenderer) {
	sdl.Quit()
	renderer.Window.Destroy()
	renderer.Renderer.Destroy()
	renderer.Icons.Icon.Destroy()
	renderer.Icons.Close.Destroy()
	FreeKitsuneRendererFonts(renderer.Fonts)
}

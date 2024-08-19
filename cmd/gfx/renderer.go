package gfx

import (
	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var RendererInstance *Renderer

type Renderer struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Fonts    RendererFonts
	Context  RendererContext
	Layers   []*RendererLayer

	IconTextures *sdlutils.TextureContainer
}

func InitializeRenderer() error {
	RendererInstance = &Renderer{
		Context:      NewRendererContext(),
		Layers:       []*RendererLayer{},
		IconTextures: sdlutils.NewTextureContainer("icons"),
	}
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return err
	}
	if err = ttf.Init(); err != nil {
		return err
	}

	if RendererInstance.Window, err = sdl.CreateWindow(constants.ApplicationName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE); err != nil {
		return err
	}
	RendererInstance.Context.WindowSize = sdl.Point{X: 800, Y: 600}

	if RendererInstance.Renderer, err = sdl.CreateRenderer(RendererInstance.Window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return err
	}
	if err := RendererInstance.Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE); err != nil {
		return err
	}

	icon, _, iconErr := sdlutils.LoadTextureRaw(RendererInstance.Renderer, constants.ApplicationIcon)
	if iconErr != nil {
		return err
	}
	RendererInstance.Window.SetIcon(icon)

	var fontErr error
	RendererInstance.Fonts, fontErr = NewRendererFonts()
	if fontErr != nil {
		return fontErr
	}

	return nil
}

func LoadRenderer() error {
	for _, layer := range RendererInstance.Layers {
		if err := layer.Load(layer); err != nil {
			return err
		}
	}

	return nil
}

func ClearRenderer() error {
	if err := RendererInstance.Renderer.Clear(); err != nil {
		return err
	}
	if err := sdlutils.SetDrawColor(RendererInstance.Renderer, constants.ApplicationBackgroundColor); err != nil {
		return err
	}
	if err := RendererInstance.Renderer.FillRect(nil); err != nil {
		return err
	}

	return nil
}

func RenderRenderer() error {
	for _, layer := range RendererInstance.Layers {
		if err := layer.Render(layer); err != nil {
			return err
		}
	}

	return nil
}

func PresentRenderer() {
	RendererInstance.Renderer.Present()
	RendererInstance.Context.Frame++
}

// TODO: make layers free themselves.
func FreeRenderer() error {
	if err := RendererInstance.Renderer.Destroy(); err != nil {
		return err
	}
	FreeRendererFonts(&RendererInstance.Fonts)
	if err := RendererInstance.Window.Destroy(); err != nil {
		return err
	}

	return nil
}

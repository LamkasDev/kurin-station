package gfx

import (
	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type KurinRenderer struct {
	Window        *sdl.Window
	Renderer      *sdl.Renderer
	Icons         KurinRendererIcons
	Fonts         KurinRendererFonts
	Context KurinRendererContext
	Layers        []*KurinRendererLayer
}

type KurinRendererIcons struct {
	Icon *sdl.Texture
}

func NewKurinRenderer() (*KurinRenderer, *error) {
	renderer := &KurinRenderer{
		Icons:         KurinRendererIcons{},
		Context: NewKurinRendererContext(),
		Layers:        []*KurinRendererLayer{},
	}
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return renderer, &err
	}
	if err = ttf.Init(); err != nil {
		return renderer, &err
	}

	if renderer.Window, err = sdl.CreateWindow(constants.ApplicationName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE); err != nil {
		return renderer, &err
	}
	renderer.Context.WindowSize = sdl.Point{X: 800, Y: 600}

	if renderer.Renderer, err = sdl.CreateRenderer(renderer.Window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return renderer, &err
	}
	if err := renderer.Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE); err != nil {
		return renderer, &err
	}

	icon, _, iconErr := sdlutils.LoadTextureRaw(renderer.Renderer, constants.ApplicationIcon)
	if iconErr != nil {
		return renderer, &err
	}
	renderer.Window.SetIcon(icon)

	var fontErr *error
	renderer.Fonts, fontErr = NewKurinRendererFonts()
	if fontErr != nil {
		return renderer, fontErr
	}

	return renderer, nil
}

func LoadKurinRenderer(renderer *KurinRenderer) *error {
	for _, layer := range renderer.Layers {
		if err := layer.Load(renderer, layer); err != nil {
			return err
		}
	}

	return nil
}

func ClearKurinRenderer(renderer *KurinRenderer) *error {
	if err := renderer.Renderer.Clear(); err != nil {
		return &err
	}
	if err := sdlutils.SetDrawColor(renderer.Renderer, constants.ApplicationBackgroundColor); err != nil {
		return &err
	}
	if err := renderer.Renderer.FillRect(nil); err != nil {
		return &err
	}

	return nil
}

func RenderKurinRenderer(renderer *KurinRenderer, game *gameplay.KurinGame) *error {
	for _, layer := range renderer.Layers {
		if err := layer.Render(renderer, layer, game); err != nil {
			return err
		}
	}

	return nil
}

func PresentKurinRenderer(renderer *KurinRenderer) {
	renderer.Renderer.Present()
	renderer.Context.Frame++
}

// TODO: make layers free themselves.
func FreeKurinRenderer(renderer *KurinRenderer) *error {
	if err := renderer.Renderer.Destroy(); err != nil {
		return &err
	}
	FreeKurinRendererFonts(&renderer.Fonts)
	if err := renderer.Window.Destroy(); err != nil {
		return &err
	}

	return nil
}

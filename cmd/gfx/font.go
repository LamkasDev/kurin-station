package gfx

import (
	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/veandco/go-sdl2/ttf"
)

type RendererFonts struct {
	Container    map[string]*ttf.Font
	Default      *ttf.Font
	DefaultSmall *ttf.Font
}

const (
	RendererFontDefault      = "default"
	RendererFontDefaultSmall = "default.small"
	RendererFontPixeled      = "pixeled"
	RendererFontOutline      = "outline"
)

func NewRendererFonts() (RendererFonts, error) {
	fonts := RendererFonts{
		Container: map[string]*ttf.Font{},
	}
	var err error
	if fonts.Container[RendererFontDefault], err = ttf.OpenFont(constants.ApplicationFontDefault, 14); err != nil {
		return fonts, err
	}
	if fonts.Container[RendererFontDefaultSmall], err = ttf.OpenFont(constants.ApplicationFontDefault, 10); err != nil {
		return fonts, err
	}
	if fonts.Container[RendererFontPixeled], err = ttf.OpenFont(constants.ApplicationFontPixeled, 24); err != nil {
		return fonts, err
	}
	if fonts.Container[RendererFontOutline], err = ttf.OpenFont(constants.ApplicationFontOutline, 24); err != nil {
		return fonts, err
	}
	fonts.Default = fonts.Container[RendererFontDefault]
	fonts.DefaultSmall = fonts.Container[RendererFontDefaultSmall]

	return fonts, nil
}

func FreeRendererFonts(fonts *RendererFonts) {
	for _, font := range fonts.Container {
		font.Close()
	}
	ttf.Quit()
}

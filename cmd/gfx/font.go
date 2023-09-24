package gfx

import (
	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/veandco/go-sdl2/ttf"
)

type KurinRendererFonts struct {
	Container map[string]*ttf.Font
}

const KurinRendererFontDefault = "default"
const KurinRendererFontPixeled = "pixeled"
const KurinRendererFontOutline = "outline"

func NewKurinRendererFonts() (KurinRendererFonts, *error) {
	fonts := KurinRendererFonts{
		Container: map[string]*ttf.Font{},
	}
	var err error
	if fonts.Container[KurinRendererFontDefault], err = ttf.OpenFont(constants.ApplicationFontDefault, 14); err != nil {
		return fonts, &err
	}
	if fonts.Container[KurinRendererFontPixeled], err = ttf.OpenFont(constants.ApplicationFontPixeled, 24); err != nil {
		return fonts, &err
	}
	if fonts.Container[KurinRendererFontOutline], err = ttf.OpenFont(constants.ApplicationFontOutline, 24); err != nil {
		return fonts, &err
	}

	return fonts, nil
}

func FreeKurinRendererFonts(fonts *KurinRendererFonts) {
	for _, font := range fonts.Container {
		font.Close()
	}
	ttf.Quit()
}

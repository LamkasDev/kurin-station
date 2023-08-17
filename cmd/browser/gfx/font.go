package gfx

import (
	"github.com/LamkasDev/kitsune/cmd/common/constants"
	"github.com/veandco/go-sdl2/ttf"
	"golang.org/x/net/html/atom"
)

type KitsuneRendererFonts struct {
	Elements *KitsuneRendererFontsContainer
}

type KitsuneRendererFontsContainer struct {
	Regular map[atom.Atom]*ttf.Font
	Bold    map[atom.Atom]*ttf.Font
}

func NewKitsuneRendererFonts() (*KitsuneRendererFonts, *error) {
	fonts := &KitsuneRendererFonts{}
	var err *error
	if fonts.Elements, err = NewKitsuneRendererFontsContainer(); err != nil {
		return nil, err
	}

	return fonts, nil
}

func NewKitsuneRendererFontsContainer() (*KitsuneRendererFontsContainer, *error) {
	container := &KitsuneRendererFontsContainer{}
	var err *error
	if container.Regular, err = NewKitsuneRendererFontsMap(constants.ApplicationFontRegular, constants.ApplicationFontBold); err != nil {
		return nil, err
	}
	if container.Bold, err = NewKitsuneRendererFontsMap(constants.ApplicationFontBold, constants.ApplicationFontBold); err != nil {
		return nil, err
	}

	return container, nil
}

func NewKitsuneRendererFontsMap(regularPath string, headingsPath string) (map[atom.Atom]*ttf.Font, *error) {
	containerMap := map[atom.Atom]*ttf.Font{}
	var err error
	if containerMap[atom.Li], err = ttf.OpenFont(regularPath, 18); err != nil {
		return nil, &err
	}
	if containerMap[atom.H1], err = ttf.OpenFont(headingsPath, 38); err != nil {
		return nil, &err
	}
	if containerMap[atom.H2], err = ttf.OpenFont(headingsPath, 28); err != nil {
		return nil, &err
	}
	if containerMap[atom.H3], err = ttf.OpenFont(headingsPath, 22); err != nil {
		return nil, &err
	}
	if containerMap[atom.H4], err = ttf.OpenFont(headingsPath, 18); err != nil {
		return nil, &err
	}
	if containerMap[atom.H5], err = ttf.OpenFont(headingsPath, 16); err != nil {
		return nil, &err
	}
	if containerMap[atom.H6], err = ttf.OpenFont(headingsPath, 13); err != nil {
		return nil, &err
	}
	if containerMap[atom.P], err = ttf.OpenFont(regularPath, 18); err != nil {
		return nil, &err
	}
	if containerMap[atom.Aside], err = ttf.OpenFont(regularPath, 18); err != nil {
		return nil, &err
	}
	if containerMap[atom.A], err = ttf.OpenFont(regularPath, 18); err != nil {
		return nil, &err
	}

	return containerMap, nil
}

func FreeKitsuneRendererFonts(fonts *KitsuneRendererFonts) {
	for _, font := range fonts.Elements.Regular {
		font.Close()
	}
	for _, font := range fonts.Elements.Bold {
		font.Close()
	}
	ttf.Quit()
}

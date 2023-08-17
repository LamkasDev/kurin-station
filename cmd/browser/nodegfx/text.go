package nodegfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/gfx"
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"github.com/LamkasDev/kitsune/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func NewKitsuneElementTextData(renderer *gfx.KitsuneRenderer, htmlNode *html.Node) (*node.KitsuneElementTextData, *error) {
	dataText := htmlNode.FirstChild.Data
	switch htmlNode.DataAtom {
	case atom.Li:
		dataText = "- " + dataText
	}

	textColor := sdl.Color{R: 0, G: 0, B: 0}
	switch htmlNode.DataAtom {
	case atom.A:
		textColor = sdl.Color{R: 0, G: 0, B: 255}
	}

	textSurface, textTexture, _ := sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Elements.Regular[htmlNode.DataAtom], dataText, textColor)
	return &node.KitsuneElementTextData{
		Text: textTexture,
		Size: sdl.Rect{
			W: textSurface.W,
			H: textSurface.H,
		},
	}, nil
}

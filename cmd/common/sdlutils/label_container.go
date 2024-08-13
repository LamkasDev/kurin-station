package sdlutils

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var LabelContainer = map[string]*RendererLabel{}

func RenderLabel(renderer *sdl.Renderer, id string, font *ttf.Font, color sdl.Color, text string, position sdl.Point, scale sdl.FPoint) (error, *sdl.Rect) {
	label, ok := LabelContainer[id]
	if !ok {
		label = NewLabel(renderer, font, color)
		LabelContainer[id] = label
	}

	return RenderLabelRaw(renderer, label, text, position, scale)
}

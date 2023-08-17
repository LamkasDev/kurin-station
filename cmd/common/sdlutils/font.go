package sdlutils

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func RenderUTF8SolidTexture(renderer *sdl.Renderer, font *ttf.Font, text string, color sdl.Color) (*sdl.Surface, *sdl.Texture, *error) {
	textSurface, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		return nil, nil, &err
	}
	textTexture, err := renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return nil, nil, &err
	}

	return textSurface, textTexture, nil
}

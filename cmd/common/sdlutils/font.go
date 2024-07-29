package sdlutils

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func CreateUTF8SolidTexture(renderer *sdl.Renderer, font *ttf.Font, text string, color sdl.Color) (*sdl.Surface, *sdl.Texture, *error) {
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

func RenderUTF8SolidTexture(renderer *sdl.Renderer, font *ttf.Font, text string, color sdl.Color, position sdl.Point, scale sdl.FPoint) (*error, *sdl.Rect) {
	textSurface, textTexture, err := CreateUTF8SolidTexture(renderer, font, text, color)
	if err != nil {
		return err, nil
	}
	rect := &sdl.Rect{
		X: position.X,
		Y: position.Y,
		W: int32(float32(textSurface.W) * scale.X),
		H: int32(float32(textSurface.H) * scale.Y),
	}
	if err := renderer.Copy(textTexture, nil, rect); err != nil {
		return &err, nil
	}

	return nil, rect
}

func RenderUTF8SolidTextureRect(renderer *sdl.Renderer, font *ttf.Font, text string, color sdl.Color, rect sdl.Rect) *error {
	_, textTexture, err := CreateUTF8SolidTexture(renderer, font, text, color)
	if err != nil {
		return err
	}
	if err := renderer.Copy(textTexture, nil, &rect); err != nil {
		return &err
	}

	return nil
}

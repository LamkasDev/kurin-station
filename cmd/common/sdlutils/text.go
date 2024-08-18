package sdlutils

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func CreateAndRenderUTF8SolidTexture(renderer *sdl.Renderer, font *ttf.Font, color sdl.Color, text string, position sdl.Point, scale sdl.FPoint) (error, *sdl.Rect) {
	texture, err := CreateUTF8SolidTexture(renderer, font, color, text)
	if err != nil {
		return err, nil
	}
	err, rect := RenderUTF8SolidTexture(renderer, texture, position, scale)
	if err != nil {
		return err, nil
	}
	texture.Texture.Destroy()

	return nil, rect
}

func CreateUTF8SolidTexture(renderer *sdl.Renderer, font *ttf.Font, color sdl.Color, text string) (*TextureWithSize, error) {
	textSurface, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		return nil, err
	}
	textTexture, err := renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return nil, err
	}
	texture := TextureWithSize{
		Texture: textTexture,
		Size: sdl.Rect{
			W: textSurface.W,
			H: textSurface.H,
		},
	}

	return &texture, nil
}

func RenderUTF8SolidTexture(renderer *sdl.Renderer, texture *TextureWithSize, position sdl.Point, scale sdl.FPoint) (error, *sdl.Rect) {
	rect := &sdl.Rect{
		X: position.X,
		Y: position.Y,
		W: int32(float32(texture.Size.W) * scale.X),
		H: int32(float32(texture.Size.H) * scale.Y),
	}
	if err := renderer.Copy(texture.Texture, nil, rect); err != nil {
		return err, nil
	}

	return nil, rect
}

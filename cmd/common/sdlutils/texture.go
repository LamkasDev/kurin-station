package sdlutils

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureWithSize struct {
	Texture *sdl.Texture
	Size    sdl.Rect
}

type TextureWithSizeAndSurface struct {
	Base    TextureWithSize
	Surface *sdl.Surface
}

func LoadTextureRaw(renderer *sdl.Renderer, path string) (*sdl.Surface, *sdl.Texture, error) {
	icon, err := img.Load(path)
	if err != nil {
		return nil, nil, err
	}
	iconTexture, err := renderer.CreateTextureFromSurface(icon)
	if err != nil {
		return nil, nil, err
	}

	return icon, iconTexture, nil
}

func LoadTexture(renderer *sdl.Renderer, path string) (TextureWithSize, error) {
	surface, texture, err := LoadTextureRaw(renderer, path)
	if err != nil {
		return TextureWithSize{}, err
	}

	return TextureWithSize{
		Texture: texture,
		Size:    sdl.Rect{W: surface.W, H: surface.H},
	}, nil
}

func LoadTextureWithSurface(renderer *sdl.Renderer, path string) (TextureWithSizeAndSurface, error) {
	surface, texture, err := LoadTextureRaw(renderer, path)
	if err != nil {
		return TextureWithSizeAndSurface{}, err
	}

	return TextureWithSizeAndSurface{
		Base: TextureWithSize{
			Texture: texture,
			Size:    sdl.Rect{W: surface.W, H: surface.H},
		},
		Surface: surface,
	}, nil
}

func GetTextureRect(renderer *sdl.Renderer, texture TextureWithSize, position sdl.Point, scale sdl.FPoint) sdl.Rect {
	return sdl.Rect{
		X: position.X,
		Y: position.Y,
		W: int32(float32(texture.Size.W) * scale.X),
		H: int32(float32(texture.Size.H) * scale.Y),
	}
}

func RenderTexture(renderer *sdl.Renderer, texture TextureWithSize, position sdl.Point, scale sdl.FPoint) error {
	rect := GetTextureRect(renderer, texture, position, scale)
	if err := renderer.Copy(texture.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

func GetPixelAt(texture TextureWithSizeAndSurface, position sdl.Point) *sdl.Color {
	pixels := texture.Surface.Pixels()
	i := int(position.Y*texture.Surface.Pitch + position.X*int32(texture.Surface.BytesPerPixel()))
	if len(pixels) <= i {
		return nil
	}
	pixel := uint32(texture.Surface.Pixels()[i])
	r, g, b, a := sdl.GetRGBA(pixel, texture.Surface.Format)

	return &sdl.Color{R: r, G: g, B: b, A: a}
}

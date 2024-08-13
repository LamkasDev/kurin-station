package sdlutils

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type RendererLabel struct {
	Text    string
	Font *ttf.Font
	Color sdl.Color
	Texture *TextureWithSize
}

func NewLabel(renderer *sdl.Renderer, font *ttf.Font, color sdl.Color) *RendererLabel {
	return &RendererLabel{
		Text:    "",
		Font: font,
		Color: color,
		Texture: nil,
	}
}

func RenderLabelRaw(renderer *sdl.Renderer, label *RendererLabel, text string, position sdl.Point, scale sdl.FPoint) (error, *sdl.Rect) {
	var err error
	if label.Text != text {
		label.Text = text
		if label.Texture != nil {
			label.Texture.Texture.Destroy()
		}
		if label.Texture, err = CreateUTF8SolidTexture(renderer, label.Font, label.Color, text); err != nil {
			return err, nil
		}
	}
	if label.Texture != nil {
		return RenderUTF8SolidTexture(renderer, label.Texture, position, scale)
	}
	
	return nil, nil
}

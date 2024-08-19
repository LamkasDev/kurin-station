package background

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type RendererLayerBackgroundData struct {
	Position         sdl.FPoint
	TextureContainer *sdlutils.TextureContainer
}

func NewRendererLayerBackground() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerBackground,
		Render: RenderRendererLayerBackground,
		Data: &RendererLayerBackgroundData{
			TextureContainer: sdlutils.NewTextureContainer("bgs"),
		},
	}
}

func LoadRendererLayerBackground(layer *gfx.RendererLayer) error {
	return nil
}

func RenderRendererLayerBackground(layer *gfx.RendererLayer) error {
	data := layer.Data.(*RendererLayerBackgroundData)
	position := sdlutils.FPointToPoint(data.Position)
	for range 6 {
		for range 3 {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, sdlutils.GetTextureFromContainer(data.TextureContainer, gfx.RendererInstance.Renderer, "bg"), position, sdl.FPoint{X: 1, Y: 1})
			position.Y += 480
		}
		position.X += 480
		position.Y = 0
	}
	data.Position.X -= 0.25
	if data.Position.X == -480 {
		data.Position.X = 0
	}

	return nil
}

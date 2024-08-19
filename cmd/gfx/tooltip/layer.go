package tooltip

import (
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerTooltipData struct{}

func NewRendererLayerTooltip() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerTooltip,
		Render: RenderRendererLayerTooltip,
		Data:   &RendererLayerTooltipData{},
	}
}

func LoadRendererLayerTooltip(layer *gfx.RendererLayer) error {
	return nil
}

func RenderRendererLayerTooltip(layer *gfx.RendererLayer) error {
	return nil
}

package tooltip

import (
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerTooltipData struct{}

func NewKurinRendererLayerTooltip() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerTooltip,
		Render: RenderKurinRendererLayerTooltip,
		Data:   &KurinRendererLayerTooltipData{},
	}
}

func LoadKurinRendererLayerTooltip(layer *gfx.RendererLayer) error {
	return nil
}

func RenderKurinRendererLayerTooltip(layer *gfx.RendererLayer) error {
	return nil
}

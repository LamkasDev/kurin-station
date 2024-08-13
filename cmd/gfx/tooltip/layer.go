package tooltip

import (
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerTooltipData struct{}

func NewKurinRendererLayerTooltip() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerTooltip,
		Render: RenderKurinRendererLayerTooltip,
		Data:   KurinRendererLayerTooltipData{},
	}
}

func LoadKurinRendererLayerTooltip(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	return nil
}

func RenderKurinRendererLayerTooltip(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	return nil
}

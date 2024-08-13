package context

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const KurinRendererLayerContextDataItemWidth = 164
const KurinRendererLayerContextDataItemHeight = 36

type KurinRendererLayerContextData struct {
	Position *sdl.Point
	Items []KurinRendererLayerContextDataItem
	HoveredItem int
}

type KurinRendererLayerContextDataItem struct {
	Text string
	Disabled bool
	OnClick func()
}

func NewKurinRendererLayerContext() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerContext,
		Render: RenderKurinRendererLayerContext,
		Data:   KurinRendererLayerContextData{
			Position: nil,
			Items: []KurinRendererLayerContextDataItem{},
			HoveredItem: -1,
		},
	}
}

func LoadKurinRendererLayerContext(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	return nil
}

func RenderKurinRendererLayerContext(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	data := layer.Data.(KurinRendererLayerContextData)
	if data.Position == nil {
		return nil
	}
	for i, item := range data.Items {
		y := data.Position.Y + int32(i) * KurinRendererLayerContextDataItemHeight
		rect := sdl.Rect{X: data.Position.X, Y: y, W: KurinRendererLayerContextDataItemWidth, H: KurinRendererLayerContextDataItemHeight}
		renderer.Renderer.SetDrawColor(255, 255, 255, 0)
		if err := renderer.Renderer.FillRect(&rect); err != nil {
			return err
		}
		renderer.Renderer.SetDrawColor(233, 233, 233, 0)
		if err := renderer.Renderer.DrawRect(&rect); err != nil {
			return err
		}
		_, text := sdlutils.RenderLabel(renderer.Renderer, fmt.Sprintf("context.%d", i), renderer.Fonts.Default, sdl.Color{R: 0, G: 0, B: 0}, item.Text, sdl.Point{X: data.Position.X + 12, Y: y + 10}, sdl.FPoint{X: 1, Y: 1})
		if data.HoveredItem == i {
			renderer.Renderer.SetDrawColor(0, 0, 0, 0)
			renderer.Renderer.DrawLine(text.X, text.Y + text.H + 2, text.X + text.W, text.Y + text.H + 2)
		}
	}

	return nil
}

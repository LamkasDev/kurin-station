package context

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	RendererLayerContextDataItemWidth  = 164
	RendererLayerContextDataItemHeight = 36
)

type RendererLayerContextData struct {
	Position    *sdl.Point
	Items       []RendererLayerContextDataItem
	HoveredItem int
}

type RendererLayerContextDataItem struct {
	Text     string
	Disabled bool
	OnClick  func()
}

func NewRendererLayerContext() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerContext,
		Render: RenderRendererLayerContext,
		Data: &RendererLayerContextData{
			Position:    nil,
			Items:       []RendererLayerContextDataItem{},
			HoveredItem: -1,
		},
	}
}

func LoadRendererLayerContext(layer *gfx.RendererLayer) error {
	return nil
}

func RenderRendererLayerContext(layer *gfx.RendererLayer) error {
	data := layer.Data.(*RendererLayerContextData)
	if data.Position == nil {
		return nil
	}
	for i, item := range data.Items {
		y := data.Position.Y + int32(i)*RendererLayerContextDataItemHeight
		rect := sdl.Rect{X: data.Position.X, Y: y, W: RendererLayerContextDataItemWidth, H: RendererLayerContextDataItemHeight}
		gfx.RendererInstance.Renderer.SetDrawColor(255, 255, 255, 0)
		if err := gfx.RendererInstance.Renderer.FillRect(&rect); err != nil {
			return err
		}
		gfx.RendererInstance.Renderer.SetDrawColor(233, 233, 233, 0)
		if err := gfx.RendererInstance.Renderer.DrawRect(&rect); err != nil {
			return err
		}
		_, text := sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("context.%d", i), gfx.RendererInstance.Fonts.Default, sdl.Color{R: 0, G: 0, B: 0}, item.Text, sdl.Point{X: data.Position.X + 12, Y: y + 10}, sdl.FPoint{X: 1, Y: 1})
		if data.HoveredItem == i {
			gfx.RendererInstance.Renderer.SetDrawColor(0, 0, 0, 0)
			gfx.RendererInstance.Renderer.DrawLine(text.X, text.Y+text.H+2, text.X+text.W, text.Y+text.H+2)
		}
	}

	return nil
}

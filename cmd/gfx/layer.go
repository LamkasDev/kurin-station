package gfx

type RendererLayer struct {
	Load   RendererLayerLoad
	Render RendererLayerRender
	Data   interface{}
}

type RendererLayerLoad func(layer *RendererLayer) error

type RendererLayerRender func(layer *RendererLayer) error

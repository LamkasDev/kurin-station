package gfx

type RendererLayer struct {
	Load   KurinRendererLayerLoad
	Render KurinRendererLayerRender
	Data   interface{}
}

type KurinRendererLayerLoad func(layer *RendererLayer) error

type KurinRendererLayerRender func(layer *RendererLayer) error

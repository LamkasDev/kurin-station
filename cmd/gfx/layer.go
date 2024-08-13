package gfx

type KurinRendererLayer struct {
	Load   KurinRendererLayerLoad
	Render KurinRendererLayerRender
	Data   interface{}
}

type KurinRendererLayerLoad func(renderer *KurinRenderer, layer *KurinRendererLayer) error

type KurinRendererLayerRender func(renderer *KurinRenderer, layer *KurinRendererLayer) error

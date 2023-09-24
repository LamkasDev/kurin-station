package structure

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinObjectRect(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, tile *gameplay.KurinTile, obj *gameplay.KurinObject) sdl.FRect {
	graphic := layer.Data.(KurinRendererLayerObjectData).Structures[obj.Type]
	texture := graphic.Textures[obj.Direction]
	return sdl.FRect{
		X: float32(tile.Position.Base.X), Y: float32(tile.Position.Base.Y),
		W: float32(texture.Size.W), H: float32(texture.Size.H),
	}
}

func RenderKurinObject(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, tile *gameplay.KurinTile, obj *gameplay.KurinObject) *error {
	graphic := layer.Data.(KurinRendererLayerObjectData).Structures[obj.Type]
	wrect := render.WorldToScreenRect(renderer, GetKurinObjectRect(renderer, layer, tile, obj))
	if err := renderer.Renderer.Copy(graphic.Textures[obj.Direction].Texture, nil, &wrect); err != nil {
		return &err
	}

	return nil
}

func RenderKurinObjectBlueprint(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, tile *gameplay.KurinTile, obj *gameplay.KurinObject, color sdl.Color) *error {
	graphic := layer.Data.(KurinRendererLayerObjectData).Structures[obj.Type]
	wrect := render.WorldToScreenRect(renderer, GetKurinObjectRect(renderer, layer, tile, obj))
	graphic.Blueprint.Texture.SetColorMod(color.R, color.G, color.B)
	if err := renderer.Renderer.Copy(graphic.Blueprint.Texture, nil, &wrect); err != nil {
		return &err
	}

	return nil
}

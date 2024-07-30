package item

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinItemRect(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame, item *gameplay.KurinItem) sdl.Rect {
	graphic := layer.Data.(KurinRendererLayerItemData).Items[item.Type]
	return render.WorldToScreenRect(renderer, sdl.FRect{
		X: float32(item.Transform.Position.Base.X) - 0.5, Y: float32(item.Transform.Position.Base.Y) - 0.5,
		W: float32(graphic.Textures[0].Base.Size.W), H: float32(graphic.Textures[0].Base.Size.H),
	})
}

func RenderKurinItem(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame, item *gameplay.KurinItem) *error {
	graphic := layer.Data.(KurinRendererLayerItemData).Items[item.Type]
	rect := GetKurinItemRect(renderer, layer, game, item)

	if game.HoveredItem == item && graphic.Outline != nil {
		if err := renderer.Renderer.CopyEx(graphic.Outline.Texture, nil, &rect, item.Transform.Rotation, nil, sdl.RendererFlip(0)); err != nil {
			return &err
		}
	}
	for _, i := range item.GetTextures(item, game) {
		if err := renderer.Renderer.CopyEx(graphic.Textures[i].Base.Texture, nil, &rect, item.Transform.Rotation, nil, sdl.RendererFlip(0)); err != nil {
			return &err
		}
	}

	return nil
}

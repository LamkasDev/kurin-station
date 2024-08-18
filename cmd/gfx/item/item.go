package item

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinItemRect(layer *gfx.RendererLayer, item *gameplay.KurinItem) sdl.Rect {
	graphic := layer.Data.(*KurinRendererLayerItemData).Items[item.Type]
	texture := graphic.Textures[0]
	return render.WorldToScreenRect(sdl.FRect{
		X: float32(item.Transform.Position.Base.X) - 0.5, Y: float32(item.Transform.Position.Base.Y) - 0.5,
		W: float32(texture.Base.Size.W), H: float32(texture.Base.Size.H),
	})
}

func RenderKurinItem(layer *gfx.RendererLayer, item *gameplay.KurinItem) error {
	graphic := layer.Data.(*KurinRendererLayerItemData).Items[item.Type]
	rect := GetKurinItemRect(layer, item)
	if gameplay.GameInstance.HoveredItem == item && graphic.Outline != nil {
		if err := gfx.RendererInstance.Renderer.CopyEx(graphic.Outline.Texture, nil, &rect, item.Transform.Rotation, nil, sdl.FLIP_NONE); err != nil {
			return err
		}
	}
	for _, i := range item.GetTextures(item) {
		if err := gfx.RendererInstance.Renderer.CopyEx(graphic.Textures[i].Base.Texture, nil, &rect, item.Transform.Rotation, nil, sdl.FLIP_NONE); err != nil {
			return err
		}
	}

	return nil
}

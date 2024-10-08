package item

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetItemRect(layer *gfx.RendererLayer, item *gameplay.Item) sdl.Rect {
	graphic := layer.Data.(*RendererLayerItemData).Items[item.Type]
	texture := graphic.Textures[0]
	return render.WorldToScreenRectF(sdl.FRect{
		X: float32(item.Transform.Position.Base.X) - 0.5, Y: float32(item.Transform.Position.Base.Y) - 0.5,
		W: float32(texture.Base.Size.W), H: float32(texture.Base.Size.H),
	})
}

func RenderItem(layer *gfx.RendererLayer, item *gameplay.Item) error {
	graphic := layer.Data.(*RendererLayerItemData).Items[item.Type]
	rect := GetItemRect(layer, item)
	if gameplay.GameInstance.HoveredItem == item && graphic.Outline != nil {
		if err := gfx.RendererInstance.Renderer.CopyEx(graphic.Outline.Texture, nil, &rect, item.Transform.Rotation, nil, sdl.FLIP_NONE); err != nil {
			return err
		}
	}
	for _, i := range item.Template.GetTextures(item) {
		if err := gfx.RendererInstance.Renderer.CopyEx(graphic.Textures[i].Base.Texture, nil, &rect, item.Transform.Rotation, nil, sdl.FLIP_NONE); err != nil {
			return err
		}
	}

	return nil
}

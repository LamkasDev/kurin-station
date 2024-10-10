package item

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetItemRect(layer *gfx.RendererLayer, item *gameplay.Item, texture *sdlutils.TextureWithSize) sdl.Rect {
	return render.WorldToScreenRectF(sdl.FRect{
		X: float32(item.Transform.Position.Base.X) - 0.5, Y: float32(item.Transform.Position.Base.Y) - 0.5,
		W: float32(texture.Size.W), H: float32(texture.Size.H),
	})
}

func RenderItem(layer *gfx.RendererLayer, item *gameplay.Item) error {
	graphic := layer.Data.(*RendererLayerItemData).Items[item.Type]
	rect := GetItemRect(layer, item, graphic.Textures[0].Base)
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

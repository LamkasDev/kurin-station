package structure

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetObjectRect(layer *gfx.RendererLayer, obj *gameplay.Object) sdl.Rect {
	graphic := layer.Data.(*RendererLayerObjectData).Structures[obj.Type]
	texture := graphic.Textures[obj.Direction][0]
	return render.WorldToScreenRectF(sdl.FRect{
		X: float32(obj.Tile.Position.Base.X), Y: float32(obj.Tile.Position.Base.Y),
		W: float32(texture.Size.W), H: float32(texture.Size.H),
	})
}

func RenderObject(layer *gfx.RendererLayer, obj *gameplay.Object) error {
	graphic := layer.Data.(*RendererLayerObjectData).Structures[obj.Type]
	rect := GetObjectRect(layer, obj)
	texture := graphic.Textures[obj.Direction][obj.Template.GetTexture(obj)]
	if graphic.Template.Smooth != nil && *graphic.Template.Smooth {
		texture = graphic.TexturesSmooth[gameplay.GetObjectDirectionHint(gameplay.GameInstance.Map, obj)]
	}
	if err := gfx.RendererInstance.Renderer.Copy(texture.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

func RenderObjectBlueprint(layer *gfx.RendererLayer, obj *gameplay.Object, color sdl.Color) error {
	graphic := layer.Data.(*RendererLayerObjectData).Structures[obj.Type]
	if graphic.Blueprint == nil {
		return nil
	}
	rect := GetObjectRect(layer, obj)
	graphic.Blueprint.Texture.SetColorMod(color.R, color.G, color.B)
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Blueprint.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

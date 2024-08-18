package structure

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinObjectRect(layer *gfx.RendererLayer, obj *gameplay.KurinObject) sdl.Rect {
	graphic := layer.Data.(*KurinRendererLayerObjectData).Structures[obj.Type]
	texture := graphic.Textures[obj.Direction][0]
	return render.WorldToScreenRect(sdl.FRect{
		X: float32(obj.Tile.Position.Base.X), Y: float32(obj.Tile.Position.Base.Y),
		W: float32(texture.Size.W), H: float32(texture.Size.H),
	})
}

func RenderKurinObject(layer *gfx.RendererLayer, obj *gameplay.KurinObject) error {
	graphic := layer.Data.(*KurinRendererLayerObjectData).Structures[obj.Type]
	rect := GetKurinObjectRect(layer, obj)
	texture := graphic.Textures[obj.Direction][obj.GetTexture(obj)]
	if graphic.Template.Smooth != nil && *graphic.Template.Smooth {
		texture = graphic.TexturesSmooth[gameplay.GetKurinObjectDirectionHint(&gameplay.GameInstance.Map, obj)]
	}
	if err := gfx.RendererInstance.Renderer.Copy(texture.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

func RenderKurinObjectBlueprint(layer *gfx.RendererLayer, obj *gameplay.KurinObject, color sdl.Color) error {
	graphic := layer.Data.(*KurinRendererLayerObjectData).Structures[obj.Type]
	rect := GetKurinObjectRect(layer, obj)
	graphic.Blueprint.Texture.SetColorMod(color.R, color.G, color.B)
	if err := gfx.RendererInstance.Renderer.Copy(graphic.Blueprint.Texture, nil, &rect); err != nil {
		return err
	}

	return nil
}

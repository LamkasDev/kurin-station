package actions

import (
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRendererLayerActionsData struct {
	Input      string
	Mode       KurinActionMode
	Index      int
	UseTexture sdlutils.TextureWithSize

	ObjectLayer *gfx.KurinRendererLayer
}

type KurinActionMode uint8

const KurinActionModeSay = KurinActionMode(0)
const KurinActionModeBuild = KurinActionMode(1)

func NewKurinRendererLayerActions(objectLayer *gfx.KurinRendererLayer) *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerActions,
		Render: RenderKurinRendererLayerActions,
		Data: KurinRendererLayerActionsData{
			Input:       "",
			Mode:        KurinActionModeSay,
			Index:       0,
			ObjectLayer: objectLayer,
		},
	}
}

func LoadKurinRendererLayerActions(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	data := layer.Data.(KurinRendererLayerActionsData)
	var err *error
	if data.UseTexture, err = sdlutils.LoadTexture(renderer.Renderer, path.Join(constants.TexturesPath, "icons", "radial_use_0.png")); err != nil {
		return err
	}
	layer.Data = data

	return nil
}

func RenderKurinRendererLayerActions(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	if renderer.RendererContext.State != gfx.KurinRendererContextStateActions {
		return nil
	}

	actionsData := layer.Data.(KurinRendererLayerActionsData)
	name := GetKurinActionModeName(actionsData.Mode)
	nameWidth, _, _ := renderer.Fonts.Container[gfx.KurinRendererFontDefault].SizeUTF8(name)

	blue := sdl.Color{R: 66, G: 135, B: 245}
	gray := sdl.Color{R: 36, G: 36, B: 36}
	size := &sdl.Rect{W: 312, H: 36}
	half := gfx.GetHalfWindowSize(&renderer.RendererContext)
	rect := &sdl.Rect{X: int32(float32(half.X) - (float32(size.W) / 2)), Y: int32(float32(half.Y)-(float32(size.H)/2)) - 92, W: size.W, H: size.H}
	irect := sdl.Rect{X: rect.X + 4, Y: rect.Y + 4, W: int32(nameWidth) + 16, H: size.H - 8}

	sdlutils.SetDrawColor(renderer.Renderer, blue)
	renderer.Renderer.FillRect(&sdl.Rect{X: rect.X - 2, Y: rect.Y - 2, W: rect.W + 4, H: rect.H + 4})

	input := actionsData.Input

	sdlutils.SetDrawColor(renderer.Renderer, gray)
	renderer.Renderer.FillRect(rect)
	sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Container[gfx.KurinRendererFontDefault], input, sdl.Color{R: 255, G: 255, B: 255}, sdl.Point{X: rect.X + 10 + irect.W, Y: rect.Y + 10}, sdl.FPoint{X: 1, Y: 1})

	sdlutils.SetDrawColor(renderer.Renderer, blue)
	renderer.Renderer.DrawRect(&irect)
	sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Container[gfx.KurinRendererFontDefault], name, blue, sdl.Point{X: rect.X + 12, Y: rect.Y + 10}, sdl.FPoint{X: 1, Y: 1})

	if actionsData.Mode == KurinActionModeBuild {
		structures := GetMenuStructureGraphics(&actionsData)
		structuresLength := len(structures)
		if structuresLength == 0 {
			return nil
		}

		menurect := &sdl.Rect{X: rect.X, Y: rect.Y + rect.H + 8, W: rect.W, H: int32(structuresLength)*72 + int32(structuresLength-1)*2}
		sdlutils.SetDrawColor(renderer.Renderer, blue)
		renderer.Renderer.FillRect(&sdl.Rect{X: menurect.X - 2, Y: menurect.Y - 2, W: menurect.W + 4, H: menurect.H + 4})
		sdlutils.SetDrawColor(renderer.Renderer, gray)
		renderer.Renderer.FillRect(menurect)

		for i, structureGraphic := range structures {
			if actionsData.Index == i {
				sdlutils.SetDrawColor(renderer.Renderer, blue)
				renderer.Renderer.FillRect(&sdl.Rect{X: menurect.X + 8, Y: menurect.Y + 8, W: 4, H: 56})
				menurect.X += 16
			}

			structureTexture := structureGraphic.Textures[0]
			structurePoint := sdl.Point{X: menurect.X + 12, Y: menurect.Y + 12}
			sdlutils.RenderTexture(renderer.Renderer, structureTexture, structurePoint, sdl.FPoint{X: 1.5, Y: 1.5})
			if structureGraphic.Template.Name != nil {
				sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Container[gfx.KurinRendererFontDefault], *structureGraphic.Template.Name, blue, sdl.Point{X: menurect.X + 72, Y: menurect.Y + 12}, sdl.FPoint{X: 1, Y: 1})
			}
			if structureGraphic.Template.Description != nil {
				sdlutils.RenderUTF8SolidTexture(renderer.Renderer, renderer.Fonts.Container[gfx.KurinRendererFontDefault], *structureGraphic.Template.Description, blue, sdl.Point{X: menurect.X + 72, Y: menurect.Y + 36}, sdl.FPoint{X: 1, Y: 1})
			}

			if actionsData.Index == i {
				menurect.X -= 16
			}

			menurect.Y += 72
			if i+1 < structuresLength {
				sdlutils.SetDrawColor(renderer.Renderer, blue)
				renderer.Renderer.FillRect(&sdl.Rect{X: menurect.X + 12, Y: menurect.Y, W: menurect.W - 24, H: 1})
				menurect.Y += 2
			}
		}
	}

	return nil
}

func GetKurinActionModeName(mode KurinActionMode) string {
	switch mode {
	case KurinActionModeSay:
		return "Say"
	case KurinActionModeBuild:
		return "Build"
	}

	return "??"
}

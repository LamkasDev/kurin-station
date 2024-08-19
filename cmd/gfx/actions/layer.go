package actions

import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"github.com/veandco/go-sdl2/sdl"
)

type RendererLayerActionsData struct {
	Input      string
	Mode       ActionMode
	Index      int
	UseTexture *sdlutils.TextureWithSize

	TurfLayer   *gfx.RendererLayer
	ObjectLayer *gfx.RendererLayer
	ItemLayer   *gfx.RendererLayer
}

type ActionMode uint8

const (
	ActionModeSay   = ActionMode(0)
	ActionModeBuild = ActionMode(1)
)

func NewRendererLayerActions(turfLayer *gfx.RendererLayer, objectLayer *gfx.RendererLayer, itemLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerActions,
		Render: RenderRendererLayerActions,
		Data: &RendererLayerActionsData{
			Input:       "",
			Mode:        ActionModeSay,
			Index:       0,
			TurfLayer:   turfLayer,
			ObjectLayer: objectLayer,
			ItemLayer:   itemLayer,
		},
	}
}

func LoadRendererLayerActions(layer *gfx.RendererLayer) error {
	data := layer.Data.(*RendererLayerActionsData)
	var err error
	if data.UseTexture, err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, path.Join(constants.TexturesPath, "icons", "radial_use.png")); err != nil {
		return err
	}

	return nil
}

func RenderRendererLayerActions(layer *gfx.RendererLayer) error {
	if gfx.RendererInstance.Context.State != gfx.RendererContextStateActions {
		return nil
	}

	data := layer.Data.(*RendererLayerActionsData)
	name := GetActionModeName(data.Mode)
	nameWidth, _, _ := gfx.RendererInstance.Fonts.Default.SizeUTF8(name)

	size := &sdl.Rect{W: 312, H: 36}
	half := gfx.GetHalfWindowSize(&gfx.RendererInstance.Context)
	rect := &sdl.Rect{X: int32(float32(half.X) - (float32(size.W) / 2)), Y: int32(float32(half.Y)-(float32(size.H)/2)) - 92, W: size.W, H: size.H}
	irect := sdl.Rect{X: rect.X + 4, Y: rect.Y + 4, W: int32(nameWidth) + 16, H: size.H - 8}

	sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Blue)
	gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: rect.X - 2, Y: rect.Y - 2, W: rect.W + 4, H: rect.H + 4})

	input := data.Input

	sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.DarkGray)
	gfx.RendererInstance.Renderer.FillRect(rect)
	sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "actions.input", gfx.RendererInstance.Fonts.Default, sdlutils.White, input, sdl.Point{X: rect.X + 10 + irect.W, Y: rect.Y + 10}, sdl.FPoint{X: 1, Y: 1})

	sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Blue)
	gfx.RendererInstance.Renderer.DrawRect(&irect)
	sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "actions.name", gfx.RendererInstance.Fonts.Default, sdl.Color{R: 66, G: 135, B: 245}, name, sdl.Point{X: rect.X + 12, Y: rect.Y + 10}, sdl.FPoint{X: 1, Y: 1})

	switch data.Mode {
	case ActionModeBuild:
		structures := GetMenuGraphics(data)
		if len(structures) == 0 {
			return nil
		}

		min := data.Index
		if min >= len(structures)-4 {
			min = mathutils.MaxInt(0, len(structures)-4)
		}
		max := mathutils.MinInt(min+4, len(structures))
		displayed := max - min

		menurect := &sdl.Rect{X: rect.X, Y: rect.Y + rect.H + 8, W: rect.W, H: int32(displayed)*72 + int32(displayed-1)*2}
		sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Blue)
		gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: menurect.X - 2, Y: menurect.Y - 2, W: menurect.W + 4, H: menurect.H + 4})
		sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.DarkGray)
		gfx.RendererInstance.Renderer.FillRect(menurect)

		for i := min; i < max; i++ {
			structureGraphic := structures[i]
			if data.Index == i {
				sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Blue)
				gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: menurect.X + 8, Y: menurect.Y + 8, W: 4, H: 56})
				menurect.X += 16
			}

			var structureTexture *sdlutils.TextureWithSize
			var structureName string
			var structureRequirements []gameplay.ItemRequirement
			switch data := structureGraphic.(type) {
			case *structure.StructureGraphic:
				structureTexture = data.Textures[0][0]
				structureName = data.Template.Name
				structureRequirements = gameplay.ObjectContainer[data.Template.Id].Requirements
				break
			case *turf.TurfGraphic:
				structureTexture = data.Textures[0]
				structureName = data.Template.Name
				break
			}

			structurePoint := sdl.Point{X: menurect.X + 12, Y: menurect.Y + 12}
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, structureTexture, structurePoint, sdl.FPoint{X: 1.5, Y: 1.5})
			sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("actions.%d.name", i), gfx.RendererInstance.Fonts.Default, sdlutils.Blue, structureName, sdl.Point{X: menurect.X + 72, Y: menurect.Y + 12}, sdl.FPoint{X: 1, Y: 1})
			if structureRequirements != nil {
				for j, requirement := range structureRequirements {
					sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("actions.%d.requirements.%d", j), gfx.RendererInstance.Fonts.Default, sdlutils.Blue, fmt.Sprintf("%dx %s", requirement.Count, requirement.Type), sdl.Point{X: menurect.X + 72, Y: menurect.Y + 36}, sdl.FPoint{X: 1, Y: 1})
				}
			}

			if data.Index == i {
				menurect.X -= 16
			}

			menurect.Y += 72
			if i+1 < max {
				sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Blue)
				gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: menurect.X + 12, Y: menurect.Y, W: menurect.W - 24, H: 1})
				menurect.Y += 2
			}
		}
	}

	return nil
}

func GetActionModeName(mode ActionMode) string {
	switch mode {
	case ActionModeSay:
		return "Say"
	case ActionModeBuild:
		return "Build"
	}

	return "??"
}

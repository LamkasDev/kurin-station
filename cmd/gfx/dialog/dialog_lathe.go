package dialog

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func NewKurinDialogLathe(layer *gfx.RendererLayer, data interface{}) *KurinDialog {
	layerData := layer.Data.(*KurinRendererLayerDialogData)
	itemLayerData := layerData.ItemLayer.Data.(*item.KurinRendererLayerItemData)

	dialog := NewKurinDialogRaw(layer, "lathe", "Lathe", "flushed")
	manufactureBar := &KurinDialogElement{
		GetRect: func(dialogRect *sdl.Rect) *sdl.Rect {
			return &sdl.Rect{X: 8, Y: 8, W: dialogRect.W - 16, H: 32}
		},
		Render: func(element *KurinDialogElement, rect *sdl.Rect) {
			sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Gray)
			gfx.RendererInstance.Renderer.FillRect(rect)

			latheData := data.(*gameplay.KurinDialogLatheData).Lathe.Data.(*gameplay.KurinObjectLatheData)
			if latheData.Order != nil {
				w := float32(rect.W) * ((float32(latheData.Order.TotalTicks) - float32(latheData.Order.TicksLeft)) / float32(latheData.Order.TotalTicks))
				sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Blue)
				gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: rect.X, Y: rect.Y, W: int32(w), H: rect.H})

				sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, itemLayerData.Items[latheData.Order.ItemType].Textures[0].Base, sdl.Rect{X: rect.X + 2, Y: rect.Y + 2, W: 28, H: 28})
				text := fmt.Sprintf("Manufacturing 1x %s%s", latheData.Order.ItemType, render.GetThreeDots())
				sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "dialog.lathe.text", gfx.RendererInstance.Fonts.Default, sdlutils.White, text, sdl.Point{X: rect.X + 36, Y: rect.Y + 8}, sdl.FPoint{X: 1, Y: 1})
			} else {
				sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "dialog.lathe.text", gfx.RendererInstance.Fonts.Default, sdlutils.White, "No order...", sdl.Point{X: rect.X + 10, Y: rect.Y + 8}, sdl.FPoint{X: 1, Y: 1})
			}
		},
	}
	dialog.Elements = append(dialog.Elements, manufactureBar)
	orderButton := NewKurinDialogElementButton(sdl.Point{X: 8, Y: 46}, "1x Rod", func(dialog *KurinDialog) {
		data := data.(*gameplay.KurinDialogLatheData)
		latheData := data.Lathe.Data.(*gameplay.KurinObjectLatheData)
		if latheData.Order == nil {
			gameplay.CreateNewOrderAtLathe(data.Lathe, gameplay.NewKurinLatheOrder("rod"))
		}
	})
	dialog.Elements = append(dialog.Elements, orderButton)
	dialog.ShouldClose = func(dialog *KurinDialog) bool {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return true
		}

		dialogData := dialog.Data.(*gameplay.KurinDialogLatheData)
		return sdlutils.GetDistanceSimpleF(sdlutils.PointToFPointCenter(gameplay.GameInstance.SelectedCharacter.Position.Base), gameplay.GetKurinObjectCenter(dialogData.Lathe)) > 1.5
	}
	dialog.Data = data

	return dialog
}

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

func NewDialogLathe(layer *gfx.RendererLayer, data interface{}) *Dialog {
	layerData := layer.Data.(*RendererLayerDialogData)
	itemLayerData := layerData.ItemLayer.Data.(*item.RendererLayerItemData)

	dialog := NewDialogRaw(layer, "lathe", "Lathe", "flushed")
	dialog.ShouldClose = func(dialog *Dialog) bool {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return true
		}

		dialogData := dialog.Data.(*gameplay.DialogLatheData)
		return sdlutils.GetDistanceSimpleF(sdlutils.PointToFPointCenter(gameplay.GameInstance.SelectedCharacter.Position.Base), gameplay.GetObjectCenter(dialogData.Lathe)) > 1.5
	}
	dialog.Data = data

	energyBar := &DialogElement{
		GetRect: func(dialogRect *sdl.Rect) *sdl.Rect {
			return &sdl.Rect{X: 8, Y: 8, W: dialogRect.W - 16, H: 32}
		},
		Render: func(element *DialogElement, rect *sdl.Rect) {
			latheData := data.(*gameplay.DialogLatheData).Lathe.Data.(*gameplay.ObjectLatheData)

			sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Gray)
			gfx.RendererInstance.Renderer.FillRect(rect)

			w := float32(rect.W) * (float32(latheData.Energy) / float32(latheData.MaxEnergy))
			sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Orange)
			gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: rect.X, Y: rect.Y, W: int32(w), H: rect.H})

			text := fmt.Sprintf("Energy (%d/%d)", latheData.Energy, latheData.MaxEnergy)
			sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "dialog.lathe.energy", gfx.RendererInstance.Fonts.Default, sdlutils.White, text, sdl.Point{X: rect.X + 8, Y: rect.Y + 8}, sdl.FPoint{X: 1, Y: 1})
		},
	}
	dialog.Elements = append(dialog.Elements, energyBar)

	manufactureBar := &DialogElement{
		GetRect: func(dialogRect *sdl.Rect) *sdl.Rect {
			return &sdl.Rect{X: 8, Y: 46, W: dialogRect.W - 16, H: 32}
		},
		Render: func(element *DialogElement, rect *sdl.Rect) {
			latheData := data.(*gameplay.DialogLatheData).Lathe.Data.(*gameplay.ObjectLatheData)

			sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Gray)
			gfx.RendererInstance.Renderer.FillRect(rect)

			if len(latheData.Orders) == 0 {
				sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "dialog.lathe.text", gfx.RendererInstance.Fonts.Default, sdlutils.White, "No order...", sdl.Point{X: rect.X + 10, Y: rect.Y + 8}, sdl.FPoint{X: 1, Y: 1})
			} else {
				w := float32(rect.W) * ((float32(latheData.Orders[0].TotalTicks) - float32(latheData.Orders[0].TicksLeft)) / float32(latheData.Orders[0].TotalTicks))
				sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.Blue)
				gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: rect.X, Y: rect.Y, W: int32(w), H: rect.H})

				sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, itemLayerData.Items[latheData.Orders[0].ItemType].Textures[0].Base, sdl.Rect{X: rect.X + 2, Y: rect.Y + 2, W: 28, H: 28})
				orderText := fmt.Sprintf("%dx %s", latheData.Orders[0].ItemCount, latheData.Orders[0].ItemType)
				if len(latheData.Orders) > 1 {
					orderText = fmt.Sprintf("%s (+%d)", orderText, len(latheData.Orders)-1)
				}
				text := fmt.Sprintf("Manufacturing %s%s", orderText, render.GetThreeDots())
				sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "dialog.lathe.text", gfx.RendererInstance.Fonts.Default, sdlutils.White, text, sdl.Point{X: rect.X + 36, Y: rect.Y + 8}, sdl.FPoint{X: 1, Y: 1})
			}
		},
	}
	dialog.Elements = append(dialog.Elements, manufactureBar)

	orderButton := NewDialogElementButton(sdl.Point{X: 8, Y: 84}, "1x Rod", func(dialog *Dialog) {
		data := data.(*gameplay.DialogLatheData)
		gameplay.CreateNewOrderAtLathe(data.Lathe, gameplay.NewLatheOrder("rod", 1))
	})
	dialog.Elements = append(dialog.Elements, orderButton)

	orderButton2 := NewDialogElementButton(sdl.Point{X: 84, Y: 84}, "3x Rod", func(dialog *Dialog) {
		data := data.(*gameplay.DialogLatheData)
		gameplay.CreateNewOrderAtLathe(data.Lathe, gameplay.NewLatheOrder("rod", 3))
	})
	dialog.Elements = append(dialog.Elements, orderButton2)

	return dialog
}

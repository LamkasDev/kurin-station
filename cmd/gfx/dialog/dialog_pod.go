package dialog

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func NewDialogConsole(layer *gfx.RendererLayer, data interface{}) *Dialog {
	dialog := NewDialogRaw(layer, "console", "Console", "flushed")
	testButton := NewDialogElementButton(sdl.Point{X: 8, Y: 8}, "test", func(dialog *Dialog) {
		gameplay.PlaySound(&gameplay.GameInstance.SoundController, "grillehit")
	})
	dialog.Elements = append(dialog.Elements, testButton)
	testLabel := NewDialogElementLabel(sdl.Point{X: 8, Y: 44}, "test", sdlutils.GetTextureFromContainer(gfx.RendererInstance.IconTextures, gfx.RendererInstance.Renderer, "flushed"))
	dialog.Elements = append(dialog.Elements, testLabel)
	testLabel2 := NewDialogElementLabel(sdl.Point{X: 8, Y: 66}, "test", nil)
	dialog.Elements = append(dialog.Elements, testLabel2)
	dialog.ShouldClose = func(dialog *Dialog) bool {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return true
		}

		dialogData := dialog.Data.(*gameplay.DialogConsoleData)
		return sdlutils.GetDistanceSimpleF(sdlutils.PointToFPointCenter(gameplay.GameInstance.SelectedCharacter.Position.Base), gameplay.GetObjectCenter(dialogData.Console)) > 1.5
	}
	dialog.Data = data

	return dialog
}

package dialog

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func NewKurinDialogPod(layer *gfx.RendererLayer, data interface{}) *KurinDialog {
	// dialogData := data.(*gameplay.KurinDialogPodData)
	dialog := NewKurinDialogRaw(layer, "pod", "Pod", "flushed")
	testButton := NewKurinDialogElementButton(sdl.Point{X: 8, Y: 8}, "test", func(dialog *KurinDialog) {
		gameplay.PlaySound(&gameplay.GameInstance.SoundController, "grillehit")
	})
	dialog.Elements = append(dialog.Elements, testButton)
	testLabel := NewKurinDialogElementLabel(sdl.Point{X: 8, Y: 44}, "test", sdlutils.GetTextureFromContainer(gfx.RendererInstance.IconTextures, gfx.RendererInstance.Renderer, "flushed"))
	dialog.Elements = append(dialog.Elements, testLabel)
	testLabel2 := NewKurinDialogElementLabel(sdl.Point{X: 8, Y: 66}, "test", nil)
	dialog.Elements = append(dialog.Elements, testLabel2)
	dialog.ShouldClose = func(dialog *KurinDialog) bool {
		if gameplay.GameInstance.SelectedCharacter == nil {
			return true
		}

		dialogData := dialog.Data.(*gameplay.KurinDialogPodData)
		return sdlutils.GetDistanceSimpleF(sdlutils.PointToFPointCenter(gameplay.GameInstance.SelectedCharacter.Position.Base), gameplay.GetKurinObjectCenter(dialogData.Pod)) > 1.5
	}
	dialog.Data = data

	return dialog
}

package dialog

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/dialog"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLayerDialogData struct {
	Layer *gfx.RendererLayer
}

func NewEventLayerDialog(layer *gfx.RendererLayer) *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerDialog,
		Process: ProcessEventLayerDialog,
		Data: &EventLayerDialogData{
			Layer: layer,
		},
	}
}

func LoadEventLayerDialog(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerDialog(layer *event.EventLayer) error {
	data := layer.Data.(*EventLayerDialogData)
	dialogData := data.Layer.Data.(*dialog.RendererLayerDialogData)
	if gameplay.GameInstance.DialogController.OpenRequest != nil {
		dialogData.Dialog = dialog.NewDialog(data.Layer, gameplay.GameInstance.DialogController.OpenRequest)
		gameplay.GameInstance.DialogController.OpenRequest = nil
	}
	if gameplay.GameInstance.DialogController.CloseRequest || gfx.RendererInstance.Context.State != gfx.RendererContextStateNone {
		dialogData.Dialog = nil
		gameplay.GameInstance.DialogController.CloseRequest = false
	}
	if dialogData.Dialog == nil {
		return nil
	}
	if event.EventManagerInstance.Keyboard.Pending != nil && *event.EventManagerInstance.Keyboard.Pending == sdl.K_ESCAPE {
		event.EventManagerInstance.Keyboard.Pending = nil
		dialogData.Dialog = nil
		return nil
	}
	if dialogData.Dialog.ShouldClose(dialogData.Dialog) {
		dialogData.Dialog = nil
		return nil
	}

	dialogSize := dialogData.Dialog.GetSize(gfx.RendererInstance.Context.WindowSize)
	dialogRect := &sdl.Rect{
		X: dialogData.Dialog.Position.X,
		Y: dialogData.Dialog.Position.Y + 32,
		W: dialogSize.X,
		H: dialogSize.Y,
	}

	// Process all the dialog elements
	for _, element := range dialogData.Dialog.Elements {
		if !element.Clickable {
			continue
		}
		rect := element.GetRect(dialogRect)
		rect.X += dialogRect.X
		rect.Y += dialogRect.Y
		element.Hovered = gfx.RendererInstance.Context.MousePosition.InRect(rect)
		if element.Hovered {
			event.EventManagerInstance.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
			if event.EventManagerInstance.Mouse.PendingLeft != nil {
				event.EventManagerInstance.Mouse.PendingLeft = nil
				element.OnClick(dialogData.Dialog)
			}
		}
	}
	dialogRect.Y -= 32
	dialogRect.H += 32

	// Dragging the bar around
	if dialogData.Dialog.Dragged {
		if event.EventManagerInstance.Mouse.PressedLeft {
			dialogData.Dialog.Position = sdlutils.AddPoints(dialogData.Dialog.Position, event.EventManagerInstance.Mouse.Delta)
		} else {
			dialogData.Dialog.Dragged = false
		}
	} else if event.EventManagerInstance.Mouse.PressedLeft && gfx.RendererInstance.Context.MousePosition.InRect(&sdl.Rect{X: dialogRect.X, Y: dialogRect.Y, W: dialogRect.W - 32, H: 32}) {
		dialogData.Dialog.Dragged = true
	}

	// Process close button
	if gfx.RendererInstance.Context.MousePosition.InRect(&sdl.Rect{X: dialogRect.X + dialogRect.W - 32, Y: dialogRect.Y, W: 32, H: 32}) {
		event.EventManagerInstance.Mouse.Cursor = sdl.SYSTEM_CURSOR_HAND
		if event.EventManagerInstance.Mouse.PendingLeft != nil {
			event.EventManagerInstance.Mouse.PendingLeft = nil
			dialogData.Dialog = nil
		}
	}

	// Cancel clicks if in dialog
	if gfx.RendererInstance.Context.MousePosition.InRect(dialogRect) {
		event.EventManagerInstance.Mouse.PendingLeft = nil
		event.EventManagerInstance.Mouse.PendingRight = nil
	}

	return nil
}

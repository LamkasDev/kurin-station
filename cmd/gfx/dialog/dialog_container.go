package dialog

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

func NewDialog(layer *gfx.RendererLayer, request *gameplay.DialogRequest) *Dialog {
	switch request.Type {
	case "console":
		return NewDialogConsole(layer, request.Data)
	case "lathe":
		return NewDialogLathe(layer, request.Data)
	}

	return NewDialogRaw(layer, request.Type, "??", "flushed")
}

package dialog

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

func NewKurinDialog(layer *gfx.RendererLayer, request *gameplay.KurinDialogRequest) *KurinDialog {
	switch request.Type {
	case "pod":
		return NewKurinDialogPod(layer, request.Data)
	case "lathe":
		return NewKurinDialogLathe(layer, request.Data)
	}

	return NewKurinDialogRaw(layer, request.Type, "??", "flushed")
}

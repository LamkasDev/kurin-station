package dialog

import (
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinDialog struct {
	Type        string
	Title       string
	Icon        string
	Elements    []*KurinDialogElement
	GetSize     KurinDialogGetSize
	ShouldClose KurinDialogShouldClose
	Movable     bool

	Position sdl.Point
	Dragged  bool
	Data     interface{}
}

type KurinDialogElement struct {
	GetRect   KurinDialogElementGetRect
	Render    KurinDialogElementRender
	OnClick   KurinDialogElementOnClick
	Clickable bool

	Hovered bool
}

type (
	KurinDialogGetSize        func(windowSize sdl.Point) sdl.Point
	KurinDialogShouldClose    func(dialog *KurinDialog) bool
	KurinDialogElementGetRect func(dialogRect *sdl.Rect) *sdl.Rect
	KurinDialogElementRender  func(element *KurinDialogElement, rect *sdl.Rect)
	KurinDialogElementOnClick func(dialog *KurinDialog)
)

func NewKurinDialogRaw(layer *gfx.RendererLayer, dialogType string, title string, icon string) *KurinDialog {
	dialog := &KurinDialog{
		Type:     dialogType,
		Title:    title,
		Icon:     icon,
		Elements: []*KurinDialogElement{},
		GetSize: func(windowSize sdl.Point) sdl.Point {
			return sdl.Point{X: 256, Y: 256}
		},
		ShouldClose: func(dialog *KurinDialog) bool {
			return false
		},
		Position: sdl.Point{
			X: 64,
			Y: 64,
		},
	}

	return dialog
}

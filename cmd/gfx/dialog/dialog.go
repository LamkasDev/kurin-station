package dialog

import (
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Dialog struct {
	Type        string
	Title       string
	Icon        string
	Elements    []*DialogElement
	GetSize     DialogGetSize
	ShouldClose DialogShouldClose
	Movable     bool

	Position sdl.Point
	Dragged  bool
	Data     interface{}
}

type DialogElement struct {
	GetRect   DialogElementGetRect
	Render    DialogElementRender
	OnClick   DialogElementOnClick
	Clickable bool

	Hovered bool
}

type (
	DialogGetSize        func(windowSize sdl.Point) sdl.Point
	DialogShouldClose    func(dialog *Dialog) bool
	DialogElementGetRect func(dialogRect *sdl.Rect) *sdl.Rect
	DialogElementRender  func(element *DialogElement, rect *sdl.Rect)
	DialogElementOnClick func(dialog *Dialog)
)

func NewDialogRaw(layer *gfx.RendererLayer, dialogType string, title string, icon string) *Dialog {
	dialog := &Dialog{
		Type:     dialogType,
		Title:    title,
		Icon:     icon,
		Elements: []*DialogElement{},
		GetSize: func(windowSize sdl.Point) sdl.Point {
			return sdl.Point{X: 256, Y: 256}
		},
		ShouldClose: func(dialog *Dialog) bool {
			return false
		},
		Position: sdl.Point{
			X: 64,
			Y: 64,
		},
	}

	return dialog
}

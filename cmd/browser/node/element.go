package node

import (
	"github.com/veandco/go-sdl2/sdl"
)

type KitsuneElement struct {
	Parent   *KitsuneElement
	Children []*KitsuneElement

	Size    KitsuneElementSize
	OwnSize KitsuneElementSize
	Margin  KitsuneElementRectLTRB
	Padding KitsuneElementRectLTRB
	Color   sdl.Color

	Data       interface{}
	CachedRect sdl.Rect
}

type KitsuneElementSize struct {
	Width  KitsuneElementValueVariable
	Height KitsuneElementValueVariable
}

type KitsuneElementValueVariable struct {
	Type  KitsuneElementValueVariableType
	Value int32
}

type KitsuneElementValueVariableType int8

const KitsuneElementValueVariableFixed = KitsuneElementValueVariableType(0x00)

func NewKitsuneElementSize(width int32, height int32) KitsuneElementSize {
	return KitsuneElementSize{
		Width: KitsuneElementValueVariable{
			Type:  KitsuneElementValueVariableFixed,
			Value: width,
		},
		Height: KitsuneElementValueVariable{
			Type:  KitsuneElementValueVariableFixed,
			Value: height,
		},
	}
}

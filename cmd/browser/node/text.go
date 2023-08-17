package node

import "github.com/veandco/go-sdl2/sdl"

type KitsuneElementTextData struct {
	Text *sdl.Texture
	Size sdl.Rect
}

type KitsuneElementLinkData struct {
	Base *KitsuneElementTextData
}

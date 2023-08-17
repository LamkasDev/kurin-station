package context

import (
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"github.com/veandco/go-sdl2/sdl"
)

type KitsuneContextRender struct {
	Title             string
	Icon              *sdl.Texture
	Document          *node.KitsuneElement
	ScrollPosition    sdl.FRect
	ScrollDestination sdl.FRect
}

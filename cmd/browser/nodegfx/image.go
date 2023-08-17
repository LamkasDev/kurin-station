package nodegfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/gfx"
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"golang.org/x/net/html"
)

func NewKitsuneElementImageData(renderer *gfx.KitsuneRenderer, htmlNode *html.Node) (*node.KitsuneElementImageData, *error) {
	return &node.KitsuneElementImageData{
		Image: nil,
	}, nil
}

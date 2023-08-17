package contextgfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/context"
	"github.com/LamkasDev/kitsune/cmd/browser/gfx"
	"github.com/LamkasDev/kitsune/cmd/browser/http"
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"github.com/LamkasDev/kitsune/cmd/browser/nodegfx"
	"golang.org/x/net/html"
)

func NewKitsuneContextRender(renderer *gfx.KitsuneRenderer, response *http.KitsuneClientResponse) (*context.KitsuneContextRender, *error) {
	rcontext := &context.KitsuneContextRender{}
	var err *error
	if rcontext.Document, err = ProcessNode(renderer, rcontext, response.Node); err != nil {
		return nil, err
	}

	return rcontext, nil
}

// TODO: move this crap onto nodegfx
func ProcessNode(renderer *gfx.KitsuneRenderer, rcontext *context.KitsuneContextRender, current *html.Node) (*node.KitsuneElement, *error) {
	if current.Type == html.ErrorNode {
		return nil, nil
	}

	element, err := nodegfx.NewKitsuneElement(renderer, rcontext, current)
	if err != nil {
		return nil, err
	}
	if element != nil {
		children, err := ProcessNodeChildren(renderer, rcontext, current.FirstChild)
		if err != nil {
			return nil, err
		}
		for _, child := range children {
			child.Parent = element
		}
		element.Children = children
		element.Size = CalculateKitsuneElementSize(element)
	}

	return element, nil
}

// TODO: move this crap onto nodegfx
func ProcessNodeChildren(renderer *gfx.KitsuneRenderer, rcontext *context.KitsuneContextRender, current *html.Node) ([]*node.KitsuneElement, *error) {
	elements := []*node.KitsuneElement{}
	for current != nil {
		element, err := ProcessNode(renderer, rcontext, current)
		if err != nil {
			return nil, err
		}
		if element != nil {
			elements = append(elements, element)
		}
		current = current.NextSibling
	}

	return elements, nil
}

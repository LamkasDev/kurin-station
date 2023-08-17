package nodegfx

import (
	"github.com/LamkasDev/kitsune/cmd/browser/context"
	"github.com/LamkasDev/kitsune/cmd/browser/gfx"
	"github.com/LamkasDev/kitsune/cmd/browser/node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func NewKitsuneElementEmpty() *node.KitsuneElement {
	return &node.KitsuneElement{
		Children: []*node.KitsuneElement{},
	}
}

func NewKitsuneElement(renderer *gfx.KitsuneRenderer, rcontext *context.KitsuneContextRender, htmlNode *html.Node) (*node.KitsuneElement, *error) {
	if htmlNode.Type == html.ErrorNode {
		return nil, nil
	}
	element := NewKitsuneElementEmpty()
	switch htmlNode.DataAtom {
	case atom.Body:
		element.Margin = node.NewKitsuneElementRect(8)
	case atom.Title:
		rcontext.Title = htmlNode.FirstChild.Data
		return nil, nil
	case atom.Ul:
		element.Margin = node.NewKitsuneElementRectHV(0, KitsuneElementMargins[htmlNode.DataAtom])
		element.Padding = node.NewKitsuneElementRectLTRB(40, 0, 0, 0)
	case atom.Li,
		atom.H1,
		atom.H2,
		atom.H3,
		atom.H4,
		atom.H5,
		atom.H6,
		atom.P,
		atom.Aside,
		atom.A:
		data, err := NewKitsuneElementTextData(renderer, htmlNode)
		if err != nil {
			return nil, err
		}
		element.OwnSize = node.NewKitsuneElementSize(data.Size.W, data.Size.H)
		element.Margin = node.NewKitsuneElementRectHV(0, KitsuneElementMargins[htmlNode.DataAtom])
		switch htmlNode.DataAtom {
		case atom.A:
			element.Data = &node.KitsuneElementLinkData{Base: data}
		default:
			element.Data = data
		}
		// TODO: if only one child node -> default, if more than one -> convert them into
	case atom.Img:
		data, err := NewKitsuneElementImageData(renderer, htmlNode)
		if err != nil {
			return nil, err
		}
		element.OwnSize = node.NewKitsuneElementSize(data.Size.W, data.Size.H)
		element.Data = data
	case atom.Html,
		atom.Header,
		atom.Head,
		atom.Div,
		atom.Nav,
		atom.Span:
		break
	default:
		if htmlNode.Type != html.DocumentNode {
			// fmt.Printf("unhandled: %+v\n", htmlNode)
			return nil, nil
		}
	}

	return element, nil
}

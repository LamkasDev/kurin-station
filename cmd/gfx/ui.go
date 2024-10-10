package gfx

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type UIAnchor uint8

var (
	UIAnchorTopLeft      = UIAnchor(0)
	UIAnchorTopCenter    = UIAnchor(1)
	UIAnchorTopRight     = UIAnchor(2)
	UIAnchorCenterLeft   = UIAnchor(3)
	UIAnchorCenter       = UIAnchor(4)
	UIAnchorCenterRight  = UIAnchor(5)
	UIAnchorBottomLeft   = UIAnchor(6)
	UIAnchorBottomCenter = UIAnchor(7)
	UIAnchorBottomRight  = UIAnchor(8)
)

func GetUIRect(texture *sdlutils.TextureWithSize, position sdl.Point, scale sdl.FPoint, anchor UIAnchor) sdl.Rect {
	scale = sdlutils.MultiplyFPoints(scale, RendererInstance.Context.WindowScale)
	rect := sdlutils.GetTextureRect(RendererInstance.Renderer, texture, position, scale)
	halfWindow := GetHalfWindowSize(&RendererInstance.Context)
	halfSize := sdlutils.DividePoint(sdl.Point{X: rect.W, Y: rect.H}, 2)
	switch anchor {
	case UIAnchorTopLeft:
		position = sdl.Point{X: position.X, Y: position.Y}
	case UIAnchorTopCenter:
		position = sdl.Point{X: halfWindow.X - halfSize.X + position.X, Y: position.Y}
	case UIAnchorTopRight:
		position = sdl.Point{X: RendererInstance.Context.WindowSize.X - rect.W - position.X, Y: position.Y}
	case UIAnchorCenterLeft:
		position = sdl.Point{X: position.X, Y: halfWindow.Y - halfSize.Y + position.Y}
	case UIAnchorCenter:
		position = sdl.Point{X: halfWindow.X - halfSize.X + position.X, Y: halfWindow.Y - halfSize.Y + position.Y}
	case UIAnchorCenterRight:
		position = sdl.Point{X: RendererInstance.Context.WindowSize.X - rect.W - position.X, Y: halfWindow.Y - halfSize.Y + position.Y}
	case UIAnchorBottomLeft:
		position = sdl.Point{X: position.X, Y: RendererInstance.Context.WindowSize.Y - rect.H - position.Y}
	case UIAnchorBottomCenter:
		position = sdl.Point{X: halfWindow.X - halfSize.X + position.X, Y: RendererInstance.Context.WindowSize.Y - rect.H - position.Y}
	case UIAnchorBottomRight:
		position = sdl.Point{X: RendererInstance.Context.WindowSize.X - rect.W - position.X, Y: RendererInstance.Context.WindowSize.Y - rect.H - position.Y}
	}

	return sdlutils.GetTextureRect(RendererInstance.Renderer, texture, position, scale)
}

func RenderUITexture(texture *sdlutils.TextureWithSize, position sdl.Point, scale sdl.FPoint, anchor UIAnchor) (sdl.Rect, error) {
	rect := GetUIRect(texture, position, scale, anchor)
	return rect, sdlutils.RenderTextureRect(RendererInstance.Renderer, texture, rect)
}

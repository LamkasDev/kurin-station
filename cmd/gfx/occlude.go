package gfx

import "github.com/veandco/go-sdl2/sdl"

func ShouldOcclude(rect sdl.Rect) bool {
	return rect.X+rect.W < 0 || rect.Y+rect.H < 0 || rect.X > RendererInstance.Context.WindowSize.X || rect.Y > RendererInstance.Context.WindowSize.Y
}

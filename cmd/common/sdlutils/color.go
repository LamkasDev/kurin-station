package sdlutils

import "github.com/veandco/go-sdl2/sdl"

func SetDrawColor(renderer *sdl.Renderer, color sdl.Color) error {
	return renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

package sdlutils

import "github.com/veandco/go-sdl2/sdl"

func SetDrawColor(renderer *sdl.Renderer, color sdl.Color) error {
	return renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

func IsColorVisible(color sdl.Color) bool {
	return color.R > 0 || color.G > 0 || color.B > 0
}

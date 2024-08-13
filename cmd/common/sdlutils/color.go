package sdlutils

import "github.com/veandco/go-sdl2/sdl"

var White = sdl.Color{R: 255, G: 255, B: 255}

func SetDrawColor(renderer *sdl.Renderer, color sdl.Color) error {
	return renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

func IsColorVisible(color sdl.Color) bool {
	return color.R > 0 || color.G > 0 || color.B > 0
}

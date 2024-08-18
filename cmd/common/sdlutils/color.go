package sdlutils

import "github.com/veandco/go-sdl2/sdl"

var (
	Gray       = sdl.Color{R: 82, G: 82, B: 82}
	DarkGray   = sdl.Color{R: 56, G: 56, B: 56}
	Blue       = sdl.Color{R: 66, G: 135, B: 245}
	Red        = sdl.Color{R: 255, G: 0, B: 0}
	White      = sdl.Color{R: 255, G: 255, B: 255}
	LightBlack = sdl.Color{R: 36, G: 36, B: 36}
	Black      = sdl.Color{R: 0, G: 0, B: 0}
)

func SetDrawColor(renderer *sdl.Renderer, color sdl.Color) error {
	return renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

func IsColorVisible(color sdl.Color) bool {
	return color.R > 0 || color.G > 0 || color.B > 0
}

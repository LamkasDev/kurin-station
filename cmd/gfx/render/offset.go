package render

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/arl/math32"
	"github.com/veandco/go-sdl2/sdl"
)

func GetCameraTileSize(renderer *gfx.KurinRenderer) sdl.FPoint {
	return sdlutils.MultiplyFPoints(gameplay.KurinTileSizeF, renderer.Context.CameraZoom)
}

func GetCameraOffset(renderer *gfx.KurinRenderer) sdl.FPoint {
	return sdlutils.SubtractFPoints(
		sdlutils.AddFPoints(
			sdlutils.MultiplyFPoints(renderer.Context.CameraPosition, renderer.Context.CameraTileSize),
			sdlutils.DivideFPointByFloat(renderer.Context.CameraTileSize, 2),
		),
		sdlutils.DivideFPointByFloat(sdlutils.PointToFPoint(renderer.Context.WindowSize), 2),
	)
}

func WorldToScreenPosition(renderer *gfx.KurinRenderer, world sdl.FPoint) sdl.Point {
	return sdlutils.FPointToPoint(sdlutils.SubtractFPoints(sdlutils.MultiplyFPoints(world, renderer.Context.CameraTileSize), renderer.Context.CameraOffset))
}

func WorldToScreenRect(renderer *gfx.KurinRenderer, world sdl.FRect) sdl.Rect {
	pos := WorldToScreenPosition(renderer, sdl.FPoint{X: world.X, Y: world.Y})
	return sdl.Rect{
		X: pos.X,
		Y: pos.Y,
		W: int32(world.W * renderer.Context.CameraZoom.X), H: int32(world.H * renderer.Context.CameraZoom.Y),
	}
}

func ScreenToWorldPosition(renderer *gfx.KurinRenderer, screen sdl.Point) sdl.Point {
	return sdl.Point{
		X: int32(math32.Floor((renderer.Context.CameraOffset.X + float32(screen.X)) / renderer.Context.CameraTileSize.X)),
		Y: int32(math32.Floor((renderer.Context.CameraOffset.Y + float32(screen.Y)) / renderer.Context.CameraTileSize.Y)),
	}
}

func ScreenToWorldRect(renderer *gfx.KurinRenderer, screen sdl.Rect) sdl.Rect {
	pos := ScreenToWorldPosition(renderer, sdl.Point{X: screen.X, Y: screen.Y})
	return sdl.Rect{
		X: pos.X,
		Y: pos.Y,
		W: screen.W, H: screen.H,
	}
}

package render

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/arl/math32"
	"github.com/veandco/go-sdl2/sdl"
)

func GetCameraTileSize(renderer *gfx.KurinRenderer) sdl.FPoint {
	return sdlutils.MultiplyFPoints(sdlutils.PointToFPoint(gameplay.KurinTileSize), renderer.RendererContext.CameraZoom)
}

func GetCameraOffset(renderer *gfx.KurinRenderer) sdl.FPoint {
	return sdlutils.SubtractFPoints(
		sdlutils.AddFPoints(
			sdlutils.MultiplyFPoints(renderer.RendererContext.CameraPosition, renderer.RendererContext.CameraTileSize),
			sdlutils.DivideFPointByFloat(renderer.RendererContext.CameraTileSize, 2),
		),
		sdlutils.DivideFPointByFloat(sdlutils.PointToFPoint(renderer.RendererContext.WindowSize), 2),
	)
}

func WorldToScreenPosition(renderer *gfx.KurinRenderer, world sdl.FPoint) sdl.Point {
	return sdlutils.FPointToPoint(sdlutils.SubtractFPoints(sdlutils.MultiplyFPoints(world, renderer.RendererContext.CameraTileSize), renderer.RendererContext.CameraOffset))
}

func WorldToScreenRect(renderer *gfx.KurinRenderer, world sdl.FRect) sdl.Rect {
	pos := WorldToScreenPosition(renderer, sdl.FPoint{X: world.X, Y: world.Y})
	return sdl.Rect{
		X: pos.X,
		Y: pos.Y,
		W: int32(world.W * renderer.RendererContext.CameraZoom.X), H: int32(world.H * renderer.RendererContext.CameraZoom.Y),
	}
}

func ScreenToWorldPosition(renderer *gfx.KurinRenderer, screen sdl.Point) sdl.Point {
	return sdl.Point{
		X: int32(math32.Floor((renderer.RendererContext.CameraOffset.X + float32(screen.X)) / renderer.RendererContext.CameraTileSize.X)),
		Y: int32(math32.Floor((renderer.RendererContext.CameraOffset.Y + float32(screen.Y)) / renderer.RendererContext.CameraTileSize.Y)),
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

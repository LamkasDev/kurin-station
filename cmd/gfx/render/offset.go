package render

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/arl/math32"
	"github.com/veandco/go-sdl2/sdl"
)

func GetCameraTileSize() sdl.FPoint {
	return sdlutils.MultiplyFPoints(gameplay.TileSizeF, gfx.RendererInstance.Context.CameraZoom)
}

func GetCameraOffset() sdl.FPoint {
	return sdlutils.SubtractFPoints(
		sdlutils.AddFPoints(
			sdlutils.MultiplyFPoints(gfx.RendererInstance.Context.CameraPosition, gfx.RendererInstance.Context.CameraTileSize),
			sdlutils.DivideFPoint(gfx.RendererInstance.Context.CameraTileSize, 2),
		),
		sdlutils.DivideFPoint(sdlutils.PointToFPoint(gfx.RendererInstance.Context.WindowSize), 2),
	)
}

func WorldToScreenPosition(world sdl.FPoint) sdl.Point {
	return sdlutils.FPointToPoint(sdlutils.SubtractFPoints(sdlutils.MultiplyFPoints(world, gfx.RendererInstance.Context.CameraTileSize), gfx.RendererInstance.Context.CameraOffset))
}

func WorldToScreenRect(world sdl.FRect) sdl.Rect {
	pos := WorldToScreenPosition(sdl.FPoint{X: world.X, Y: world.Y})
	return sdl.Rect{
		X: pos.X,
		Y: pos.Y,
		W: int32(world.W * gfx.RendererInstance.Context.CameraZoom.X), H: int32(world.H * gfx.RendererInstance.Context.CameraZoom.Y),
	}
}

func ScreenToWorldPosition(screen sdl.Point) sdl.Point {
	return sdl.Point{
		X: int32(math32.Floor((gfx.RendererInstance.Context.CameraOffset.X + float32(screen.X)) / gfx.RendererInstance.Context.CameraTileSize.X)),
		Y: int32(math32.Floor((gfx.RendererInstance.Context.CameraOffset.Y + float32(screen.Y)) / gfx.RendererInstance.Context.CameraTileSize.Y)),
	}
}

func ScreenToWorldRect(screen sdl.Rect) sdl.Rect {
	pos := ScreenToWorldPosition(sdl.Point{X: screen.X, Y: screen.Y})
	return sdl.Rect{
		X: pos.X,
		Y: pos.Y,
		W: screen.W, H: screen.H,
	}
}

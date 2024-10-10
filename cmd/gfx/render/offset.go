package render

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func GetCameraTileSize() sdl.FPoint {
	return sdlutils.MultiplyFPoints(gameplay.TileSizeF, gfx.RendererInstance.Context.CameraZoom)
}

func GetCameraOffset() sdl.FPoint {
	return sdlutils.SubtractFPoints(
		sdlutils.AddFPoints(
			sdlutils.MultiplyFPoints(gfx.RendererInstance.Context.CameraPosition, gfx.RendererInstance.Context.CameraTileSizeF),
			sdlutils.DivideFPoint(gfx.RendererInstance.Context.CameraTileSizeF, 2),
		),
		sdlutils.DivideFPoint(sdlutils.PointToFPoint(gfx.RendererInstance.Context.WindowSize), 2),
	)
}

func WorldToScreenPosition(world sdl.Point) sdl.Point {
	return sdl.Point{
		X: int32((float32(world.X) * gfx.RendererInstance.Context.CameraTileSizeF.X) - gfx.RendererInstance.Context.CameraOffsetF.X),
		Y: int32((float32(world.Y) * gfx.RendererInstance.Context.CameraTileSizeF.Y) - gfx.RendererInstance.Context.CameraOffsetF.Y),
	}
}

func WorldToScreenPositionF(world sdl.FPoint) sdl.Point {
	return sdl.Point{
		X: int32((world.X * gfx.RendererInstance.Context.CameraTileSizeF.X) - gfx.RendererInstance.Context.CameraOffsetF.X),
		Y: int32((world.Y * gfx.RendererInstance.Context.CameraTileSizeF.Y) - gfx.RendererInstance.Context.CameraOffsetF.Y),
	}
}

func WorldToScreenRectF(world sdl.FRect) sdl.Rect {
	pos := WorldToScreenPositionF(sdl.FPoint{X: world.X, Y: world.Y})
	return sdl.Rect{
		X: pos.X,
		Y: pos.Y,
		W: int32(world.W * gfx.RendererInstance.Context.CameraZoom.X),
		H: int32(world.H * gfx.RendererInstance.Context.CameraZoom.Y),
	}
}

func ScreenToWorldPosition(screen sdl.Point) sdl.Point {
	return sdl.Point{
		X: int32((gfx.RendererInstance.Context.CameraOffsetF.X + float32(screen.X)) / gfx.RendererInstance.Context.CameraTileSizeF.X),
		Y: int32((gfx.RendererInstance.Context.CameraOffsetF.Y + float32(screen.Y)) / gfx.RendererInstance.Context.CameraTileSizeF.Y),
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

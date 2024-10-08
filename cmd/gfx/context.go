package gfx

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type RendererContext struct {
	WindowSize                sdl.Point
	MousePosition             sdl.Point
	Frame                     uint64
	CameraMode                RendererCameraMode
	CameraPosition            sdl.FPoint
	CameraPositionDestination sdl.FPoint
	CameraZoom                sdl.FPoint
	CameraZoomDestination     sdl.FPoint
	CameraTileSize            sdl.Point
	CameraTileSizeF           sdl.FPoint
	CameraOffset              sdl.Point
	CameraOffsetF             sdl.FPoint
	State                     RendererContextState
}

type RendererCameraMode uint8

const (
	RendererCameraModeCharacter = RendererCameraMode(0)
	RendererCameraModeFree      = RendererCameraMode(1)
)

type RendererContextState uint8

const (
	RendererContextStateNone    = RendererContextState(0)
	RendererContextStateActions = RendererContextState(1)
	RendererContextStateTool    = RendererContextState(2)
)

func NewRendererContext() RendererContext {
	return RendererContext{
		WindowSize:                sdl.Point{},
		MousePosition:             sdl.Point{},
		Frame:                     0,
		CameraMode:                RendererCameraModeCharacter,
		CameraPosition:            sdl.FPoint{},
		CameraPositionDestination: sdl.FPoint{},
		CameraZoom:                sdl.FPoint{X: 4, Y: 4},
		CameraZoomDestination:     sdl.FPoint{X: 4, Y: 4},
		CameraTileSize:            sdl.Point{},
		CameraTileSizeF:           sdl.FPoint{},
		CameraOffset:              sdl.Point{},
		CameraOffsetF:             sdl.FPoint{},
		State:                     RendererContextStateNone,
	}
}

func GetHalfWindowSize(context *RendererContext) sdl.Point {
	return sdl.Point{
		X: context.WindowSize.X / 2,
		Y: context.WindowSize.Y / 2,
	}
}

func GetHoveredOffset(context *RendererContext, base sdl.Rect) sdl.Point {
	return sdl.Point{
		X: int32(float32(context.MousePosition.X-base.X) / context.CameraZoom.X),
		Y: int32(float32(context.MousePosition.Y-base.Y) / context.CameraZoom.Y),
	}
}

func GetHoveredOffsetUnscaled(context *RendererContext, base sdl.Point) sdl.Point {
	return sdl.Point{
		X: context.MousePosition.X - base.X,
		Y: context.MousePosition.Y - base.Y,
	}
}

func IsHoveredOffsetSolid(texture *sdlutils.TextureWithSizeAndSurface, offset sdl.Point) bool {
	if offset.InRect(&sdl.Rect{W: texture.Base.Size.W, H: texture.Base.Size.H}) {
		hoveredColor := sdlutils.GetPixelAt(texture, offset)
		if hoveredColor != nil && sdlutils.IsColorVisible(*hoveredColor) {
			return true
		}
	}

	return false
}

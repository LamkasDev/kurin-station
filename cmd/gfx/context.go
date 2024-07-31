package gfx

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRendererContext struct {
	WindowSize                sdl.Point
	MousePosition             sdl.Point
	Frame                     uint64
	CameraMode                KurinRendererCameraMode
	CameraPosition            sdl.FPoint
	CameraPositionDestination sdl.FPoint
	CameraZoom                sdl.FPoint
	CameraZoomDestination     sdl.FPoint
	CameraTileSize            sdl.FPoint
	CameraOffset              sdl.FPoint
	State                     KurinRendererContextState
}

type KurinRendererCameraMode uint8

const KurinRendererCameraModeCharacter = KurinRendererCameraMode(0)
const KurinRendererCameraModeFree = KurinRendererCameraMode(1)

type KurinRendererContextState uint8

const KurinRendererContextStateNone = KurinRendererContextState(0)
const KurinRendererContextStateActions = KurinRendererContextState(1)
const KurinRendererContextStateTool = KurinRendererContextState(2)

func NewKurinRendererContext() KurinRendererContext {
	return KurinRendererContext{
		WindowSize:                sdl.Point{},
		MousePosition:             sdl.Point{},
		Frame:                     0,
		CameraMode:                KurinRendererCameraModeCharacter,
		CameraPosition:            sdl.FPoint{},
		CameraPositionDestination: sdl.FPoint{},
		CameraZoom:                sdl.FPoint{X: 4, Y: 4},
		CameraZoomDestination:     sdl.FPoint{X: 4, Y: 4},
		CameraTileSize:            sdl.FPoint{},
		CameraOffset:              sdl.FPoint{},
		State:                     KurinRendererContextStateNone,
	}
}

func GetHalfWindowSize(context *KurinRendererContext) sdl.Point {
	return sdl.Point{
		X: context.WindowSize.X / 2,
		Y: context.WindowSize.Y / 2,
	}
}

func GetHoveredOffset(context *KurinRendererContext, base sdl.Rect) sdl.Point {
	return sdl.Point{
		X: int32(float32(context.MousePosition.X-base.X) / context.CameraZoom.X),
		Y: int32(float32(context.MousePosition.Y-base.Y) / context.CameraZoom.Y),
	}
}

func GetHoveredOffsetUnscaled(context *KurinRendererContext, base sdl.Point) sdl.Point {
	return sdl.Point{
		X: context.MousePosition.X-base.X,
		Y: context.MousePosition.Y-base.Y,
	}
}

func IsHoveredOffsetSolid(texture sdlutils.TextureWithSizeAndSurface, offset sdl.Point) bool {
	if offset.InRect(&sdl.Rect{W: texture.Base.Size.W, H: texture.Base.Size.H}) {
		hoveredColor := sdlutils.GetPixelAt(texture, offset)
		if hoveredColor != nil && sdlutils.IsColorVisible(*hoveredColor) {
			return true
		}
	}

	return false
}

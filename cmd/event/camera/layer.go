package camera

import (
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/game/timing"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerCameraData struct{}

func NewKurinEventLayerCamera() *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadKurinEventLayerCamera,
		Process: ProcessKurinEventLayerCamera,
		Data:    &KurinEventLayerCameraData{},
	}
}

func LoadKurinEventLayerCamera(layer *event.EventLayer) error {
	return nil
}

func ProcessKurinEventLayerCamera(layer *event.EventLayer) error {
	if event.EventManagerInstance.Keyboard.InputMode {
		return nil
	}

	switch gfx.RendererInstance.Context.CameraMode {
	case gfx.KurinRendererCameraModeCharacter:
		gfx.RendererInstance.Context.CameraPosition = gameplay.GameInstance.SelectedCharacter.PositionRender
		gfx.RendererInstance.Context.CameraPositionDestination = gfx.RendererInstance.Context.CameraPosition
	case gfx.KurinRendererCameraModeFree:
		delay := float32(60)
		if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_w]; pressed {
			gfx.RendererInstance.Context.CameraPositionDestination.Y -= timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_s]; pressed {
			gfx.RendererInstance.Context.CameraPositionDestination.Y += timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_a]; pressed {
			gfx.RendererInstance.Context.CameraPositionDestination.X -= timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_d]; pressed {
			gfx.RendererInstance.Context.CameraPositionDestination.X += timing.KurinTimingGlobal.FrameTime / delay
		}

		gfx.RendererInstance.Context.CameraPosition = mathutils.LerpFPoint(gfx.RendererInstance.Context.CameraPosition, gfx.RendererInstance.Context.CameraPositionDestination, 0.4)
	}

	if event.EventManagerInstance.Mouse.PendingScroll != 0 {
		zoom := float32(event.EventManagerInstance.Mouse.PendingScroll) * float32(0.5)
		zoomDestination := mathutils.ClampFloat32(gfx.RendererInstance.Context.CameraZoomDestination.X+zoom, 1, 4)
		gfx.RendererInstance.Context.CameraZoomDestination = sdl.FPoint{
			X: zoomDestination,
			Y: zoomDestination,
		}
	}
	gfx.RendererInstance.Context.CameraZoom = mathutils.LerpFPoint(gfx.RendererInstance.Context.CameraZoom, gfx.RendererInstance.Context.CameraZoomDestination, 0.4)

	return nil
}

package camera

import (
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/game/timing"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerCameraData struct {
}

func NewKurinEventLayerCamera() *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerCamera,
		Process: ProcessKurinEventLayerCamera,
		Data: KurinEventLayerCameraData{},
	}
}

func LoadKurinEventLayerCamera(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	return nil
}

func ProcessKurinEventLayerCamera(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	if manager.Keyboard.InputMode {
		return nil
	}

	switch manager.Renderer.Context.CameraMode {
	case gfx.KurinRendererCameraModeCharacter:
		manager.Renderer.Context.CameraPosition = gameplay.KurinGameInstance.SelectedCharacter.PositionRender
		manager.Renderer.Context.CameraPositionDestination = manager.Renderer.Context.CameraPosition
	case gfx.KurinRendererCameraModeFree:
		delay := float32(60)
		if pressed := manager.Keyboard.Pressed[sdl.K_w]; pressed {
			manager.Renderer.Context.CameraPositionDestination.Y -= timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := manager.Keyboard.Pressed[sdl.K_s]; pressed {
			manager.Renderer.Context.CameraPositionDestination.Y += timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := manager.Keyboard.Pressed[sdl.K_a]; pressed {
			manager.Renderer.Context.CameraPositionDestination.X -= timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := manager.Keyboard.Pressed[sdl.K_d]; pressed {
			manager.Renderer.Context.CameraPositionDestination.X += timing.KurinTimingGlobal.FrameTime / delay
		}

		manager.Renderer.Context.CameraPosition = mathutils.LerpFPoint(manager.Renderer.Context.CameraPosition, manager.Renderer.Context.CameraPositionDestination, 0.4)
	}

	if manager.Mouse.Scroll != 0 {
		zoom := float32(manager.Mouse.Scroll) * float32(0.5)
		zoomDestination := mathutils.ClampFloat32(manager.Renderer.Context.CameraZoomDestination.X+zoom, 1, 4)
		manager.Renderer.Context.CameraZoomDestination = sdl.FPoint{
			X: zoomDestination,
			Y: zoomDestination,
		}
	}
	manager.Renderer.Context.CameraZoom = mathutils.LerpFPoint(manager.Renderer.Context.CameraZoom, manager.Renderer.Context.CameraZoomDestination, 0.4)

	return nil
}

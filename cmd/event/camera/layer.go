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

func LoadKurinEventLayerCamera(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerCamera(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	if manager.Keyboard.InputMode {
		return nil
	}

	switch manager.Renderer.RendererContext.CameraMode {
	case gfx.KurinRendererCameraModeCharacter:
		manager.Renderer.RendererContext.CameraPosition = game.SelectedCharacter.PositionRender
		manager.Renderer.RendererContext.CameraPositionDestination = manager.Renderer.RendererContext.CameraPosition
	case gfx.KurinRendererCameraModeFree:
		delay := float32(60)
		if pressed := manager.Keyboard.Pressed[sdl.K_w]; pressed {
			manager.Renderer.RendererContext.CameraPositionDestination.Y -= timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := manager.Keyboard.Pressed[sdl.K_s]; pressed {
			manager.Renderer.RendererContext.CameraPositionDestination.Y += timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := manager.Keyboard.Pressed[sdl.K_a]; pressed {
			manager.Renderer.RendererContext.CameraPositionDestination.X -= timing.KurinTimingGlobal.FrameTime / delay
		}
		if pressed := manager.Keyboard.Pressed[sdl.K_d]; pressed {
			manager.Renderer.RendererContext.CameraPositionDestination.X += timing.KurinTimingGlobal.FrameTime / delay
		}

		manager.Renderer.RendererContext.CameraPosition = mathutils.LerpFPoint(manager.Renderer.RendererContext.CameraPosition, manager.Renderer.RendererContext.CameraPositionDestination, 0.4)
	}

	if manager.Mouse.Scroll != 0 {
		zoom := float32(manager.Mouse.Scroll) * float32(0.5)
		zoomDestination := mathutils.ClampFloat32(manager.Renderer.RendererContext.CameraZoomDestination.X+zoom, 1, 4)
		manager.Renderer.RendererContext.CameraZoomDestination = sdl.FPoint{
			X: zoomDestination,
			Y: zoomDestination,
		}
	}
	manager.Renderer.RendererContext.CameraZoom = mathutils.LerpFPoint(manager.Renderer.RendererContext.CameraZoom, manager.Renderer.RendererContext.CameraZoomDestination, 0.4)

	return nil
}

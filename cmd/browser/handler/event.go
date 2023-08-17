package handler

import (
	"github.com/LamkasDev/kitsune/cmd/browser/gfx"
	"github.com/LamkasDev/kitsune/cmd/browser/life"
	"github.com/LamkasDev/kitsune/cmd/common/mathutils"
	"github.com/veandco/go-sdl2/sdl"
)

func HandleEvents(instance *life.KitsuneInstance) (bool, *error) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return false, nil
		case *sdl.WindowEvent:
			switch event.(*sdl.WindowEvent).Event {
			case sdl.WINDOWEVENT_RESIZED:
				w, h := instance.Renderer.Window.GetSize()
				instance.Renderer.WindowContext.WindowSize = sdl.Rect{
					W: w,
					H: h,
				}
			}
		case *sdl.MouseMotionEvent:
			instance.Renderer.WindowContext.MousePosition = sdl.Rect{
				X: event.(*sdl.MouseMotionEvent).X,
				Y: event.(*sdl.MouseMotionEvent).Y,
			}
		case *sdl.MouseWheelEvent:
			scroll := float32(event.(*sdl.MouseWheelEvent).Y * 32)
			instance.Context.ScrollDestination.Y = mathutils.ClampFloat32(instance.Context.ScrollDestination.Y+scroll, gfx.GetScrollbarLimit(instance.Renderer), 0)
		case *sdl.MouseButtonEvent:
			switch event.(*sdl.MouseButtonEvent).Type {
			case sdl.MOUSEBUTTONDOWN:
				break
			}
		}
	}

	return true, nil
}

func HandleEventsFrame(instance *life.KitsuneInstance) *error {
	instance.Context.ScrollPosition.Y = mathutils.Lerp(instance.Context.ScrollPosition.Y, instance.Context.ScrollDestination.Y, 0.4)
	return nil
}

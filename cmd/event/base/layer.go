package base

import (
	"errors"

	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerBaseData struct{}

func NewKurinEventLayerBase() *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadKurinEventLayerBase,
		Process: ProcessKurinEventLayerBase,
		Data:    &KurinEventLayerBaseData{},
	}
}

func LoadKurinEventLayerBase(layer *event.EventLayer) error {
	return nil
}

func ProcessKurinEventLayerBase(layer *event.EventLayer) error {
	event.EventManagerInstance.Mouse.Delta = sdl.Point{}
	for sdlEvent := sdl.PollEvent(); sdlEvent != nil; sdlEvent = sdl.PollEvent() {
		switch val := sdlEvent.(type) {
		case sdl.QuitEvent:
			err := errors.New("time to bounce")
			event.EventManagerInstance.Close = true
			return err
		case sdl.WindowEvent:
			switch val.Event {
			case sdl.WINDOWEVENT_RESIZED:
				w, h := gfx.RendererInstance.Window.GetSize()
				gfx.RendererInstance.Context.WindowSize = sdl.Point{X: w, Y: h}
			}
		case sdl.MouseMotionEvent:
			event.EventManagerInstance.Mouse.Delta = sdl.Point{
				X: val.XRel,
				Y: val.YRel,
			}
			gfx.RendererInstance.Context.MousePosition = sdl.Point{
				X: val.X,
				Y: val.Y,
			}
		case sdl.MouseWheelEvent:
			event.EventManagerInstance.Mouse.PendingScroll = val.Y
		case sdl.KeyboardEvent:
			key := val.Keysym.Sym
			switch val.Type {
			case sdl.KEYDOWN:
				event.EventManagerInstance.Keyboard.Pressed[key] = true
				event.EventManagerInstance.Keyboard.Pending = &key
			case sdl.KEYUP:
				event.EventManagerInstance.Keyboard.Pressed[key] = false
			}
		case sdl.MouseButtonEvent:
			switch val.Type {
			case sdl.MOUSEBUTTONDOWN:
				wpos := render.ScreenToWorldPosition(gfx.RendererInstance.Context.MousePosition)
				switch val.Button {
				case sdl.ButtonLeft:
					event.EventManagerInstance.Mouse.PressedLeft = true
					event.EventManagerInstance.Mouse.PendingLeft = &wpos
				case sdl.ButtonRight:
					event.EventManagerInstance.Mouse.PressedRight = true
					event.EventManagerInstance.Mouse.PendingRight = &wpos
				}
			case sdl.MOUSEBUTTONUP:
				switch val.Button {
				case sdl.ButtonLeft:
					event.EventManagerInstance.Mouse.PressedLeft = false
				case sdl.ButtonRight:
					event.EventManagerInstance.Mouse.PressedRight = false
				}
			}
		case sdl.TextInputEvent:
			event.EventManagerInstance.Keyboard.Input = val.GetText()
		}
	}
	gameplay.ProcessKurinGame()

	return nil
}

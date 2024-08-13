package base

import (
	"errors"

	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerBaseData struct {
}

func NewKurinEventLayerBase() *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerBase,
		Process: ProcessKurinEventLayerBase,
		Data: KurinEventLayerBaseData{},
	}
}

func LoadKurinEventLayerBase(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	return nil
}

func ProcessKurinEventLayerBase(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch val := event.(type) {
		case sdl.QuitEvent:
			err := errors.New("time to bounce")
			manager.Close = true
			return err
		case sdl.WindowEvent:
			switch val.Event {
			case sdl.WINDOWEVENT_RESIZED:
				w, h := manager.Renderer.Window.GetSize()
				manager.Renderer.Context.WindowSize = sdl.Point{X: w, Y: h}
			}
		case sdl.MouseMotionEvent:
			manager.Renderer.Context.MousePosition = sdl.Point{
				X: val.X,
				Y: val.Y,
			}
		case sdl.MouseWheelEvent:
			manager.Mouse.Scroll = val.Y
		case sdl.KeyboardEvent:
			key := val.Keysym.Sym
			switch val.Type {
			case sdl.KEYDOWN:
				manager.Keyboard.Pressed[key] = true
				manager.Keyboard.Pending = &key
			case sdl.KEYUP:
				manager.Keyboard.Pressed[key] = false
			}
		case sdl.MouseButtonEvent:
			switch val.Type {
			case sdl.MOUSEBUTTONDOWN:
				wpos := render.ScreenToWorldPosition(manager.Renderer, manager.Renderer.Context.MousePosition)
				switch val.Button {
				case sdl.ButtonLeft:
					manager.Mouse.PendingLeft = &wpos
				case sdl.ButtonRight:
					manager.Mouse.PendingRight = &wpos
				}
			}
		case sdl.TextInputEvent:
			manager.Keyboard.Input = val.GetText()
		}
	}
	gameplay.ProcessKurinGame()

	return nil
}

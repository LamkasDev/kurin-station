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
	}
}

func LoadKurinEventLayerBase(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	layer.Data = KurinEventLayerBaseData{}

	return nil
}

func ProcessKurinEventLayerBase(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch val := event.(type) {
		case sdl.QuitEvent:
			err := errors.New("time to bounce")
			return &err
		case sdl.WindowEvent:
			switch val.Event {
			case sdl.WINDOWEVENT_RESIZED:
				w, h := manager.Renderer.Window.GetSize()
				manager.Renderer.WindowContext.WindowSize = sdl.Point{X: w, Y: h}
			}
		case sdl.MouseMotionEvent:
			manager.Renderer.WindowContext.MousePosition = sdl.Point{
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
				wpos := render.ScreenToWorldPosition(manager.Renderer, manager.Renderer.WindowContext.MousePosition)
				switch val.Button {
				case sdl.ButtonLeft:
					manager.Mouse.PendingLeft = &wpos
					manager.Mouse.LastLeft = &wpos
				case sdl.ButtonRight:
					manager.Mouse.PendingRight = &wpos
					manager.Mouse.LastRight = &wpos
				}
			}
		case sdl.TextInputEvent:
			manager.Keyboard.Input = val.GetText()
		}
	}

	return nil
}

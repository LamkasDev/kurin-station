package event

import (
	"github.com/veandco/go-sdl2/sdl"
)

var EventManagerInstance *EventManager

type EventManager struct {
	Layers   []*EventLayer
	Mouse    Mouse
	Keyboard Keyboard
	Close    bool
}

func InitializeEventManager() error {
	EventManagerInstance = &EventManager{
		Layers:   []*EventLayer{},
		Mouse:    NewMouse(),
		Keyboard: NewKeyboard(),
	}

	return nil
}

func LoadEventManager() error {
	for _, layer := range EventManagerInstance.Layers {
		if err := layer.Load(layer); err != nil {
			return err
		}
	}

	return nil
}

func ProcessEventManager() error {
	for _, layer := range EventManagerInstance.Layers {
		if err := layer.Process(layer); err != nil {
			return err
		}
	}
	sdl.SetCursor(EventManagerInstance.Mouse.Cursors[EventManagerInstance.Mouse.Cursor])
	EventManagerInstance.Mouse.Cursor = sdl.SYSTEM_CURSOR_ARROW
	EventManagerInstance.Mouse.PendingLeft = nil
	EventManagerInstance.Mouse.PendingRight = nil
	EventManagerInstance.Mouse.PendingScroll = 0
	EventManagerInstance.Keyboard.Pending = nil
	EventManagerInstance.Keyboard.Input = ""

	return nil
}

func FreeEventManager() error {
	return nil
}

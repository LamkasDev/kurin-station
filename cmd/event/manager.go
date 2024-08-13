package event

import (
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventManager struct {
	Layers   []*KurinEventLayer
	Renderer *gfx.KurinRenderer
	Mouse    KurinMouse
	Keyboard KurinKeyboard
	Close bool
}

func NewKurinEventManager(renderer *gfx.KurinRenderer) (KurinEventManager, error) {
	manager := KurinEventManager{
		Layers:   []*KurinEventLayer{},
		Renderer: renderer,
		Mouse:    NewKurinMouse(),
		Keyboard: NewKurinKeyboard(),
	}

	return manager, nil
}

func LoadKurinEventManager(manager *KurinEventManager) error {
	for _, layer := range manager.Layers {
		if err := layer.Load(manager, layer); err != nil {
			return err
		}
	}

	return nil
}

func ProcessKurinEventManager(manager *KurinEventManager) error {
	for _, layer := range manager.Layers {
		if err := layer.Process(manager, layer); err != nil {
			return err
		}
	}
	sdl.SetCursor(manager.Mouse.Cursors[manager.Mouse.Cursor])
	manager.Mouse.Cursor = sdl.SYSTEM_CURSOR_ARROW
	manager.Mouse.PendingLeft = nil
	manager.Mouse.PendingRight = nil
	manager.Mouse.Scroll = 0
	manager.Keyboard.Pending = nil
	manager.Keyboard.Input = ""

	return nil
}

func FreeKurinEventManager(manager *KurinEventManager) error {
	return nil
}

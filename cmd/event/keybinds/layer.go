package keybinds

import (
	"os"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/serialization"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerKeybindsData struct {
}

func NewKurinEventLayerKeybinds() *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerKeybinds,
		Process: ProcessKurinEventLayerKeybinds,
		Data:    KurinEventLayerKeybindsData{},
	}
}

func LoadKurinEventLayerKeybinds(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	return nil
}

func ProcessKurinEventLayerKeybinds(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	if manager.Keyboard.Pending != nil {
		switch *manager.Keyboard.Pending {
		case sdl.K_x:
			if gameplay.KurinGameInstance.SelectedCharacter == nil {
				return nil
			}
			switch gameplay.KurinGameInstance.SelectedCharacter.ActiveHand {
			case gameplay.KurinHandLeft:
				gameplay.KurinGameInstance.SelectedCharacter.ActiveHand = gameplay.KurinHandRight
			case gameplay.KurinHandRight:
				gameplay.KurinGameInstance.SelectedCharacter.ActiveHand = gameplay.KurinHandLeft
			}
		case sdl.K_q:
			if gameplay.KurinGameInstance.SelectedCharacter == nil {
				return nil
			}
			gameplay.DropKurinItemFromCharacter(gameplay.KurinGameInstance.SelectedCharacter)
		case sdl.K_r:
			if gameplay.KurinGameInstance.SelectedCharacter == nil {
				return nil
			}
			item := gameplay.KurinGameInstance.SelectedCharacter.Inventory.Hands[gameplay.KurinGameInstance.SelectedCharacter.ActiveHand]
			if !gameplay.DropKurinItemFromCharacter(gameplay.KurinGameInstance.SelectedCharacter) {
				return nil
			}
			wpos := render.ScreenToWorldPosition(manager.Renderer, manager.Renderer.Context.MousePosition)
			force := gameplay.KurinForce{
				Item: item,
				Target: sdlutils.PointToFPointCenter(wpos),
			}
			gameplay.KurinGameInstance.ForceController.Forces[item] = &force
		case sdl.K_f:
			switch manager.Renderer.Context.CameraMode {
			case gfx.KurinRendererCameraModeCharacter:
				manager.Renderer.Context.CameraMode = gfx.KurinRendererCameraModeFree
				gameplay.KurinGameInstance.SelectedCharacter = nil
			case gfx.KurinRendererCameraModeFree:
				manager.Renderer.Context.CameraMode = gfx.KurinRendererCameraModeCharacter
				gameplay.KurinGameInstance.SelectedCharacter = gameplay.KurinGameInstance.Characters[0]
			}
		case sdl.K_s:
			if manager.Keyboard.Pressed[sdl.K_LCTRL] {
				data := serialization.EncodeKurinGame(gameplay.KurinGameInstance)
				if _, err := os.Stat(path.Join(constants.TempSavesPath, "save.dat")); err == nil {
					os.Remove(path.Join(constants.TempSavesPath, "save.dat"))
				}
				os.WriteFile(path.Join(constants.TempSavesPath, "save.dat"), data, 777)
			}
		case sdl.K_l:
			if manager.Keyboard.Pressed[sdl.K_LCTRL] {
				data, _ := os.ReadFile(path.Join(constants.TempSavesPath, "save.dat"))
				serialization.DecodeKurinGame(data, gameplay.KurinGameInstance)
			}
		default:
			return nil
		}
		manager.Keyboard.Pending = nil
	}

	return nil
}

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
	"github.com/kelindar/binary"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLayerKeybindsData struct{}

func NewEventLayerKeybinds() *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerKeybinds,
		Process: ProcessEventLayerKeybinds,
		Data:    &EventLayerKeybindsData{},
	}
}

func LoadEventLayerKeybinds(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerKeybinds(layer *event.EventLayer) error {
	if event.EventManagerInstance.Keyboard.Pending != nil {
		switch *event.EventManagerInstance.Keyboard.Pending {
		case sdl.K_x:
			if gameplay.GameInstance.SelectedCharacter == nil {
				return nil
			}
			switch gameplay.GetActiveHand(gameplay.GameInstance.SelectedCharacter) {
			case gameplay.HandLeft:
				gameplay.GetInventory(gameplay.GameInstance.SelectedCharacter).ActiveHand = gameplay.HandRight
			case gameplay.HandRight:
				gameplay.GetInventory(gameplay.GameInstance.SelectedCharacter).ActiveHand = gameplay.HandLeft
			}
		case sdl.K_q:
			if gameplay.GameInstance.SelectedCharacter == nil || gameplay.GameInstance.SelectedCharacter.Health.Dead {
				return nil
			}
			gameplay.DropItemFromCharacter(gameplay.GameInstance.SelectedCharacter)
		case sdl.K_r:
			if gameplay.GameInstance.SelectedCharacter == nil || gameplay.GameInstance.SelectedCharacter.Health.Dead {
				return nil
			}
			item := gameplay.GetHeldItem(gameplay.GameInstance.SelectedCharacter)
			if !gameplay.DropItemFromCharacter(gameplay.GameInstance.SelectedCharacter) {
				return nil
			}
			force := gameplay.NewForce(item.Transform.Position, sdlutils.PointToFPointCenter(render.ScreenToWorldPosition(gfx.RendererInstance.Context.MousePosition)), gameplay.GameInstance.SelectedCharacter.Id, item)
			gameplay.GameInstance.ForceController.Items[item] = force
		case sdl.K_f:
			switch gfx.RendererInstance.Context.CameraMode {
			case gfx.RendererCameraModeCharacter:
				gfx.RendererInstance.Context.CameraMode = gfx.RendererCameraModeFree
				gameplay.GameInstance.SelectedCharacter = nil
			case gfx.RendererCameraModeFree:
				gfx.RendererInstance.Context.CameraMode = gfx.RendererCameraModeCharacter
				gameplay.GameInstance.SelectedCharacter = gameplay.GameInstance.Map.Mobs[0]
			}
		case sdl.K_s:
			if event.EventManagerInstance.Keyboard.Pressed[sdl.K_LCTRL] {
				data := serialization.EncodeGame()
				if _, err := os.Stat(path.Join(constants.TempSavesPath, "save.dat")); err == nil {
					os.Remove(path.Join(constants.TempSavesPath, "save.dat"))
				}
				os.WriteFile(path.Join(constants.TempSavesPath, "save.dat"), data, 777)
			}
		case sdl.K_l:
			if event.EventManagerInstance.Keyboard.Pressed[sdl.K_LCTRL] {
				rawData, _ := os.ReadFile(path.Join(constants.TempSavesPath, "save.dat"))
				var data serialization.GameData
				if err := binary.Unmarshal(rawData, &data); err != nil {
					panic(err)
				}
				gameplay.CloseDialog()
				serialization.PredecodeGame(data)
				serialization.DecodeGame(data)
			}
		default:
			return nil
		}
		event.EventManagerInstance.Keyboard.Pending = nil
	}

	return nil
}

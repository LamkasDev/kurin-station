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
			switch gameplay.GameInstance.SelectedCharacter.ActiveHand {
			case gameplay.HandLeft:
				gameplay.GameInstance.SelectedCharacter.ActiveHand = gameplay.HandRight
			case gameplay.HandRight:
				gameplay.GameInstance.SelectedCharacter.ActiveHand = gameplay.HandLeft
			}
		case sdl.K_q:
			if gameplay.GameInstance.SelectedCharacter == nil {
				return nil
			}
			gameplay.DropItemFromCharacter(gameplay.GameInstance.SelectedCharacter)
		case sdl.K_r:
			if gameplay.GameInstance.SelectedCharacter == nil {
				return nil
			}
			item := gameplay.GameInstance.SelectedCharacter.Inventory.Hands[gameplay.GameInstance.SelectedCharacter.ActiveHand]
			if !gameplay.DropItemFromCharacter(gameplay.GameInstance.SelectedCharacter) {
				return nil
			}
			// this shit is trash
			force := gameplay.Force{
				Item:   item,
				Target: sdlutils.PointToFPointCenter(render.ScreenToWorldPosition(gfx.RendererInstance.Context.MousePosition)),
			}
			force.Delta = sdlutils.SubtractFPoints(force.Target, item.Transform.Position.Base)
			distance := sdlutils.GetDistanceF(item.Transform.Position.Base, force.Target)
			if distance != 0 {
				force.Delta.X /= distance * 3
				force.Delta.Y /= distance * 3
			}
			if gameplay.GameInstance.HoveredTile == nil {
				force.Target.X += force.Delta.X * 100
				force.Target.Y += force.Delta.Y * 100
			}
			gameplay.GameInstance.ForceController.Forces[item] = &force
		case sdl.K_f:
			switch gfx.RendererInstance.Context.CameraMode {
			case gfx.RendererCameraModeCharacter:
				gfx.RendererInstance.Context.CameraMode = gfx.RendererCameraModeFree
				gameplay.GameInstance.SelectedCharacter = nil
			case gfx.RendererCameraModeFree:
				gfx.RendererInstance.Context.CameraMode = gfx.RendererCameraModeCharacter
				gameplay.GameInstance.SelectedCharacter = gameplay.GameInstance.Characters[0]
			}
		case sdl.K_s:
			if event.EventManagerInstance.Keyboard.Pressed[sdl.K_LCTRL] {
				data := serialization.EncodeGame(gameplay.GameInstance)
				if _, err := os.Stat(path.Join(constants.TempSavesPath, "save.dat")); err == nil {
					os.Remove(path.Join(constants.TempSavesPath, "save.dat"))
				}
				os.WriteFile(path.Join(constants.TempSavesPath, "save.dat"), data, 777)
			}
		case sdl.K_l:
			if event.EventManagerInstance.Keyboard.Pressed[sdl.K_LCTRL] {
				data, _ := os.ReadFile(path.Join(constants.TempSavesPath, "save.dat"))
				serialization.DecodeGame(data, gameplay.GameInstance)
			}
		default:
			return nil
		}
		event.EventManagerInstance.Keyboard.Pending = nil
	}

	return nil
}

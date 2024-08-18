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

type KurinEventLayerKeybindsData struct{}

func NewKurinEventLayerKeybinds() *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadKurinEventLayerKeybinds,
		Process: ProcessKurinEventLayerKeybinds,
		Data:    &KurinEventLayerKeybindsData{},
	}
}

func LoadKurinEventLayerKeybinds(layer *event.EventLayer) error {
	return nil
}

func ProcessKurinEventLayerKeybinds(layer *event.EventLayer) error {
	if event.EventManagerInstance.Keyboard.Pending != nil {
		switch *event.EventManagerInstance.Keyboard.Pending {
		case sdl.K_x:
			if gameplay.GameInstance.SelectedCharacter == nil {
				return nil
			}
			switch gameplay.GameInstance.SelectedCharacter.ActiveHand {
			case gameplay.KurinHandLeft:
				gameplay.GameInstance.SelectedCharacter.ActiveHand = gameplay.KurinHandRight
			case gameplay.KurinHandRight:
				gameplay.GameInstance.SelectedCharacter.ActiveHand = gameplay.KurinHandLeft
			}
		case sdl.K_q:
			if gameplay.GameInstance.SelectedCharacter == nil {
				return nil
			}
			gameplay.DropKurinItemFromCharacter(gameplay.GameInstance.SelectedCharacter)
		case sdl.K_r:
			if gameplay.GameInstance.SelectedCharacter == nil {
				return nil
			}
			item := gameplay.GameInstance.SelectedCharacter.Inventory.Hands[gameplay.GameInstance.SelectedCharacter.ActiveHand]
			if !gameplay.DropKurinItemFromCharacter(gameplay.GameInstance.SelectedCharacter) {
				return nil
			}
			// this shit is trash
			position := render.ScreenToWorldPosition(gfx.RendererInstance.Context.MousePosition)
			force := gameplay.KurinForce{
				Item:   item,
				Target: sdlutils.PointToFPointCenter(position),
			}
			force.Delta = sdlutils.SubtractFPoints(force.Target, item.Transform.Position.Base)
			distance := sdlutils.GetDistanceF(item.Transform.Position.Base, force.Target)
			if distance != 0 {
				force.Delta.X /= distance * 3
				force.Delta.Y /= distance * 3
			}
			if gameplay.GetKurinTileAt(&gameplay.GameInstance.Map, sdlutils.Vector3{Base: position, Z: 0}) == nil {
				force.Target.X += force.Delta.X * 100
				force.Target.Y += force.Delta.Y * 100
			}
			gameplay.GameInstance.ForceController.Forces[item] = &force
		case sdl.K_f:
			switch gfx.RendererInstance.Context.CameraMode {
			case gfx.KurinRendererCameraModeCharacter:
				gfx.RendererInstance.Context.CameraMode = gfx.KurinRendererCameraModeFree
				gameplay.GameInstance.SelectedCharacter = nil
			case gfx.KurinRendererCameraModeFree:
				gfx.RendererInstance.Context.CameraMode = gfx.KurinRendererCameraModeCharacter
				gameplay.GameInstance.SelectedCharacter = gameplay.GameInstance.Characters[0]
			}
		case sdl.K_s:
			if event.EventManagerInstance.Keyboard.Pressed[sdl.K_LCTRL] {
				data := serialization.EncodeKurinGame(gameplay.GameInstance)
				if _, err := os.Stat(path.Join(constants.TempSavesPath, "save.dat")); err == nil {
					os.Remove(path.Join(constants.TempSavesPath, "save.dat"))
				}
				os.WriteFile(path.Join(constants.TempSavesPath, "save.dat"), data, 777)
			}
		case sdl.K_l:
			if event.EventManagerInstance.Keyboard.Pressed[sdl.K_LCTRL] {
				data, _ := os.ReadFile(path.Join(constants.TempSavesPath, "save.dat"))
				serialization.DecodeKurinGame(data, gameplay.GameInstance)
			}
		default:
			return nil
		}
		event.EventManagerInstance.Keyboard.Pending = nil
	}

	return nil
}

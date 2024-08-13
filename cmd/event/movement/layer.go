package movement

import (
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/game/timing"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const KurinMovementDelay = float32(120)

type KurinEventLayerMovementData struct {
}

func NewKurinEventLayerMovement() *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerMovement,
		Process: ProcessKurinEventLayerMovement,
		Data:    KurinEventLayerMovementData{},
	}
}

func LoadKurinEventLayerMovement(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	return nil
}

func ProcessKurinEventLayerMovement(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	for _, character := range gameplay.KurinGameInstance.Characters {
		character.PositionRender = mathutils.LerpFPoint(character.PositionRender, sdlutils.PointToFPoint(character.Position.Base), 0.2)
	}
	if manager.Renderer.Context.CameraMode != gfx.KurinRendererCameraModeCharacter || manager.Keyboard.InputMode {
		return nil
	}

	gameplay.KurinGameInstance.SelectedCharacter.Moving = false
	if pressed := manager.Keyboard.Pressed[sdl.K_w]; pressed {
		gameplay.KurinGameInstance.SelectedCharacter.Movement.Y -= timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		gameplay.KurinGameInstance.SelectedCharacter.Moving = true
	} else if pressed := manager.Keyboard.Pressed[sdl.K_s]; pressed {
		gameplay.KurinGameInstance.SelectedCharacter.Movement.Y += timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		gameplay.KurinGameInstance.SelectedCharacter.Moving = true
	} else {
		gameplay.KurinGameInstance.SelectedCharacter.Movement.Y = 0
	}
	if pressed := manager.Keyboard.Pressed[sdl.K_a]; pressed {
		gameplay.KurinGameInstance.SelectedCharacter.Movement.X -= timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		gameplay.KurinGameInstance.SelectedCharacter.Moving = true
	} else if pressed := manager.Keyboard.Pressed[sdl.K_d]; pressed {
		gameplay.KurinGameInstance.SelectedCharacter.Movement.X += timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		gameplay.KurinGameInstance.SelectedCharacter.Moving = true
	} else {
		gameplay.KurinGameInstance.SelectedCharacter.Movement.X = 0
	}

	if gameplay.KurinGameInstance.SelectedCharacter.Moving {
		gameplay.KurinGameInstance.SelectedCharacter.Direction = gameplay.GetFacingDirection(sdl.FPoint{}, gameplay.KurinGameInstance.SelectedCharacter.Movement)
	}
	direction := gameplay.KurinGameInstance.SelectedCharacter.Direction

	if gameplay.KurinGameInstance.SelectedCharacter.Movement.Y >= 1 || gameplay.KurinGameInstance.SelectedCharacter.Movement.Y <= -1 {
		position := sdlutils.Vector3{
			Base: sdl.Point{
				X: gameplay.KurinGameInstance.SelectedCharacter.Position.Base.X,
				Y: gameplay.KurinGameInstance.SelectedCharacter.Position.Base.Y + int32(gameplay.KurinGameInstance.SelectedCharacter.Movement.Y),
			},
			Z: gameplay.KurinGameInstance.SelectedCharacter.Position.Z,
		}
		gameplay.KurinGameInstance.SelectedCharacter.Movement.Y = 0
		gameplay.MoveKurinCharacter(gameplay.KurinGameInstance.SelectedCharacter, position)
		gameplay.KurinGameInstance.SelectedCharacter.Direction = direction
	}

	if gameplay.KurinGameInstance.SelectedCharacter.Movement.X >= 1 || gameplay.KurinGameInstance.SelectedCharacter.Movement.X <= -1 {
		position := sdlutils.Vector3{
			Base: sdl.Point{
				X: gameplay.KurinGameInstance.SelectedCharacter.Position.Base.X + int32(gameplay.KurinGameInstance.SelectedCharacter.Movement.X),
				Y: gameplay.KurinGameInstance.SelectedCharacter.Position.Base.Y,
			},
			Z: gameplay.KurinGameInstance.SelectedCharacter.Position.Z,
		}
		gameplay.KurinGameInstance.SelectedCharacter.Movement.X = 0
		gameplay.MoveKurinCharacter(gameplay.KurinGameInstance.SelectedCharacter, position)
		gameplay.KurinGameInstance.SelectedCharacter.Direction = direction
	}

	return nil
}

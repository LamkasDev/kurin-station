package movement

import (
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/game/timing"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const KurinMovementDelay = float32(120)

type KurinEventLayerMovementData struct{}

func NewKurinEventLayerMovement() *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadKurinEventLayerMovement,
		Process: ProcessKurinEventLayerMovement,
		Data:    &KurinEventLayerMovementData{},
	}
}

func LoadKurinEventLayerMovement(layer *event.EventLayer) error {
	return nil
}

func ProcessKurinEventLayerMovement(layer *event.EventLayer) error {
	for _, character := range gameplay.GameInstance.Characters {
		character.PositionRender = mathutils.LerpFPoint(character.PositionRender, sdlutils.PointToFPoint(character.Position.Base), 0.2)
	}
	if gfx.RendererInstance.Context.CameraMode != gfx.KurinRendererCameraModeCharacter || event.EventManagerInstance.Keyboard.InputMode {
		return nil
	}

	gameplay.GameInstance.SelectedCharacter.Moving = false
	if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_w]; pressed {
		gameplay.GameInstance.SelectedCharacter.Movement.Y -= timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		gameplay.GameInstance.SelectedCharacter.Moving = true
	} else if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_s]; pressed {
		gameplay.GameInstance.SelectedCharacter.Movement.Y += timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		gameplay.GameInstance.SelectedCharacter.Moving = true
	} else {
		gameplay.GameInstance.SelectedCharacter.Movement.Y = 0
	}
	if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_a]; pressed {
		gameplay.GameInstance.SelectedCharacter.Movement.X -= timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		gameplay.GameInstance.SelectedCharacter.Moving = true
	} else if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_d]; pressed {
		gameplay.GameInstance.SelectedCharacter.Movement.X += timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		gameplay.GameInstance.SelectedCharacter.Moving = true
	} else {
		gameplay.GameInstance.SelectedCharacter.Movement.X = 0
	}

	if gameplay.GameInstance.SelectedCharacter.Moving {
		gameplay.GameInstance.SelectedCharacter.Direction = common.GetFacingDirectionF(sdl.FPoint{}, gameplay.GameInstance.SelectedCharacter.Movement)
	}
	direction := gameplay.GameInstance.SelectedCharacter.Direction

	if gameplay.GameInstance.SelectedCharacter.Movement.Y >= 1 || gameplay.GameInstance.SelectedCharacter.Movement.Y <= -1 {
		position := sdlutils.Vector3{
			Base: sdl.Point{
				X: gameplay.GameInstance.SelectedCharacter.Position.Base.X,
				Y: gameplay.GameInstance.SelectedCharacter.Position.Base.Y + int32(gameplay.GameInstance.SelectedCharacter.Movement.Y),
			},
			Z: gameplay.GameInstance.SelectedCharacter.Position.Z,
		}
		gameplay.GameInstance.SelectedCharacter.Movement.Y = 0
		gameplay.MoveKurinCharacter(gameplay.GameInstance.SelectedCharacter, position)
		gameplay.GameInstance.SelectedCharacter.Direction = direction
	}

	if gameplay.GameInstance.SelectedCharacter.Movement.X >= 1 || gameplay.GameInstance.SelectedCharacter.Movement.X <= -1 {
		position := sdlutils.Vector3{
			Base: sdl.Point{
				X: gameplay.GameInstance.SelectedCharacter.Position.Base.X + int32(gameplay.GameInstance.SelectedCharacter.Movement.X),
				Y: gameplay.GameInstance.SelectedCharacter.Position.Base.Y,
			},
			Z: gameplay.GameInstance.SelectedCharacter.Position.Z,
		}
		gameplay.GameInstance.SelectedCharacter.Movement.X = 0
		gameplay.MoveKurinCharacter(gameplay.GameInstance.SelectedCharacter, position)
		gameplay.GameInstance.SelectedCharacter.Direction = direction
	}

	return nil
}

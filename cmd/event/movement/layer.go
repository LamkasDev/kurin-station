package movement

import (
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLayerMovementData struct{}

func NewEventLayerMovement() *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerMovement,
		Process: ProcessEventLayerMovement,
		Data:    &EventLayerMovementData{},
	}
}

func LoadEventLayerMovement(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerMovement(layer *event.EventLayer) error {
	for _, character := range gameplay.GameInstance.Characters {
		character.PositionRender = mathutils.LerpFPoint(character.PositionRender, sdlutils.PointToFPoint(character.Position.Base), 0.2)
	}
	if gfx.RendererInstance.Context.CameraMode != gfx.RendererCameraModeCharacter || event.EventManagerInstance.Keyboard.InputMode {
		return nil
	}

	gameplay.GameInstance.SelectedCharacter.Movement = sdl.Point{}
	if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_w]; pressed {
		gameplay.GameInstance.SelectedCharacter.Movement.Y = -1
	} else if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_s]; pressed {
		gameplay.GameInstance.SelectedCharacter.Movement.Y = 1
	} else {
		gameplay.GameInstance.SelectedCharacter.Movement.Y = 0
	}
	if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_a]; pressed {
		gameplay.GameInstance.SelectedCharacter.Movement.X = -1
	} else if pressed := event.EventManagerInstance.Keyboard.Pressed[sdl.K_d]; pressed {
		gameplay.GameInstance.SelectedCharacter.Movement.X = 1
	} else {
		gameplay.GameInstance.SelectedCharacter.Movement.X = 0
	}

	if !sdlutils.IsPointZero(gameplay.GameInstance.SelectedCharacter.Movement) {
		gameplay.GameInstance.SelectedCharacter.Direction = common.GetFacingDirection(sdl.Point{}, gameplay.GameInstance.SelectedCharacter.Movement)
		gameplay.GameInstance.SelectedCharacter.MovementTicks++
	} else {
		gameplay.GameInstance.SelectedCharacter.MovementTicks = 0
	}

	if gameplay.GameInstance.SelectedCharacter.MovementTicks >= gameplay.CharacterMovementTicks {
		position := sdlutils.Vector3{Base: sdlutils.AddPoints(gameplay.GameInstance.SelectedCharacter.Position.Base, gameplay.GameInstance.SelectedCharacter.Movement)}
		gameplay.GameInstance.SelectedCharacter.MovementTicks = 0
		gameplay.MoveCharacter(gameplay.GameInstance.SelectedCharacter, position)
	}

	return nil
}

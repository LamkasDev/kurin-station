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

func LoadKurinEventLayerMovement(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerMovement(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	for _, character := range game.Characters {
		character.PositionRender = mathutils.LerpFPoint(character.PositionRender, sdlutils.PointToFPoint(character.Position.Base), 0.2)
	}
	if manager.Renderer.RendererContext.CameraMode != gfx.KurinRendererCameraModeCharacter || manager.Keyboard.InputMode {
		return nil
	}

	game.SelectedCharacter.Moving = false
	if pressed := manager.Keyboard.Pressed[sdl.K_w]; pressed {
		game.SelectedCharacter.Movement.Y -= timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		game.SelectedCharacter.Moving = true
	} else if pressed := manager.Keyboard.Pressed[sdl.K_s]; pressed {
		game.SelectedCharacter.Movement.Y += timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		game.SelectedCharacter.Moving = true
	} else {
		game.SelectedCharacter.Movement.Y = 0
	}
	if pressed := manager.Keyboard.Pressed[sdl.K_a]; pressed {
		game.SelectedCharacter.Movement.X -= timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		game.SelectedCharacter.Moving = true
	} else if pressed := manager.Keyboard.Pressed[sdl.K_d]; pressed {
		game.SelectedCharacter.Movement.X += timing.KurinTimingGlobal.FrameTime / KurinMovementDelay
		game.SelectedCharacter.Moving = true
	} else {
		game.SelectedCharacter.Movement.X = 0
	}

	if game.SelectedCharacter.Moving {
		game.SelectedCharacter.Direction = gameplay.GetFacingDirection(sdl.FPoint{}, game.SelectedCharacter.Movement)
	}
	direction := game.SelectedCharacter.Direction

	if game.SelectedCharacter.Movement.Y >= 1 || game.SelectedCharacter.Movement.Y <= -1 {
		position := sdlutils.Vector3{
			Base: sdl.Point{
				X: game.SelectedCharacter.Position.Base.X,
				Y: game.SelectedCharacter.Position.Base.Y + int32(game.SelectedCharacter.Movement.Y),
			},
			Z: game.SelectedCharacter.Position.Z,
		}
		game.SelectedCharacter.Movement.Y = 0
		gameplay.MoveKurinCharacter(game.SelectedCharacter, position)
		game.SelectedCharacter.Direction = direction
	}

	if game.SelectedCharacter.Movement.X >= 1 || game.SelectedCharacter.Movement.X <= -1 {
		position := sdlutils.Vector3{
			Base: sdl.Point{
				X: game.SelectedCharacter.Position.Base.X + int32(game.SelectedCharacter.Movement.X),
				Y: game.SelectedCharacter.Position.Base.Y,
			},
			Z: game.SelectedCharacter.Position.Z,
		}
		game.SelectedCharacter.Movement.X = 0
		gameplay.MoveKurinCharacter(game.SelectedCharacter, position)
		game.SelectedCharacter.Direction = direction
	}

	return nil
}

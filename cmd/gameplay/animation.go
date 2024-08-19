package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

type Animation struct {
	Type  string
	Step  int
	Ticks int32
}

func NewAnimation(animationType string) *Animation {
	return &Animation{
		Type:  animationType,
		Step:  -1,
		Ticks: 0,
	}
}

func GetAnimationOffset(character *Character) sdl.FPoint {
	if character.AnimationController.Direction {
		switch character.Direction {
		case common.DirectionNorth:
			return sdl.FPoint{
				X: -character.AnimationController.Position.X,
				Y: -character.AnimationController.Position.Y,
			}
		case common.DirectionSouth:
			return character.AnimationController.Position
		case common.DirectionEast:
			return sdl.FPoint{
				X: character.AnimationController.Position.Y,
				Y: character.AnimationController.Position.X,
			}
		case common.DirectionWest:
			return sdl.FPoint{
				X: -character.AnimationController.Position.Y,
				Y: -character.AnimationController.Position.X,
			}
		}
	}

	return character.AnimationController.Position
}

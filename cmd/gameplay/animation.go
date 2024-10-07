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

func GetAnimationOffset(mob *Mob) sdl.FPoint {
	if mob.AnimationController.Direction {
		switch mob.Direction {
		case common.DirectionNorth:
			return sdl.FPoint{
				X: -mob.AnimationController.Position.X,
				Y: -mob.AnimationController.Position.Y,
			}
		case common.DirectionSouth:
			return mob.AnimationController.Position
		case common.DirectionEast:
			return sdl.FPoint{
				X: mob.AnimationController.Position.Y,
				Y: mob.AnimationController.Position.X,
			}
		case common.DirectionWest:
			return sdl.FPoint{
				X: -mob.AnimationController.Position.Y,
				Y: -mob.AnimationController.Position.X,
			}
		}
	}

	return mob.AnimationController.Position
}

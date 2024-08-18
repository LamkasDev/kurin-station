package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinAnimation struct {
	Type  string
	Step  int
	Ticks int32
}

func NewKurinAnimation(animationType string) *KurinAnimation {
	return &KurinAnimation{
		Type:  animationType,
		Step:  -1,
		Ticks: 0,
	}
}

func GetAnimationOffset(character *KurinCharacter) sdl.FPoint {
	if character.AnimationController.Direction {
		switch character.Direction {
		case common.KurinDirectionNorth:
			return sdl.FPoint{
				X: -character.AnimationController.Position.X,
				Y: -character.AnimationController.Position.Y,
			}
		case common.KurinDirectionSouth:
			return character.AnimationController.Position
		case common.KurinDirectionEast:
			return sdl.FPoint{
				X: character.AnimationController.Position.Y,
				Y: character.AnimationController.Position.X,
			}
		case common.KurinDirectionWest:
			return sdl.FPoint{
				X: -character.AnimationController.Position.Y,
				Y: -character.AnimationController.Position.X,
			}
		}
	}

	return character.AnimationController.Position
}

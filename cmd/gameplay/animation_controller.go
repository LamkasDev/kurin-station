package gameplay

import (
	"github.com/veandco/go-sdl2/sdl"
)

type KurinAnimationController struct {
	Animation     *KurinAnimation
	Position      sdl.FPoint
	PositionShift sdl.FPoint
	Direction     bool
}

func NewKurinAnimationController() KurinAnimationController {
	return KurinAnimationController{
		Animation:     nil,
		Position:      sdl.FPoint{},
		PositionShift: sdl.FPoint{},
		Direction:     false,
	}
}

func PlayKurinCharacterAnimation(character *KurinCharacter, atype string) {
	if character.AnimationController.Animation != nil {
		EndKurinCharacterAnimation(character)
	}
	character.AnimationController.Animation = NewKurinAnimation(atype)
}

func EndKurinCharacterAnimation(character *KurinCharacter) {
	character.AnimationController.Animation = nil
	character.AnimationController.Position = sdl.FPoint{}
	character.AnimationController.PositionShift = sdl.FPoint{}
}

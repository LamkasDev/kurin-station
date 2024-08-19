package gameplay

import (
	"github.com/veandco/go-sdl2/sdl"
)

type AnimationController struct {
	Animation     *Animation
	Position      sdl.FPoint
	PositionShift sdl.FPoint
	Direction     bool
}

func NewAnimationController() AnimationController {
	return AnimationController{
		Animation:     nil,
		Position:      sdl.FPoint{},
		PositionShift: sdl.FPoint{},
		Direction:     false,
	}
}

func PlayCharacterAnimation(character *Character, atype string) {
	if character.AnimationController.Animation != nil {
		EndCharacterAnimation(character)
	}
	character.AnimationController.Animation = NewAnimation(atype)
}

func EndCharacterAnimation(character *Character) {
	character.AnimationController.Animation = nil
	character.AnimationController.Position = sdl.FPoint{}
	character.AnimationController.PositionShift = sdl.FPoint{}
}

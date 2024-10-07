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

func PlayMobAnimation(mob *Mob, atype string) {
	if mob.AnimationController.Animation != nil {
		EndMobAnimation(mob)
	}
	mob.AnimationController.Animation = NewAnimation(atype)
}

func EndMobAnimation(mob *Mob) {
	mob.AnimationController.Animation = nil
	mob.AnimationController.Position = sdl.FPoint{}
	mob.AnimationController.PositionShift = sdl.FPoint{}
}

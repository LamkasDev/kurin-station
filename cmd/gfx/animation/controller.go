package animation

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
)

func ProcessMobAnimation(mob *gameplay.Mob, template *templates.AnimationTemplate) {
	mob.AnimationController.Animation.Ticks++
	mob.AnimationController.Position = sdlutils.AddFPoints(mob.AnimationController.Position, mob.AnimationController.PositionShift)

	step := template.Steps[mob.AnimationController.Animation.Step]
	if mob.AnimationController.Animation.Ticks >= step.Ticks {
		AdvanceMobAnimation(mob, template)
	}
}

func AdvanceMobAnimation(mob *gameplay.Mob, template *templates.AnimationTemplate) {
	mob.AnimationController.Animation.Ticks = 0
	mob.AnimationController.Animation.Step++
	if mob.AnimationController.Animation.Step >= len(template.Steps) {
		gameplay.EndMobAnimation(mob)
		return
	}

	step := template.Steps[mob.AnimationController.Animation.Step]
	mob.AnimationController.PositionShift = sdlutils.DivideFPoint(step.Offset, float32(step.Ticks))
	mob.AnimationController.Direction = *step.Direction
}

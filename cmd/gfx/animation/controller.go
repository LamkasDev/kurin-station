package animation

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
)

func ProcessKurinCharacterAnimation(character *gameplay.KurinCharacter, template *templates.KurinAnimationTemplate) {
	character.AnimationController.Animation.Ticks++
	character.AnimationController.Position = sdlutils.AddFPoints(character.AnimationController.Position, character.AnimationController.PositionShift)

	step := template.Steps[character.AnimationController.Animation.Step]
	if character.AnimationController.Animation.Ticks >= step.Ticks {
		AdvanceKurinCharacterAnimation(character, template)
	}
}

func AdvanceKurinCharacterAnimation(character *gameplay.KurinCharacter, template *templates.KurinAnimationTemplate) {
	character.AnimationController.Animation.Ticks = 0
	character.AnimationController.Animation.Step++
	if character.AnimationController.Animation.Step >= len(template.Steps) {
		gameplay.EndKurinCharacterAnimation(character)
		return
	}

	step := template.Steps[character.AnimationController.Animation.Step]
	character.AnimationController.PositionShift = sdlutils.DivideFPointByFloat(step.Offset, float32(step.Ticks))
	character.AnimationController.Direction = *step.Direction
}

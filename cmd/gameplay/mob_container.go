package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

var MobContainer = map[string]*MobTemplate{}

func RegisterMobs() {
	MobContainer["character"] = NewMobTemplateCharacter()
	MobContainer["cat"] = NewMobTemplateCat()
	MobContainer["tarantula"] = NewMobTemplateTarantula()
}

func NewMob(mobType string, faction Faction) *Mob {
	mob := &Mob{
		Id:                  GetNextId(),
		Type:                mobType,
		Gender:              DefaultGender,
		Faction:             faction,
		Position:            sdlutils.Vector3{},
		Direction:           common.DirectionEast,
		Fatigue:             0,
		Health:              NewHealth(),
		PositionRender:      sdl.FPoint{},
		Movement:            sdl.Point{},
		MovementTicks:       0,
		Thinktree:           NewThinktree(),
		AnimationController: NewAnimationController(),
		Template:            MobContainer[mobType],
	}
	mob.JobTracker = NewJobTracker(mob)
	mob.Template.Initialize(mob)

	return mob
}

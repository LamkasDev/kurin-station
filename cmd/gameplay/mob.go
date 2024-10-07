package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	DefaultGender    = "f"
	MobMovementTicks = 10
)

type Mob struct {
	Id        uint32
	Type      string
	Gender    string
	Faction   Faction
	Position  sdlutils.Vector3
	Direction common.Direction
	Fatigue   int32

	PositionRender      sdl.FPoint
	Movement            sdl.Point
	MovementTicks       uint8
	JobTracker          *JobTracker
	Thinktree           Thinktree
	AnimationController AnimationController

	Template *MobTemplate
	Data     interface{}
}

func ProcessMob(mob *Mob) {
	if mob.Fatigue > 0 {
		mob.Fatigue--
	}
	if !ProcessJobTracker(mob.JobTracker) {
		ProcessThinktree(mob)
	}
}

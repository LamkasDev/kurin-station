package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
	"robpike.io/filter"
)

const (
	DefaultGender    = "f"
	MobMovementTicks = 10
)

type Mob struct {
	Id         uint32
	Type       string
	Gender     string
	Faction    Faction
	Position   sdlutils.Vector3
	Direction  common.Direction
	Fatigue    int32
	Health     Health
	JobTracker *JobTracker

	PositionRender      sdl.FPoint
	Movement            sdl.Point
	MovementTicks       uint8
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

func GetMobsOnTile(kmap *Map, tile *Tile) []*Mob {
	return filter.Choose(GameInstance.Map.Mobs, func(mob *Mob) bool {
		return sdlutils.CompareVector3(mob.Position, tile.Position)
	}).([]*Mob)
}

func GetMobsOnTileWithout(kmap *Map, tile *Tile, id uint32) []*Mob {
	return filter.Choose(GameInstance.Map.Mobs, func(mob *Mob) bool {
		return sdlutils.CompareVector3(mob.Position, tile.Position) && mob.Id != id
	}).([]*Mob)
}

func GetMobDescription(mob *Mob) string {
	return mob.Type
}

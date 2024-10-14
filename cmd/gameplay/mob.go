package gameplay

import (
	"slices"
	"sort"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
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

func FindMobs(kmap *Map, predicate func(mob *Mob) bool) []*Mob {
	return slices.Collect(func(yield func(*Mob) bool) {
		for _, mob := range kmap.Mobs {
			if predicate(mob) {
				if !yield(mob) {
					return // triggered in "break"
				}
			}
		}
	})
}

func FindMob(kmap *Map, predicate func(mob *Mob) bool) *Mob {
	if i := slices.IndexFunc(kmap.Mobs, predicate); i >= 0 {
		return kmap.Mobs[i]
	}

	return nil
}

// TODO: z-levels
func FindClosestMob(kmap *Map, position sdlutils.Vector3, predicate func(mob *Mob) bool) *Mob {
	mobs := FindMobs(kmap, predicate)
	sort.Slice(mobs, func(i, j int) bool {
		id := sdlutils.GetDistanceSimple(position.Base, mobs[i].Position.Base)
		jd := sdlutils.GetDistanceSimple(position.Base, mobs[j].Position.Base)
		return id < jd
	})
	if len(mobs) == 0 {
		return nil
	}

	return mobs[0]
}

func GetMobsOnTile(kmap *Map, tile *Tile) []*Mob {
	return FindMobs(kmap, func(mob *Mob) bool {
		return sdlutils.CompareVector3(mob.Position, tile.Position) && !mob.Health.Dead
	})
}

func GetMobsOnTileWithout(kmap *Map, tile *Tile, id uint32) []*Mob {
	return FindMobs(kmap, func(mob *Mob) bool {
		return sdlutils.CompareVector3(mob.Position, tile.Position) && mob.Id != id
	})
}

func GetMobDescription(mob *Mob) string {
	return mob.Type
}

func IsMobAggro(mob *Mob) bool {
	return mob.Health.LastDamageTicks != 0 && GameInstance.Ticks <= mob.Health.LastDamageTicks+600
}

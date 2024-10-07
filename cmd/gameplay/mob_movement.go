package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

func TeleportMobRandom(mob *Mob) {
	for {
		position := GetRandomMapPosition(&GameInstance.Map, GameInstance.Map.BaseZ)
		if MoveMob(mob, position) {
			mob.PositionRender = sdlutils.PointToFPoint(position.Base)
			break
		}
	}
}

func MoveMobDirection(mob *Mob, direction common.Direction) bool {
	return MoveMob(mob, common.GetPositionInDirectionV(mob.Position, direction))
}

func MoveMob(mob *Mob, position sdlutils.Vector3) bool {
	if CanEnterMapPosition(&GameInstance.Map, position) == EnteranceStatusNo {
		return false
	}
	TurnMobTo(mob, position.Base)
	tile := GetTileAt(&GameInstance.Map, position)
	object := GetObjectAtTile(tile)
	if object != nil {
		switch data := object.Data.(type) {
		case *ObjectAirlockData:
			if !data.Open {
				object.Template.OnInteraction(object, nil)
				return false
			}
		}
	}
	mob.Position = position

	return true
}

func TurnMobTo(mob *Mob, position sdl.Point) {
	mob.Direction = common.GetFacingDirection(mob.Position.Base, position)
}

func FollowPath(mob *Mob, path *Path) bool {
	if path.Index == len(path.Nodes) {
		return true
	}
	mob.MovementTicks++
	if mob.MovementTicks >= MobMovementTicks {
		node := path.Nodes[path.Index]
		mob.MovementTicks = 0
		if node.Position.Z != mob.Position.Z {
			return false
		}
		if !MoveMob(mob, node.Position) {
			return false
		}
		path.Index++
	}

	return false
}

package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
)

type MobData struct {
	Id         uint32
	Type       string
	Gender     string
	Faction    gameplay.Faction
	Position   sdlutils.Vector3
	Direction  common.Direction
	Fatigue    int32
	ActiveHand gameplay.Hand
	Inventory  InventoryData
	JobTracker JobTrackerData
}

func EncodeMob(mob *gameplay.Mob) MobData {
	data := MobData{
		Id:        mob.Id,
		Type:      mob.Type,
		Gender:    mob.Gender,
		Position:  mob.Position,
		Direction: mob.Direction,
		Fatigue:   mob.Fatigue,
	}

	return data
}

func DecodeMob(data MobData) *gameplay.Mob {
	mob := gameplay.NewMob(data.Type, data.Faction)
	mob.Id = data.Id
	mob.Gender = data.Gender
	mob.Position = data.Position
	mob.PositionRender = sdlutils.PointToFPoint(data.Position.Base)
	mob.Direction = data.Direction
	mob.Fatigue = data.Fatigue

	return mob
}

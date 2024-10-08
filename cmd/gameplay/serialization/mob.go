package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/kelindar/binary"
)

type MobData struct {
	Id         uint32
	Type       string
	Gender     string
	Faction    gameplay.Faction
	Position   sdlutils.Vector3
	Direction  common.Direction
	Fatigue    int32
	JobTracker JobTrackerData
	Data       []byte
}

func EncodeMob(mob *gameplay.Mob) MobData {
	data := MobData{
		Id:         mob.Id,
		Type:       mob.Type,
		Gender:     mob.Gender,
		Position:   mob.Position,
		Direction:  mob.Direction,
		Fatigue:    mob.Fatigue,
		JobTracker: EncodeJobTracker(mob.JobTracker),
	}
	switch mob.Data.(type) {
	case *gameplay.MobCharacterData:
		mobData, _ := binary.Marshal(EncodeCharacterData(mob))
		data.Data = mobData
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
	DecodeJobTracker(data.JobTracker, mob)
	switch mob.Data.(type) {
	case *gameplay.MobCharacterData:
		var mobData CharacterData
		binary.Unmarshal(data.Data, &mobData)
		DecodeCharacterData(mobData, mob)
	}

	return mob
}

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
	Health     HealthData
	JobTracker JobTrackerData
	Data       []byte
	ExtraData  []byte
}

func EncodeMob(mob *gameplay.Mob) MobData {
	data := MobData{
		Id:         mob.Id,
		Type:       mob.Type,
		Gender:     mob.Gender,
		Position:   mob.Position,
		Direction:  mob.Direction,
		Fatigue:    mob.Fatigue,
		Health:     EncodeHealthData(mob),
		JobTracker: EncodeJobTracker(mob.JobTracker),
	}
	if mob.Data != nil {
		data.Data = mob.Template.EncodeData(mob)
	}
	switch mob.Data.(type) {
	case *gameplay.MobCharacterData:
		extraData, _ := binary.Marshal(EncodeCharacterData(mob))
		data.ExtraData = extraData
	}

	return data
}

func PredecodeMob(kmap *gameplay.Map, data MobData) *gameplay.Mob {
	mob := gameplay.NewMob(data.Type, data.Faction)
	mob.Id = data.Id
	mob.Gender = data.Gender
	mob.Position = data.Position
	mob.PositionRender = sdlutils.PointToFPoint(data.Position.Base)
	mob.Direction = data.Direction
	mob.Fatigue = data.Fatigue
	DecodeHealthData(data.Health, mob)
	DecodeJobTracker(kmap, data.JobTracker, mob)
	if data.Data != nil {
		mob.Template.DecodeData(mob, data.Data)
	}
	switch mob.Data.(type) {
	case *gameplay.MobCharacterData:
		var extraData CharacterData
		binary.Unmarshal(data.ExtraData, &extraData)
		DecodeCharacterData(extraData, mob)
	}
	kmap.Mobs = append(kmap.Mobs, mob)

	return mob
}

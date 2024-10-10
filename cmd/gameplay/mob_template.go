package gameplay

import (
	"github.com/kelindar/binary"
)

type MobTemplate struct {
	Type       string
	Initialize MobInitialize
	Process    MobProcess
	OnDeath    MobOnDeath
	EncodeData MobEncodeData
	DecodeData MobDecodeData
}

type (
	MobInitialize func(mob *Mob)
	MobProcess    func(mob *Mob)
	MobOnDeath    func(mob *Mob)
	MobEncodeData func(mob *Mob) []byte
	MobDecodeData func(mob *Mob, data []byte)
)

func NewMobTemplateRaw[D any](mobType string) *MobTemplate {
	return &MobTemplate{
		Type: mobType,
		Initialize: func(mob *Mob) {
		},
		Process: ProcessMob,
		OnDeath: func(mob *Mob) {},
		EncodeData: func(mob *Mob) []byte {
			if mob.Data == nil {
				return []byte{}
			}

			mobData := mob.Data.(D)
			data, _ := binary.Marshal(&mobData)
			return data
		},
		DecodeData: func(mob *Mob, data []byte) {
			if len(data) == 0 {
				return
			}

			var mobData D
			binary.Unmarshal(data, &mobData)
			mob.Data = mobData
		},
	}
}

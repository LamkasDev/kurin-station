package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
)

type KurinCharacterData struct {
	Id         uint32
	Type       string
	Species    string
	Position   sdlutils.Vector3
	Direction  common.KurinDirection
	Fatigue    int32
	ActiveHand gameplay.KurinHand
	Inventory  KurinInventoryData
	JobTracker KurinJobTrackerData
}

func EncodeKurinCharacter(character *gameplay.KurinCharacter) KurinCharacterData {
	data := KurinCharacterData{
		Id:         character.Id,
		Type:       character.Type,
		Species:    character.Species,
		Position:   character.Position,
		Direction:  character.Direction,
		Fatigue:    character.Fatigue,
		ActiveHand: character.ActiveHand,
		Inventory:  EncodeKurinInventory(&character.Inventory),
		JobTracker: EncodeKurinJobTracker(character.JobTracker),
	}

	return data
}

func DecodeKurinCharacter(data KurinCharacterData) *gameplay.KurinCharacter {
	character := gameplay.NewKurinCharacter()
	character.Id = data.Id
	character.Type = data.Type
	character.Species = data.Species
	character.Position = data.Position
	character.PositionRender = sdlutils.PointToFPoint(data.Position.Base)
	character.Direction = data.Direction
	character.Fatigue = data.Fatigue
	character.ActiveHand = data.ActiveHand
	character.Inventory = DecodeKurinInventory(data.Inventory, character)
	character.JobTracker = DecodeKurinJobTracker(character, data.JobTracker)

	return character
}

package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
)

type CharacterData struct {
	Id         uint32
	Species    string
	Gender     string
	Faction    gameplay.Faction
	Position   sdlutils.Vector3
	Direction  common.Direction
	Fatigue    int32
	ActiveHand gameplay.Hand
	Inventory  InventoryData
	JobTracker JobTrackerData
}

func EncodeCharacter(character *gameplay.Character) CharacterData {
	data := CharacterData{
		Id:         character.Id,
		Species:    character.Species,
		Gender:     character.Gender,
		Position:   character.Position,
		Direction:  character.Direction,
		Fatigue:    character.Fatigue,
		ActiveHand: character.ActiveHand,
		Inventory:  EncodeInventory(&character.Inventory),
		JobTracker: EncodeJobTracker(character.JobTracker),
	}

	return data
}

func DecodeCharacter(data CharacterData) *gameplay.Character {
	character := gameplay.NewCharacter(data.Faction)
	character.Id = data.Id
	character.Gender = data.Gender
	character.Species = data.Species
	character.Position = data.Position
	character.PositionRender = sdlutils.PointToFPoint(data.Position.Base)
	character.Direction = data.Direction
	character.Fatigue = data.Fatigue
	character.ActiveHand = data.ActiveHand
	character.Inventory = DecodeInventory(data.Inventory, character)
	character.JobTracker = DecodeJobTracker(character, data.JobTracker)

	return character
}

package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type CharacterData struct {
	Inventory  InventoryData
	JobTracker JobTrackerData
}

func EncodeCharacter(character *gameplay.Mob) CharacterData {
	characterData := character.Data.(*gameplay.MobCharacterData)
	data := CharacterData{
		Inventory: EncodeInventory(characterData.Inventory),
	}

	return data
}

func DecodeCharacter(data CharacterData) *gameplay.Mob {
	character := gameplay.NewMob("character", gameplay.FactionPlayer)
	character.Data.(*gameplay.MobCharacterData).Inventory = DecodeInventory(data.Inventory, character)

	return character
}

package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type CharacterData struct {
	Inventory InventoryData
}

func EncodeCharacterData(character *gameplay.Mob) CharacterData {
	characterData := character.Data.(*gameplay.MobCharacterData)
	data := CharacterData{
		Inventory: EncodeInventory(characterData.Inventory),
	}

	return data
}

func DecodeCharacterData(data CharacterData, character *gameplay.Mob) {
	character.Data.(*gameplay.MobCharacterData).Inventory = DecodeInventory(data.Inventory, character)
}

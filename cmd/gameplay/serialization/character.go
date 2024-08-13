package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinCharacterData struct {
	Id uint32
	Type string
	Species string
	Position       sdlutils.Vector3
	Direction gameplay.KurinDirection

	ActiveHand gameplay.KurinHand
	Fatigue int32

	Inventory           KurinInventoryData
}

func EncodeKurinCharacter(character *gameplay.KurinCharacter) KurinCharacterData {
	data := KurinCharacterData{
		Id: character.Id,
		Type: character.Type,
		Species: character.Species,
		Position: character.Position,
		Direction: character.Direction,
		ActiveHand: character.ActiveHand,
		Fatigue: character.Fatigue,
		Inventory: KurinInventoryData{
			Left: KurinItemData{},
			Right: KurinItemData{},
		},
	}
	left := character.Inventory.Hands[gameplay.KurinHandLeft]
	if left != nil {
		data.Inventory.Left = EncodeKurinItem(left)
	}
	right := character.Inventory.Hands[gameplay.KurinHandRight]
	if right != nil {
		data.Inventory.Right = EncodeKurinItem(right)
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
	character.ActiveHand = data.ActiveHand
	character.Fatigue = data.Fatigue
	character.Inventory = gameplay.NewKurinInventory()
	if data.Inventory.Left.Id != 0 {
		item := DecodeKurinItem(data.Inventory.Left)
		item.Character = character
		character.Inventory.Hands[gameplay.KurinHandLeft] = item
	}
	if data.Inventory.Right.Id != 0 {
		item := DecodeKurinItem(data.Inventory.Right)
		item.Character = character
		character.Inventory.Hands[gameplay.KurinHandRight] = item
	}

	return character
}

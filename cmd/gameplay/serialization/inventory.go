package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type InventoryData struct {
	Left  *ItemData
	Right *ItemData
}

func EncodeInventory(inventory *gameplay.Inventory) InventoryData {
	data := InventoryData{}
	left := inventory.Hands[gameplay.HandLeft]
	if left != nil {
		item := EncodeItem(left)
		data.Left = &item
	}
	right := inventory.Hands[gameplay.HandRight]
	if right != nil {
		item := EncodeItem(right)
		data.Right = &item
	}

	return data
}

func DecodeInventory(data InventoryData, character *gameplay.Character) gameplay.Inventory {
	inventory := gameplay.NewInventory()
	if data.Left != nil {
		item := DecodeItem(*data.Left)
		item.Character = character
		inventory.Hands[gameplay.HandLeft] = item
	}
	if data.Right != nil {
		item := DecodeItem(*data.Right)
		item.Character = character
		inventory.Hands[gameplay.HandRight] = item
	}

	return inventory
}

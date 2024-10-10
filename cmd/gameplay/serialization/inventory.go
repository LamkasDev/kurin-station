package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type InventoryData struct {
	ActiveHand gameplay.Hand
	Left       *ItemData
	Right      *ItemData
}

func EncodeInventory(inventory *gameplay.Inventory) InventoryData {
	data := InventoryData{
		ActiveHand: inventory.ActiveHand,
	}
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

func DecodeInventory(data InventoryData, character *gameplay.Mob) *gameplay.Inventory {
	inventory := gameplay.NewInventory()
	inventory.ActiveHand = data.ActiveHand
	if data.Left != nil {
		item := PredecodeItem(*data.Left)
		item.Mob = character
		inventory.Hands[gameplay.HandLeft] = item
	}
	if data.Right != nil {
		item := PredecodeItem(*data.Right)
		item.Mob = character
		inventory.Hands[gameplay.HandRight] = item
	}

	return inventory
}

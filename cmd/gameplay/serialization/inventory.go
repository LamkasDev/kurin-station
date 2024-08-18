package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type KurinInventoryData struct {
	Left  *KurinItemData
	Right *KurinItemData
}

func EncodeKurinInventory(inventory *gameplay.KurinInventory) KurinInventoryData {
	data := KurinInventoryData{}
	left := inventory.Hands[gameplay.KurinHandLeft]
	if left != nil {
		item := EncodeKurinItem(left)
		data.Left = &item
	}
	right := inventory.Hands[gameplay.KurinHandRight]
	if right != nil {
		item := EncodeKurinItem(right)
		data.Right = &item
	}

	return data
}

func DecodeKurinInventory(data KurinInventoryData, character *gameplay.KurinCharacter) gameplay.KurinInventory {
	inventory := gameplay.NewKurinInventory()
	if data.Left != nil {
		item := DecodeKurinItem(*data.Left)
		item.Character = character
		inventory.Hands[gameplay.KurinHandLeft] = item
	}
	if data.Right != nil {
		item := DecodeKurinItem(*data.Right)
		item.Character = character
		inventory.Hands[gameplay.KurinHandRight] = item
	}

	return inventory
}

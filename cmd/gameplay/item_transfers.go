package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"golang.org/x/exp/slices"
)

func RemoveItemFromMapRaw(kmap *Map, item *Item) bool {
	i := slices.Index(kmap.Items, item)
	if i == -1 {
		return false
	}

	kmap.Items = slices.Delete(kmap.Items, i, i+1)
	return true
}

func RemoveItemFromCharacterRaw(item *Item, character *Mob) bool {
	characterData := character.Data.(*MobCharacterData)
	for hand := range characterData.Inventory.Hands {
		if characterData.Inventory.Hands[hand] == item {
			item.Mob = nil
			characterData.Inventory.Hands[hand] = nil
			return true
		}
	}

	return false
}

func AddItemToMapRaw(item *Item, kmap *Map, transform *sdlutils.Transform) *Item {
	item.Transform = transform
	kmap.Items = append(kmap.Items, item)

	return item
}

func AddItemToCharacterRaw(item *Item, character *Mob) bool {
	handItem := GetHeldItem(character)
	if handItem == nil {
		GetInventory(character).Hands[GetActiveHand(character)] = item
		item.Mob = character
		return true
	} else if handItem.Type == item.Type {
		transfer := min(handItem.Template.MaxCount-handItem.Count, item.Count)
		handItem.Count += transfer
		item.Count -= transfer
		if item.Count == 0 {
			return true
		}
		return false
	}

	return false
}

func TransferItemToCharacterRaw(item *Item, kmap *Map, character *Mob) bool {
	if !AddItemToCharacterRaw(item, character) {
		return false
	}

	item.Transform = nil
	RemoveItemFromMapRaw(kmap, item)
	return true
}

func TransferItemFromCharacterRaw(item *Item, kmap *Map, character *Mob) bool {
	if !RemoveItemFromCharacterRaw(item, character) {
		return false
	}
	AddItemToMapRaw(item, kmap, &sdlutils.Transform{
		Position: sdlutils.Vector3ToFVector3Center(character.Position),
		Rotation: 0,
	})

	return true
}

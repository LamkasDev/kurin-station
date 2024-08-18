package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"golang.org/x/exp/slices"
)

func RemoveKurinItemFromMapRaw(kmap *KurinMap, item *KurinItem) bool {
	i := slices.Index(kmap.Items, item)
	if i == -1 {
		return false
	}

	kmap.Items = slices.Delete(kmap.Items, i, i+1)
	return true
}

func RemoveKurinItemFromCharacterRaw(item *KurinItem, character *KurinCharacter) bool {
	for hand := range character.Inventory.Hands {
		if character.Inventory.Hands[hand] == item {
			item.Character = nil
			character.Inventory.Hands[hand] = nil
			return true
		}
	}

	return false
}

func AddKurinItemToMapRaw(item *KurinItem, kmap *KurinMap, transform *sdlutils.Transform) {
	item.Transform = transform
	kmap.Items = append(kmap.Items, item)
}

func AddKurinItemToCharacterRaw(item *KurinItem, character *KurinCharacter) bool {
	handItem := character.Inventory.Hands[character.ActiveHand]
	if handItem == nil {
		character.Inventory.Hands[character.ActiveHand] = item
		item.Character = character
		return true
	} else if handItem.Type == item.Type {
		transfer := min(handItem.MaxCount-handItem.Count, item.Count)
		handItem.Count += transfer
		item.Count -= transfer
		if item.Count == 0 {
			return true
		}
		return false
	}

	return false
}

func TransferKurinItemToCharacterRaw(item *KurinItem, kmap *KurinMap, character *KurinCharacter) bool {
	if !AddKurinItemToCharacterRaw(item, character) {
		return false
	}

	item.Transform = nil
	RemoveKurinItemFromMapRaw(kmap, item)
	return true
}

func TransferKurinItemFromCharacterRaw(item *KurinItem, kmap *KurinMap, character *KurinCharacter) bool {
	if !RemoveKurinItemFromCharacterRaw(item, character) {
		return false
	}
	AddKurinItemToMapRaw(item, kmap, &sdlutils.Transform{
		Position: sdlutils.Vector3ToFVector3Center(character.Position),
		Rotation: 0,
	})

	return true
}

package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type KurinItem struct {
	Type     string
	Transform *sdlutils.Transform
}

func NewKurinItem(itype string, transform *sdlutils.Transform) *KurinItem {
	return &KurinItem{
		Type:     itype,
		Transform: transform,
	}
}

func NewKurinItemRandom(itype string, kmap *KurinMap) *KurinItem {
	item := &KurinItem{
		Type: itype,
	}
	for {
		position := sdlutils.Vector3{Base: sdl.Point{X: int32(rand.Float32() * float32(kmap.Size.Base.X)), Y: int32(rand.Float32() * float32(kmap.Size.Base.Y))}, Z: 0}
		if CanEnterPosition(kmap, position) {
			item.Transform = &sdlutils.Transform{
				Position: sdlutils.FVector3{
					Base: sdl.FPoint{
						X: float32(position.Base.X) + 0.5,
						Y: float32(position.Base.Y) + 0.5,
					},
					Z: position.Z,
				},
				Rotation: 0,
			}
			break
		}
	}

	return item
}

func RawRemoveKurinItemFromMap(item *KurinItem, kmap *KurinMap) bool {
	i := slices.Index(kmap.Items, item)
	if i == -1 {
		return false
	}

    kmap.Items[i] = kmap.Items[len(kmap.Items)-1]
    kmap.Items = kmap.Items[:len(kmap.Items)-1]
	return true
}

func RawRemoveKurinItemFromCharacter(item *KurinItem, character *KurinCharacter) bool {
	for hand := range character.Inventory.Hands {
		if character.Inventory.Hands[hand] == item {
			character.Inventory.Hands[hand] = nil
			return true
		}
	}

	return false
}

func RawAddKurinItemToMap(item *KurinItem, kmap *KurinMap) {
    kmap.Items = append(kmap.Items, item)
}

func RawAddKurinItemToCharacter(item *KurinItem, character *KurinCharacter) bool {
	if !RawIsKurinCharacterHandEmpty(character) {
		return false
	}

	character.Inventory.Hands[character.ActiveHand] = item
	return true
}

func RawIsKurinCharacterHandEmpty(character *KurinCharacter) bool {
	return character.Inventory.Hands[character.ActiveHand] == nil
}

func RawTransferKurinItemToCharacter(item *KurinItem, kmap *KurinMap, character *KurinCharacter) bool {
	if !RawAddKurinItemToCharacter(item, character) {
		return false
	}

	item.Transform = nil
	RawRemoveKurinItemFromMap(item, kmap)
	return true
}

func RawTransferKurinItemFromCharacter(item *KurinItem, kmap *KurinMap, character *KurinCharacter) bool {
	if !RawRemoveKurinItemFromCharacter(item, character) {
		return false
	}

	item.Transform = &sdlutils.Transform{
		Position: sdlutils.FVector3{
			Base: sdl.FPoint{
				X: float32(character.Position.Base.X) + 0.5,
				Y: float32(character.Position.Base.Y) + 0.5,
			},
			Z: character.Position.Z,
		},
		Rotation: 0,
	}
	RawAddKurinItemToMap(item, kmap)
	return true
}

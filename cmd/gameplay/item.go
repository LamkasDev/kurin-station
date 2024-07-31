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

	GetTextures KurinItemGetTextures
	GetTextureHand KurinItemGetTextureHand
	Interact KurinItemInteract
	Process  KurinItemProcess
	Data interface{}
}

type KurinItemGetTextures func(item *KurinItem, game *KurinGame) []int
type KurinItemGetTextureHand func(item *KurinItem, game *KurinGame) int
type KurinItemInteract func(item *KurinItem, game *KurinGame)
type KurinItemProcess func(item *KurinItem, game *KurinGame)

func NewKurinItemRandom(itemType string, kmap *KurinMap) *KurinItem {
	item := NewKurinItem(itemType, nil)
	for {
		position := sdlutils.Vector3{Base: sdl.Point{X: int32(rand.Float32() * float32(kmap.Size.Base.X)), Y: int32(rand.Float32() * float32(kmap.Size.Base.Y))}, Z: 0}
		if CanEnterPosition(kmap, position) {
			item.Transform = &sdlutils.Transform{
				Position:  sdlutils.Vector3ToFVector3Center(position),
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
		Position:  sdlutils.Vector3ToFVector3Center(character.Position),
		Rotation: 0,
	}
	RawAddKurinItemToMap(item, kmap)
	return true
}

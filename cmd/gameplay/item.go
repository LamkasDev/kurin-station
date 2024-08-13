package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type KurinItem struct {
	Id uint32
	Type     string
	Transform *sdlutils.Transform
	Character *KurinCharacter

	GetTextures KurinItemGetTextures
	GetTextureHand KurinItemGetTextureHand
	OnHandInteraction KurinItemOnHandInteraction
	OnTileInteraction KurinItemOnTileInteraction
	EncodeData KurinItemEncodeData
	DecodeData KurinItemDecodeData
	CanHit bool
	Process  KurinItemProcess
	Data interface{}
}

type KurinItemGetTextures func(item *KurinItem) []int
type KurinItemGetTextureHand func(item *KurinItem) int
type KurinItemOnHandInteraction func(item *KurinItem)
type KurinItemOnTileInteraction func(item *KurinItem, tile *KurinTile) bool
type KurinItemEncodeData func(item *KurinItem) []byte
type KurinItemDecodeData func(item *KurinItem, data []byte)
type KurinItemProcess func(item *KurinItem)

func NewKurinItemRandom(itemType string, kmap *KurinMap) *KurinItem {
	item := NewKurinItem(itemType)
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

func RemoveKurinItemFromMapRaw(item *KurinItem, kmap *KurinMap) bool {
	i := slices.Index(kmap.Items, item)
	if i == -1 {
		return false
	}

    kmap.Items[i] = kmap.Items[len(kmap.Items)-1]
    kmap.Items = kmap.Items[:len(kmap.Items)-1]
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
	if !IsKurinCharacterHandEmptyRaw(character) {
		return false
	}

	character.Inventory.Hands[character.ActiveHand] = item
	item.Character = character
	return true
}

func IsKurinCharacterHandEmptyRaw(character *KurinCharacter) bool {
	return character.Inventory.Hands[character.ActiveHand] == nil
}

func TransferKurinItemToCharacterRaw(item *KurinItem, kmap *KurinMap, character *KurinCharacter) bool {
	if !AddKurinItemToCharacterRaw(item, character) {
		return false
	}

	item.Transform = nil
	RemoveKurinItemFromMapRaw(item, kmap)
	return true
}

func TransferKurinItemFromCharacterRaw(item *KurinItem, kmap *KurinMap, character *KurinCharacter) bool {
	if !RemoveKurinItemFromCharacterRaw(item, character) {
		return false
	}

	AddKurinItemToMapRaw(item, kmap, &sdlutils.Transform{
		Position:  sdlutils.Vector3ToFVector3Center(character.Position),
		Rotation: 0,
	})
	return true
}

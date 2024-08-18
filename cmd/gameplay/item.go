package gameplay

import (
	"sort"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"robpike.io/filter"
)

type KurinItem struct {
	Id        uint32
	Type      string
	Count     uint16
	Reserved  bool
	Transform *sdlutils.Transform
	Character *KurinCharacter

	Process           KurinItemProcess
	GetTextures       KurinItemGetTextures
	GetTextureHand    KurinItemGetTextureHand
	OnHandInteraction KurinItemOnHandInteraction
	OnTileInteraction KurinItemOnTileInteraction
	EncodeData        KurinItemEncodeData
	DecodeData        KurinItemDecodeData
	CanHit            bool
	MaxCount          uint16
	Data              interface{}
}

type (
	KurinItemProcess           func(item *KurinItem)
	KurinItemGetTextures       func(item *KurinItem) []int
	KurinItemGetTextureHand    func(item *KurinItem) int
	KurinItemOnHandInteraction func(item *KurinItem)
	KurinItemOnTileInteraction func(item *KurinItem, tile *KurinTile) bool
	KurinItemEncodeData        func(item *KurinItem) []byte
	KurinItemDecodeData        func(item *KurinItem, data []byte)
)

func NewKurinItemRandom(kmap *KurinMap, itemType string, count uint16) *KurinItem {
	item := NewKurinItem(itemType, count)
	for {
		position := GetRandomMapPosition(kmap)
		if CanEnterMapPosition(kmap, position) {
			item.Transform = &sdlutils.Transform{
				Position: sdlutils.Vector3ToFVector3Center(position),
				Rotation: 0,
			}
			break
		}
	}

	return item
}

func FindItemsOfType(kmap *KurinMap, itemType string, reservation bool) []*KurinItem {
	return filter.Choose(kmap.Items, func(item *KurinItem) bool {
		return item.Type == itemType && (!reservation || !item.Reserved)
	}).([]*KurinItem)
}

func FindClosestItemOfType(kmap *KurinMap, position sdlutils.Vector3, itemType string, reservation bool) *KurinItem {
	items := FindItemsOfType(kmap, itemType, reservation)
	start := sdlutils.PointToFPoint(position.Base)
	sort.Slice(items, func(i, j int) bool {
		return sdlutils.GetDistanceSimpleF(start, items[i].Transform.Position.Base) < sdlutils.GetDistanceSimpleF(start, items[j].Transform.Position.Base)
	})
	if len(items) == 0 {
		return nil
	}

	return items[0]
}

func ReserveKurinItem(item *KurinItem) {
	item.Reserved = true
}

func UnreserveKurinItem(item *KurinItem) {
	item.Reserved = false
}

package gameplay

import (
	"slices"
	"sort"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type Item struct {
	Id        uint32
	Type      string
	Count     uint16
	Reserved  bool
	Transform *sdlutils.Transform
	Mob       *Mob

	Template *ItemTemplate
	Data     interface{}
}

type ItemRequirement struct {
	Type  string
	Count uint16
}

func NewItemRandom(kmap *Map, itemType string, count uint16) *Item {
	item := NewItem(itemType, count)
	for {
		position := GetRandomMapPosition(kmap, kmap.BaseZ)
		if CanEnterMapPosition(kmap, position) == EnteranceStatusYes {
			item.Transform = &sdlutils.Transform{
				Position: sdlutils.Vector3ToFVector3Center(position),
				Rotation: 0,
			}
			break
		}
	}

	return item
}

func FindItems(kmap *Map, predicate func(item *Item) bool) []*Item {
	return slices.Collect(func(yield func(*Item) bool) {
		for _, item := range kmap.Items {
			if predicate(item) {
				if !yield(item) {
					return // triggered in "break"
				}
			}
		}
	})
}

func GetItemsOnTile(kmap *Map, tile *Tile) []*Item {
	return FindItems(kmap, func(item *Item) bool {
		return item.Transform.Position.Z == tile.Position.Z &&
			sdlutils.ComparePoints(sdlutils.FPointToPointFloored(item.Transform.Position.Base), tile.Position.Base)
	})
}

func FindItemsOfType(kmap *Map, itemType string, z uint8, reservation bool) []*Item {
	return FindItems(kmap, func(item *Item) bool {
		return item.Type == itemType && item.Transform.Position.Z == z && (!reservation || !item.Reserved)
	})
}

// TODO: z-levels
func FindClosestItemOfType(kmap *Map, position sdlutils.Vector3, itemType string, reservation bool) *Item {
	items := FindItemsOfType(kmap, itemType, position.Z, reservation)
	start := sdlutils.PointToFPoint(position.Base)
	sort.Slice(items, func(i, j int) bool {
		id := sdlutils.GetDistanceSimpleF(start, items[i].Transform.Position.Base)
		jd := sdlutils.GetDistanceSimpleF(start, items[j].Transform.Position.Base)
		return id < jd
	})
	if len(items) == 0 {
		return nil
	}

	return items[0]
}

func GetItemDescription(item *Item) string {
	return item.Type
}

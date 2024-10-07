package gameplay

import (
	"sort"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"robpike.io/filter"
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
		position := GetRandomMapPosition(kmap)
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

func GetItemsOnTile(kmap *Map, tile *Tile) []*Item {
	return filter.Choose(kmap.Items, func(item *Item) bool {
		return item.Transform.Position.Z == tile.Position.Z &&
			sdlutils.ComparePoints(sdlutils.FPointToPointFloored(item.Transform.Position.Base), tile.Position.Base)
	}).([]*Item)
}

func FindItemsOfType(kmap *Map, itemType string, reservation bool) []*Item {
	return filter.Choose(kmap.Items, func(item *Item) bool {
		return item.Type == itemType && (!reservation || !item.Reserved)
	}).([]*Item)
}

func FindClosestItemOfType(kmap *Map, position sdlutils.Vector3, itemType string, reservation bool) *Item {
	items := FindItemsOfType(kmap, itemType, reservation)
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

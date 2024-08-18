package gameplay

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

var KurinTileSize = sdl.Point{
	X: 32,
	Y: 32,
}

var KurinTileSizeF = sdl.FPoint{
	X: 32,
	Y: 32,
}

type KurinTile struct {
	Type     string
	Position sdlutils.Vector3

	Job     *KurinJobDriver
	Objects []*KurinObject
}

func NewKurinTile(tileType string, position sdlutils.Vector3) *KurinTile {
	return &KurinTile{
		Type:     tileType,
		Position: position,
		Job:      nil,
		Objects:  []*KurinObject{},
	}
}

func GetKurinTileAt(kmap *KurinMap, position sdlutils.Vector3) *KurinTile {
	if IsMapPositionOutOfBounds(kmap, position) {
		return nil
	}

	return kmap.Tiles[position.Base.X][position.Base.Y][position.Z]
}

func GetKurinTileInDirection(kmap *KurinMap, tile *KurinTile, direction common.KurinDirection) *KurinTile {
	return GetKurinTileAt(kmap, common.GetPositionInDirectionV(tile.Position, direction))
}

func DoesMapPositionHaveKurinTileNeighbour(kmap *KurinMap, position sdlutils.Vector3) bool {
	if GetKurinTileAt(kmap, common.GetPositionInDirectionV(position, common.KurinDirectionNorth)) != nil {
		return true
	}
	if GetKurinTileAt(kmap, common.GetPositionInDirectionV(position, common.KurinDirectionEast)) != nil {
		return true
	}
	if GetKurinTileAt(kmap, common.GetPositionInDirectionV(position, common.KurinDirectionSouth)) != nil {
		return true
	}
	if GetKurinTileAt(kmap, common.GetPositionInDirectionV(position, common.KurinDirectionWest)) != nil {
		return true
	}

	return false
}

func CanEnterKurinTile(tile *KurinTile) bool {
	return len(tile.Objects) == 0
}

func CanBuildKurinTileAtMapPosition(kmap *KurinMap, position sdlutils.Vector3) bool {
	return !IsMapPositionOutOfBounds(kmap, position) &&
		GetKurinTileAt(kmap, position) == nil &&
		DoesMapPositionHaveKurinTileNeighbour(kmap, position)
}

func GetKurinTileDescription(tile *KurinTile) string {
	text := fmt.Sprintf("[%d_%d] %s", tile.Position.Base.X, tile.Position.Base.Y, tile.Type)
	object := GetKurinObjectAtTile(tile)
	if object != nil {
		text = fmt.Sprintf("%s %s", text, object.Type)
	}

	return text
}

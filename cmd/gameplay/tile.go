package gameplay

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

var TileSize = sdl.Point{
	X: 32,
	Y: 32,
}

var TileSizeF = sdl.FPoint{
	X: 32,
	Y: 32,
}

type Tile struct {
	Type     string
	Position sdlutils.Vector3

	Job     *JobDriver
	Objects []*Object
}

func NewTile(tileType string, position sdlutils.Vector3) *Tile {
	return &Tile{
		Type:     tileType,
		Position: position,
		Job:      nil,
		Objects:  []*Object{},
	}
}

func GetTileAt(kmap *Map, position sdlutils.Vector3) *Tile {
	if IsMapPositionOutOfBounds(kmap, position) {
		return nil
	}

	return kmap.Tiles[position.Base.X][position.Base.Y][position.Z]
}

func GetTileInDirection(kmap *Map, tile *Tile, direction common.Direction) *Tile {
	return GetTileAt(kmap, common.GetPositionInDirectionV(tile.Position, direction))
}

func DoesMapPositionHaveTileNeighbour(kmap *Map, position sdlutils.Vector3) bool {
	if GetTileAt(kmap, common.GetPositionInDirectionV(position, common.DirectionNorth)) != nil {
		return true
	}
	if GetTileAt(kmap, common.GetPositionInDirectionV(position, common.DirectionEast)) != nil {
		return true
	}
	if GetTileAt(kmap, common.GetPositionInDirectionV(position, common.DirectionSouth)) != nil {
		return true
	}
	if GetTileAt(kmap, common.GetPositionInDirectionV(position, common.DirectionWest)) != nil {
		return true
	}

	return false
}

func CanBuildTileAtMapPosition(kmap *Map, position sdlutils.Vector3) bool {
	return !IsMapPositionOutOfBounds(kmap, position) &&
		GetTileAt(kmap, position) == nil &&
		DoesMapPositionHaveTileNeighbour(kmap, position)
}

func GetTileDescription(tile *Tile) string {
	text := fmt.Sprintf("[%d_%d] %s", tile.Position.Base.X, tile.Position.Base.Y, tile.Type)
	object := GetObjectAtTile(tile)
	if object != nil {
		text = fmt.Sprintf("%s %s", text, object.Type)
	}

	return text
}

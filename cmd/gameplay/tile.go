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

const (
	TileIdFloor    = uint8(0)
	TileIdBlank    = uint8(1)
	TileIdCatwalk  = uint8(2)
	TileIdAsteroid = uint8(3)
)

var TileIdMap = map[uint8]string{
	TileIdFloor:    "floor",
	TileIdBlank:    "blank",
	TileIdCatwalk:  "catwalk",
	TileIdAsteroid: "asteroid",
}

type Tile struct {
	Type     uint8
	Position sdlutils.Vector3
	Seed     uint8
	Job      *JobDriver
	Objects  []*Object

	Template *TileTemplate
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

func CanDestroyTileAtMapPosition(kmap *Map, position sdlutils.Vector3) bool {
	if !IsMapPositionOutOfBounds(kmap, position) {
		return false
	}
	tile := GetTileAt(kmap, position)

	return tile.Type != TileIdAsteroid
}

func GetTileDescription(tile *Tile) string {
	text := fmt.Sprintf("[%d_%d] %s", tile.Position.Base.X, tile.Position.Base.Y, TileIdMap[tile.Type])
	object := GetObjectAtTile(tile)
	if object != nil {
		text = fmt.Sprintf("%s %s", text, object.Type)
	}

	return text
}

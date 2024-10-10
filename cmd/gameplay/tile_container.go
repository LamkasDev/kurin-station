package gameplay

import (
	"hash/fnv"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

var (
	TileHasher    = fnv.New32a()
	TileContainer = map[uint8]*TileTemplate{}
)

func RegisterTiles() {
	TileContainer[TileIdAsteroid] = NewTileTemplateAsteroid()
	TileContainer[TileIdBlank] = NewTileTemplate("blank")
	TileContainer[TileIdCatwalk] = NewTileTemplate("catwalk")
	TileContainer[TileIdFloor] = NewTileTemplate("floor")
}

func NewTile(tileType uint8, position sdlutils.Vector3) *Tile {
	tile := &Tile{
		Type:     tileType,
		Position: position,
		Job:      nil,
		Objects:  []*Object{},
		Template: TileContainer[tileType],
	}
	TileHasher.Write([]byte{byte(tile.Position.Base.X >> 8), byte(tile.Position.Base.X), byte(tile.Position.Base.Y >> 8), byte(tile.Position.Base.Y)})
	tile.Seed = uint8(TileHasher.Sum32() % 100)
	TileHasher.Reset()
	tile.Template.OnCreate(tile)

	return tile
}

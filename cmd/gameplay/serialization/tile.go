package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type TileData struct {
	Type     string
	Position sdlutils.Vector3
}

func EncodeTile(tile *gameplay.Tile) TileData {
	data := TileData{
		Type:     tile.Type,
		Position: tile.Position,
	}

	return data
}

func DecodeTile(kmap *gameplay.Map, data TileData) *gameplay.Tile {
	tile := gameplay.CreateTileRaw(kmap, data.Position, data.Type)

	return tile
}

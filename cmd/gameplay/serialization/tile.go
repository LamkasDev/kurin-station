package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinTileData struct {
	Type     string
	Position sdlutils.Vector3
}

func EncodeKurinTile(tile *gameplay.KurinTile) KurinTileData {
	data := KurinTileData{
		Type:     tile.Type,
		Position: tile.Position,
	}

	return data
}

func DecodeKurinTile(kmap *gameplay.KurinMap, data KurinTileData) *gameplay.KurinTile {
	tile := gameplay.CreateKurinTileRaw(kmap, data.Position, data.Type)

	return tile
}

package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

type EnteranceStatus uint8

const (
	EnteranceStatusYes      = EnteranceStatus(0)
	EnteranceStatusNo       = EnteranceStatus(1)
	EnteranceStatusPossible = EnteranceStatus(2)
)

func CanEnterMapPosition(kmap *Map, position sdlutils.Vector3) EnteranceStatus {
	tile := GetTileAt(kmap, position)
	if tile == nil {
		return EnteranceStatusNo
	}

	return CanEnterTile(tile)
}

func CanEnterTile(tile *Tile) EnteranceStatus {
	for _, object := range tile.Objects {
		if !object.Template.IsPassable(object) {
			if object.Type == "airlock" {
				return EnteranceStatusPossible
			}

			return EnteranceStatusNo
		}
	}

	return EnteranceStatusYes
}

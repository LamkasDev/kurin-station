package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type MapData struct {
	Size    sdlutils.Vector3
	BaseZ   uint8
	Tiles   []TileData
	Objects []ObjectData
	Items   []ItemData
}

func EncodeMap(kmap *gameplay.Map) MapData {
	data := MapData{
		Size:  kmap.Size,
		BaseZ: kmap.BaseZ,
		Tiles: []TileData{},
		Items: []ItemData{},
	}
	for x := range kmap.Size.Base.X {
		for y := range kmap.Size.Base.Y {
			for z := range kmap.Size.Z {
				tile := kmap.Tiles[x][y][z]
				if tile == nil {
					continue
				}
				data.Tiles = append(data.Tiles, EncodeTile(tile))
			}
		}
	}
	for _, obj := range kmap.Objects {
		data.Objects = append(data.Objects, EncodeObject(obj))
	}
	for _, item := range kmap.Items {
		data.Items = append(data.Items, EncodeItem(item))
	}

	return data
}

func DecodeMap(data MapData) gameplay.Map {
	kmap := gameplay.NewMap(data.Size, data.BaseZ)
	for _, tileData := range data.Tiles {
		DecodeTile(&kmap, tileData)
	}
	for _, objData := range data.Objects {
		DecodeObject(&kmap, objData)
	}
	for _, itemData := range data.Items {
		kmap.Items = append(kmap.Items, DecodeItem(itemData))
	}

	return kmap
}

package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinMapData struct {
	Size    sdlutils.Vector3
	Tiles   []KurinTileData
	Objects []KurinObjectData
	Items   []KurinItemData
}

func EncodeKurinMap(kmap *gameplay.KurinMap) KurinMapData {
	data := KurinMapData{
		Size:  kmap.Size,
		Tiles: []KurinTileData{},
		Items: []KurinItemData{},
	}
	for x := range kmap.Size.Base.X {
		for y := range kmap.Size.Base.Y {
			for z := range kmap.Size.Z {
				tile := kmap.Tiles[x][y][z]
				if tile == nil {
					continue
				}
				data.Tiles = append(data.Tiles, EncodeKurinTile(tile))
			}
		}
	}
	for _, obj := range kmap.Objects {
		data.Objects = append(data.Objects, EncodeKurinObject(obj))
	}
	for _, item := range kmap.Items {
		data.Items = append(data.Items, EncodeKurinItem(item))
	}

	return data
}

func DecodeKurinMap(data KurinMapData) gameplay.KurinMap {
	kmap := gameplay.NewKurinMap(data.Size)
	for _, tileData := range data.Tiles {
		DecodeKurinTile(&kmap, tileData)
	}
	for _, objData := range data.Objects {
		DecodeKurinObject(&kmap, objData)
	}
	for _, itemData := range data.Items {
		kmap.Items = append(kmap.Items, DecodeKurinItem(itemData))
	}

	return kmap
}

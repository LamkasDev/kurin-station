package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinMapData struct {
	Size sdlutils.Vector3
	Objects []KurinObjectData
	Items []KurinItemData
}

func EncodeKurinMap(kmap *gameplay.KurinMap) KurinMapData {
	data := KurinMapData{
		Size: kmap.Size,
		Items: []KurinItemData{},
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
	for _, objData := range data.Objects {
		DecodeKurinObject(&kmap, objData)
	}
	for _, itemData := range data.Items {
		kmap.Items = append(kmap.Items, DecodeKurinItem(itemData))
	}

	return kmap
}

package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinItemData struct {
	Id        uint32
	Type      string
	Count     uint16
	Transform *sdlutils.Transform
	Data      []byte
}

func EncodeKurinItem(item *gameplay.KurinItem) KurinItemData {
	return KurinItemData{
		Id:        item.Id,
		Type:      item.Type,
		Count:     item.Count,
		Transform: item.Transform,
		Data:      item.EncodeData(item),
	}
}

func DecodeKurinItem(data KurinItemData) *gameplay.KurinItem {
	item := gameplay.NewKurinItem(data.Type, data.Count)
	item.Id = data.Id
	item.Transform = data.Transform
	item.DecodeData(item, data.Data)

	return item
}

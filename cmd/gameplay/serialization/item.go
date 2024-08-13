package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinItemData struct {
	Id   uint32
	Type string
	Transform *sdlutils.Transform
	Data []byte
}

func EncodeKurinItem(item *gameplay.KurinItem) KurinItemData {
	return KurinItemData{
		Id: item.Id,
		Type: item.Type,
		Transform: item.Transform,
		Data: item.EncodeData(item),
	}
}

func DecodeKurinItem(data KurinItemData) *gameplay.KurinItem {
	item := gameplay.NewKurinItem(data.Type)
	item.Id = data.Id
	item.Transform = data.Transform
	item.DecodeData(item, data.Data)

	return item
}

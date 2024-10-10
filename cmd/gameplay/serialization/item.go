package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type ItemData struct {
	Id        uint32
	Type      string
	Count     uint16
	Transform *sdlutils.Transform
	Data      []byte
}

func EncodeItem(item *gameplay.Item) ItemData {
	return ItemData{
		Id:        item.Id,
		Type:      item.Type,
		Count:     item.Count,
		Transform: item.Transform,
		Data:      item.Template.EncodeData(item),
	}
}

func PredecodeItem(data ItemData) *gameplay.Item {
	item := gameplay.NewItem(data.Type, data.Count)
	item.Id = data.Id
	item.Transform = data.Transform
	item.Template.DecodeData(item, data.Data)

	return item
}

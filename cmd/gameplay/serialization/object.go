package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
)

type ObjectData struct {
	Id        uint32
	Type      string
	Position  sdlutils.Vector3
	Direction common.Direction
	Health    uint16
	Data      []byte
}

func EncodeObject(obj *gameplay.Object) ObjectData {
	data := ObjectData{
		Id:        obj.Id,
		Type:      obj.Type,
		Position:  obj.Tile.Position,
		Direction: obj.Direction,
		Health:    obj.Health,
	}
	if obj.Data != nil {
		data.Data = obj.Template.EncodeData(obj)
	}

	return data
}

func PredecodeObject(kmap *gameplay.Map, data ObjectData) *gameplay.Object {
	obj := gameplay.CreateObjectRaw(kmap, kmap.Tiles[data.Position.Base.X][data.Position.Base.Y][data.Position.Z], data.Type)
	obj.Id = data.Id
	obj.Direction = data.Direction
	obj.Health = data.Health

	return obj
}

func DecodeObject(kmap *gameplay.Map, obj *gameplay.Object, data ObjectData) *gameplay.Object {
	if data.Data != nil {
		obj.Template.DecodeData(obj, data.Data)
	}

	return obj
}

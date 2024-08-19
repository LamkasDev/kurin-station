package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

type ObjectData struct {
	Id        uint32
	Type      string
	Position  sdl.Point
	Direction common.Direction
	Health    uint16
	Data      []byte
}

func EncodeObject(obj *gameplay.Object) ObjectData {
	data := ObjectData{
		Id:        obj.Id,
		Type:      obj.Type,
		Position:  obj.Tile.Position.Base,
		Direction: obj.Direction,
		Health:    obj.Health,
		Data:      obj.Template.EncodeData(obj),
	}

	return data
}

func DecodeObject(kmap *gameplay.Map, data ObjectData) *gameplay.Object {
	obj := gameplay.CreateObjectRaw(kmap, kmap.Tiles[data.Position.X][data.Position.Y][0], data.Type)
	obj.Id = data.Id
	obj.Direction = data.Direction
	obj.Health = data.Health
	obj.Template.DecodeData(obj, data.Data)

	return obj
}

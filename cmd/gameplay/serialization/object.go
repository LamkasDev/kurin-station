package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinObjectData struct {
	Id   uint32
	Type string
	Position sdl.Point
	Direction gameplay.KurinDirection
	Health uint8
	Data []byte
}

func EncodeKurinObject(obj *gameplay.KurinObject) KurinObjectData {
	data := KurinObjectData{
		Id: obj.Id,
		Type: obj.Type,
		Position: obj.Tile.Position.Base,
		Direction: obj.Direction,
		Health: obj.Health,
		Data: obj.EncodeData(obj),
	}

	return data
}

func DecodeKurinObject(kmap *gameplay.KurinMap, data KurinObjectData) *gameplay.KurinObject {
	obj := gameplay.CreateKurinObjectRaw(kmap, kmap.Tiles[data.Position.X][data.Position.Y][0], data.Type)
	obj.Id = data.Id
	obj.Direction = data.Direction
	obj.Health = data.Health
	obj.DecodeData(obj, data.Data)

	return obj
}

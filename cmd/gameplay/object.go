package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinObject struct {
	Id        uint32
	Type      string
	Tile      *KurinTile
	Direction common.KurinDirection
	Health    uint16

	Process       KurinObjectProcess
	GetTexture    KurinObjectGetTexture
	OnInteraction KurinObjectOnInteraction
	OnCreate      KurinObjectOnDestroy
	OnDestroy     KurinObjectOnDestroy
	EncodeData    KurinObjectEncodeData
	DecodeData    KurinObjectDecodeData
	Data          interface{}
}

type (
	KurinObjectProcess       func(object *KurinObject)
	KurinObjectGetTexture    func(object *KurinObject) int
	KurinObjectOnInteraction func(object *KurinObject, item *KurinItem) bool
	KurinObjectOnCreate      func(object *KurinObject)
	KurinObjectOnDestroy     func(object *KurinObject)
	KurinObjectEncodeData    func(object *KurinObject) []byte
	KurinObjectDecodeData    func(object *KurinObject, data []byte)
)

func GetKurinObjectAtTile(tile *KurinTile) *KurinObject {
	if len(tile.Objects) == 0 {
		return nil
	}

	return tile.Objects[len(tile.Objects)-1]
}

func GetKurinObjectAtMapPosition(kmap *KurinMap, position sdlutils.Vector3) *KurinObject {
	tile := GetKurinTileAt(kmap, position)
	if tile == nil {
		return nil
	}

	return GetKurinObjectAtTile(tile)
}

func GetKurinObjectInDirection(kmap *KurinMap, object *KurinObject, direction common.KurinDirection) *KurinObject {
	return GetKurinObjectAtMapPosition(kmap, common.GetPositionInDirectionV(object.Tile.Position, direction))
}

func GetKurinObjectDirectionHint(kmap *KurinMap, obj *KurinObject) string {
	direction := ""
	if neighbour := GetKurinObjectInDirection(kmap, obj, common.KurinDirectionNorth); neighbour != nil && neighbour.Type == obj.Type {
		direction += "n"
	}
	if neighbour := GetKurinObjectInDirection(kmap, obj, common.KurinDirectionEast); neighbour != nil && neighbour.Type == obj.Type {
		direction += "e"
	}
	if neighbour := GetKurinObjectInDirection(kmap, obj, common.KurinDirectionSouth); neighbour != nil && neighbour.Type == obj.Type {
		direction += "s"
	}
	if neighbour := GetKurinObjectInDirection(kmap, obj, common.KurinDirectionWest); neighbour != nil && neighbour.Type == obj.Type {
		direction += "w"
	}

	return direction
}

func CanBuildKurinObjectAtMapPosition(kmap *KurinMap, position sdlutils.Vector3) bool {
	tile := GetKurinTileAt(kmap, position)
	if tile == nil {
		return false
	}

	return GetKurinObjectAtTile(tile) == nil
}

func GetKurinObjectSize(object *KurinObject) sdl.Point {
	switch object.Type {
	case "pod":
		return sdl.Point{X: 2, Y: 2}
	case "big_thruster":
		return sdl.Point{X: 3, Y: 1}
	}

	return sdl.Point{X: 1, Y: 1}
}

func GetKurinObjectCenter(object *KurinObject) sdl.FPoint {
	size := GetKurinObjectSize(object)
	return sdl.FPoint{
		X: float32(object.Tile.Position.Base.X) + float32(size.X)/2,
		Y: float32(object.Tile.Position.Base.Y) + float32(size.Y)/2,
	}
}

func HitKurinObject(object *KurinObject) {
	PlaySound(&GameInstance.SoundController, "grillehit")
	CreateKurinParticle(&GameInstance.ParticleController, NewKurinParticleCross(sdlutils.Vector3ToFVector3Center(object.Tile.Position), 0.75, sdl.Color{R: 210, G: 210, B: 210}))
	object.Health--
	if object.Health <= 0 {
		DestroyKurinObject(object)
	}
}

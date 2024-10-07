package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
	"robpike.io/filter"
)

type Object struct {
	Id        uint32
	Type      string
	Tile      *Tile
	Direction common.Direction
	Health    uint16

	Template *ObjectTemplate
	Data     interface{}
}

func GetObjectAtTile(tile *Tile) *Object {
	if len(tile.Objects) == 0 {
		return nil
	}

	return tile.Objects[len(tile.Objects)-1]
}

func GetObjectAtMapPosition(kmap *Map, position sdlutils.Vector3) *Object {
	tile := GetTileAt(kmap, position)
	if tile == nil {
		return nil
	}

	return GetObjectAtTile(tile)
}

func GetObjectInDirection(kmap *Map, object *Object, direction common.Direction) *Object {
	return GetObjectAtMapPosition(kmap, common.GetPositionInDirectionV(object.Tile.Position, direction))
}

func CanObjectsJoinHint(a *Object, b *Object) bool {
	if a.Template.Smooth && b.Template.Smooth {
		return true
	}

	return a.Type == b.Type
}

func GetObjectDirectionHint(kmap *Map, obj *Object) string {
	direction := ""
	if neighbour := GetObjectInDirection(kmap, obj, common.DirectionNorth); neighbour != nil && CanObjectsJoinHint(neighbour, obj) {
		direction += "n"
	}
	if neighbour := GetObjectInDirection(kmap, obj, common.DirectionEast); neighbour != nil && CanObjectsJoinHint(neighbour, obj) {
		direction += "e"
	}
	if neighbour := GetObjectInDirection(kmap, obj, common.DirectionSouth); neighbour != nil && CanObjectsJoinHint(neighbour, obj) {
		direction += "s"
	}
	if neighbour := GetObjectInDirection(kmap, obj, common.DirectionWest); neighbour != nil && CanObjectsJoinHint(neighbour, obj) {
		direction += "w"
	}

	return direction
}

func CanBuildObjectAtMapPosition(kmap *Map, position sdlutils.Vector3) bool {
	tile := GetTileAt(kmap, position)
	if tile == nil {
		return false
	}

	return GetObjectAtTile(tile) == nil
}

func GetObjectSize(object *Object) sdl.Point {
	switch object.Type {
	case "pod":
		return sdl.Point{X: 2, Y: 2}
	case "big_thruster":
		return sdl.Point{X: 3, Y: 1}
	}

	return sdl.Point{X: 1, Y: 1}
}

func GetObjectCenter(object *Object) sdl.FPoint {
	size := GetObjectSize(object)
	return sdl.FPoint{
		X: float32(object.Tile.Position.Base.X) + float32(size.X)/2,
		Y: float32(object.Tile.Position.Base.Y) + float32(size.Y)/2,
	}
}

func InteractObject(mob *Mob, object *Object) {
	MobHitObject(mob, object)
}

func HitObject(object *Object) {
	PlaySound(&GameInstance.SoundController, "grillehit")
	particle := NewParticleCross(
		sdlutils.Vector3ToFVector3Center(object.Tile.Position),
		0.75,
		sdl.Color{R: 210, G: 210, B: 210},
	)
	CreateParticle(&GameInstance.ParticleController, particle)
	object.Health--
	if object.Health <= 0 {
		DestroyObject(object)
	}
}

func FindObjectsOfType(kmap *Map, z uint8, objectType string) []*Object {
	return filter.Choose(kmap.Objects, func(object *Object) bool {
		return object.Tile.Position.Z == z && object.Type == objectType
	}).([]*Object)
}

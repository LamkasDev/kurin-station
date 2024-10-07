package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
)

var ObjectContainer = map[string]*ObjectTemplate{}

func RegisterObjects() {
	ObjectContainer["airlock"] = NewObjectTemplateAirlock()
	ObjectContainer["big_thruster"] = NewObjectTemplateThruster("big_thruster", sdlutils.White)
	ObjectContainer["broken_grille"] = NewObjectTemplateBrokenGrille()
	ObjectContainer["console"] = NewObjectTemplateConsole()
	ObjectContainer["displaced"] = NewObjectTemplateDisplaced()
	ObjectContainer["grille"] = NewObjectTemplateGrille()
	ObjectContainer["lathe"] = NewObjectTemplateLathe()
	ObjectContainer["lattice_l"] = NewObjectTemplate[interface{}]("lattice_l", false)
	ObjectContainer["lattice_r"] = NewObjectTemplate[interface{}]("lattice_r", false)
	ObjectContainer["pod"] = NewObjectTemplate[interface{}]("pod", false)
	ObjectContainer["shuttle_wall"] = NewObjectTemplateWall("shuttle_wall")
	ObjectContainer["small_thruster_l"] = NewObjectTemplateThruster("small_thruster_l", sdlutils.Blue)
	ObjectContainer["small_thruster_r"] = NewObjectTemplateThruster("small_thruster_r", sdlutils.Blue)
	ObjectContainer["telepad"] = NewObjectTemplateTelepad()
	ObjectContainer["teleporter"] = NewObjectTemplateTeleporter()
	ObjectContainer["wall"] = NewObjectTemplateWall("wall")
	ObjectContainer["window"] = NewObjectTemplate[interface{}]("window", false)
}

func NewObject(tile *Tile, objectType string) *Object {
	object := &Object{
		Id:        GetNextId(),
		Type:      objectType,
		Tile:      tile,
		Direction: common.DirectionSouth,
		Template:  ObjectContainer[objectType],
	}
	object.Health = object.Template.MaxHealth
	object.Data = object.Template.GetDefaultData()

	return object
}

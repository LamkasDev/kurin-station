package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type ObjectTeleporterData struct {
	Target     sdlutils.Vector3
	Processing bool
	Charged    bool
	Ticks      uint16
}

func NewObjectTemplateTeleporter() *ObjectTemplate {
	template := NewObjectTemplate[*ObjectTeleporterData]("teleporter", false)
	template.Process = func(object *Object) {
		data := object.Data.(*ObjectTeleporterData)
		if data.Processing {
			data.Ticks++
			if data.Ticks >= 120 {
				mobs := GetMobsOnTile(&GameInstance.Map, object.Tile)
				if len(mobs) > 0 {
					for _, mob := range mobs {
						MoveMob(mob, data.Target)
					}
					PlaySound(&GameInstance.SoundController, "synth_yes")
				} else {
					PlaySound(&GameInstance.SoundController, "synth_no")
				}
				data.Charged = false
				data.Processing = false
				data.Ticks = 0
			}
		} else if !data.Charged {
			data.Ticks++
			if data.Ticks >= 360 {
				data.Charged = true
				data.Ticks = 0
			}
		} else if !data.Processing {
			if GameInstance.Ticks%20 != 0 {
				return
			}
			mobs := GetMobsOnTile(&GameInstance.Map, object.Tile)
			if len(mobs) == 0 {
				return
			}
			data.Processing = true
		}
	}
	template.OnInteraction = func(object *Object, item *Item) bool {
		return true
	}
	template.GetTexture = func(object *Object) int {
		data := object.Data.(*ObjectTeleporterData)
		if data.Processing {
			if data.Ticks%20 < 10 {
				return 1
			}

			return 2
		}
		if !data.Charged {
			if data.Ticks%120 < 60 {
				return 0
			}

			return 1
		}

		return 1
	}
	template.GetDefaultData = func() interface{} {
		return &ObjectTeleporterData{}
	}
	template.IsPassable = func(object *Object) bool {
		return true
	}
	template.MaxHealth = 100

	return template
}

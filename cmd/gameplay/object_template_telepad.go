package gameplay

import "robpike.io/filter"

type ObjectTelepadData struct {
	Processing bool
	Ticks      uint8
}

func NewObjectTemplateTelepad() *ObjectTemplate {
	template := NewObjectTemplate[*ObjectTelepadData]("telepad", false)
	template.Process = func(object *Object) {
		data := object.Data.(*ObjectTelepadData)
		if !data.Processing {
			return
		}
		data.Ticks++
		if data.Ticks >= 90 {
			items := GetItemsOnTile(&GameInstance.Map, object.Tile)
			items = filter.Choose(items, func(item *Item) bool {
				return item.Template.Price > 0
			}).([]*Item)
			for _, item := range items {
				if !RemoveItemFromMapRaw(&GameInstance.Map, item) {
					continue
				}
				AddCredits(item.Template.Price * uint32(item.Count))
			}
			if len(items) > 0 {
				PlaySound(&GameInstance.SoundController, "synth_yes")
			} else {
				PlaySound(&GameInstance.SoundController, "synth_no")
			}
			data.Processing = false
			data.Ticks = 0
		}
	}
	template.OnInteraction = func(object *Object, item *Item) bool {
		data := object.Data.(*ObjectTelepadData)
		if !data.Processing {
			data.Processing = true
		}

		return true
	}
	template.GetTexture = func(object *Object) int {
		data := object.Data.(*ObjectTelepadData)
		if data.Ticks%20 < 10 {
			return 0
		}

		return 1
	}
	template.GetDefaultData = func() interface{} {
		return &ObjectTelepadData{
			Processing: false,
		}
	}
	template.IsPassable = func(object *Object) bool {
		return true
	}
	template.MaxHealth = 100

	return template
}

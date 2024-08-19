package gameplay

func NewObjectTemplateWall(wallType string) *ObjectTemplate {
	template := NewObjectTemplate[interface{}](wallType, true)
	template.OnInteraction = func(object *Object, item *Item) bool {
		if item != nil {
			switch data := item.Data.(type) {
			case ItemWelderData:
				if data.Enabled && object.Health < object.Template.MaxHealth {
					PlaySoundVolume(&GameInstance.SoundController, "welder", 0.5)
					object.Health++
				}
				return true
			}
		}

		return false
	}
	template.OnDestroy = func(object *Object) {
		CreateObject(object.Tile, "displaced")
	}
	template.MaxHealth = 10

	return template
}

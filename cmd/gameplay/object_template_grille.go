package gameplay

func NewObjectTemplateGrille() *ObjectTemplate {
	template := NewObjectTemplate[interface{}]("grille", false)
	template.OnInteraction = func(object *Object, item *Item) bool {
		if item != nil {
			switch data := item.Data.(type) {
			case ItemWelderData:
				if data.Enabled && object.Health < 3 {
					PlaySoundVolume(&GameInstance.SoundController, "welder", 0.5)
					object.Health++
				}
				return true
			}
		}

		return false
	}
	template.OnDestroy = func(object *Object) {
		CreateObject(object.Tile, "broken_grille")
	}
	template.MaxHealth = 3

	return template
}

package gameplay

type ObjectPodData struct {
	Enabled bool
}

func NewObjectTemplatePod() *ObjectTemplate {
	template := NewObjectTemplate[*ObjectPodData]("pod", false)
	template.OnInteraction = func(object *Object, item *Item) bool {
		if item != nil && item.Type == "credit" {
			if !RemoveItemFromCharacterRaw(item, item.Character) {
				return false
			}
			PlaySound(&GameInstance.SoundController, "jingle")
			GameInstance.Credits++
			return true
		}

		OpenDialog(&DialogRequest{Type: "pod", Data: &DialogPodData{Pod: object}})
		return true
	}
	template.GetDefaultData = func() interface{} {
		return &ObjectPodData{
			Enabled: false,
		}
	}
	template.MaxHealth = 100

	return template
}

package gameplay

type ObjectConsoleData struct {
	Enabled bool
}

func NewObjectTemplateConsole() *ObjectTemplate {
	template := NewObjectTemplate[*ObjectConsoleData]("console", false)
	template.OnInteraction = func(object *Object, item *Item) bool {
		if item != nil && item.Type == "credit" {
			if !RemoveItemFromCharacterRaw(item, item.Mob) {
				return false
			}
			PlaySound(&GameInstance.SoundController, "jingle")
			AddCredits(1)
			return true
		}

		OpenDialog(&DialogRequest{Type: "console", Data: &DialogConsoleData{Console: object}})
		return true
	}
	template.GetDefaultData = func() interface{} {
		return &ObjectConsoleData{
			Enabled: false,
		}
	}
	template.MaxHealth = 100

	return template
}

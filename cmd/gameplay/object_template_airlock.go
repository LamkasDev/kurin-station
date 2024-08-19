package gameplay

import "github.com/arl/math32"

type ObjectAirlockData struct {
	Open           bool
	AnimationTicks uint8
}

func NewObjectTemplateAirlock() *ObjectTemplate {
	template := NewObjectTemplate[*ObjectAirlockData]("airlock", false)
	template.Process = func(object *Object) {
		data := object.Data.(*ObjectAirlockData)
		if data.Open {
			if data.AnimationTicks < 50 {
				data.AnimationTicks++
			}
		} else {
			if data.AnimationTicks > 0 {
				data.AnimationTicks--
			}
		}
	}
	template.IsPassable = func(object *Object) bool {
		data := object.Data.(*ObjectAirlockData)
		return data.Open && data.AnimationTicks >= 20
	}
	template.OnInteraction = func(object *Object, item *Item) bool {
		data := object.Data.(*ObjectAirlockData)
		data.Open = !data.Open
		if data.Open {
			PlaySoundVolume(&GameInstance.SoundController, "airlockopen", 0.5)
		} else {
			PlaySoundVolume(&GameInstance.SoundController, "airlockclose", 0.5)
		}

		return true
	}
	template.GetTexture = func(object *Object) int {
		data := object.Data.(*ObjectAirlockData)
		return int(math32.Floor(float32(data.AnimationTicks) / 10))
	}
	template.GetDefaultData = func() interface{} {
		return &ObjectAirlockData{}
	}
	template.MaxHealth = 10

	return template
}

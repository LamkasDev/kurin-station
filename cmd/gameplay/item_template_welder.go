package gameplay

import (
	"math"
)

type ItemWelderData struct {
	Enabled bool
}

func NewItemTemplateWelder() *ItemTemplate {
	template := NewItemTemplate[*ItemWelderData]("welder", 1, 3)
	template.GetTextures = func(item *Item) []int {
		if item.Data.(*ItemWelderData).Enabled {
			if int64(math.Floor(float64(GameInstance.Ticks)/8))%2 == 0 {
				return []int{0, 2}
			}
			return []int{0, 1}
		}

		return []int{0}
	}
	template.GetTextureHand = func(item *Item) int {
		if item.Data.(*ItemWelderData).Enabled {
			if int64(math.Floor(float64(GameInstance.Ticks)/8))%2 == 0 {
				return 1
			}
			return 0
		}

		return 0
	}
	template.OnHandInteraction = func(item *Item) {
		data := item.Data.(*ItemWelderData)
		data.Enabled = !data.Enabled
		if data.Enabled {
			PlaySoundVolume(&GameInstance.SoundController, "welderactivate", 0.5)
		} else {
			PlaySoundVolume(&GameInstance.SoundController, "welderdeactivate", 0.5)
		}
	}
	template.CanHit = false
	template.GetDefaultData = func() interface{} {
		return &ItemWelderData{}
	}

	return template
}

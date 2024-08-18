package gameplay

import (
	"math"
)

type KurinItemWelderData struct {
	Enabled bool
}

func NewKurinItemWelder() *KurinItem {
	item := NewKurinItemRaw[*KurinItemWelderData]("welder", 1)
	item.GetTextures = GetTexturesKurinItemWelder
	item.GetTextureHand = GetTextureHandKurinItemWelder
	item.OnHandInteraction = func(item *KurinItem) {
		data := item.Data.(*KurinItemWelderData)
		data.Enabled = !data.Enabled
		if data.Enabled {
			PlaySoundVolume(&GameInstance.SoundController, "welderactivate", 0.5)
		} else {
			PlaySoundVolume(&GameInstance.SoundController, "welderdeactivate", 0.5)
		}
	}
	item.CanHit = false
	item.Data = &KurinItemWelderData{
		Enabled: false,
	}

	return item
}

func GetTexturesKurinItemWelder(item *KurinItem) []int {
	if item.Data.(*KurinItemWelderData).Enabled {
		if int64(math.Floor(float64(GameInstance.Ticks)/8))%2 == 0 {
			return []int{0, 2}
		}
		return []int{0, 1}
	}

	return []int{0}
}

func GetTextureHandKurinItemWelder(item *KurinItem) int {
	if item.Data.(*KurinItemWelderData).Enabled {
		if int64(math.Floor(float64(GameInstance.Ticks)/8))%2 == 0 {
			return 1
		}
		return 0
	}

	return 0
}

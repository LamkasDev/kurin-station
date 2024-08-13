package gameplay

import (
	"math"

	"github.com/kelindar/binary"
)

type KurinItemWelderData struct {
	Enabled bool
}

func NewKurinItemWelder() *KurinItem {
	item := NewKurinItemRaw("welder")
	item.GetTextures = GetTexturesKurinItemWelder
	item.GetTextureHand = GetTextureHandKurinItemWelder
	item.OnHandInteraction = func(item *KurinItem) {
		data := item.Data.(KurinItemWelderData)
		data.Enabled = !data.Enabled
		if data.Enabled {
			PlaySoundVolume(&KurinGameInstance.SoundController, "welderactivate", 0.5)
		} else {
			PlaySoundVolume(&KurinGameInstance.SoundController, "welderdeactivate", 0.5)
		}
		item.Data = data
	}
	item.EncodeData = func(item *KurinItem) []byte {
		itemData := item.Data.(KurinItemWelderData)
		data, _ := binary.Marshal(&itemData)
		return data
	}
	item.DecodeData = func(item *KurinItem, data []byte) {
		var itemData KurinItemWelderData 
		binary.Unmarshal(data, &itemData)
		item.Data = itemData
	}
	item.CanHit = false
	item.Data = KurinItemWelderData{
		Enabled: false,
	}

	return item
}

func GetTexturesKurinItemWelder(item *KurinItem) []int {
	if(item.Data.(KurinItemWelderData).Enabled) {
		if int64(math.Floor(float64(KurinGameInstance.Ticks) / 8)) % 2 == 0 {
			return []int{0, 2}
		}
		return []int{0, 1}
	}
	
	return []int{0}
}

func GetTextureHandKurinItemWelder(item *KurinItem) int {
	if(item.Data.(KurinItemWelderData).Enabled) {
		if int64(math.Floor(float64(KurinGameInstance.Ticks) / 8)) % 2 == 0 {
			return 1
		}
		return 0
	}
	
	return 0
}

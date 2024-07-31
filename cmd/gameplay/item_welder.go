package gameplay

import (
	"math"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type KurinItemWelderData struct {
	Enabled bool
}

func NewKurinItemWelder(transform *sdlutils.Transform) *KurinItem {
	return &KurinItem{
		Type:      "welder",
		Transform: transform,
		GetTextures: GetTexturesKurinItemWelder,
		GetTextureHand: GetTextureHandKurinItemWelder,
		Interact: func(item *KurinItem, game *KurinGame) {
			data := item.Data.(KurinItemWelderData)
			data.Enabled = !data.Enabled
			item.Data = data
		},
		Process: func(item *KurinItem, game *KurinGame) {},
		Data: KurinItemWelderData{
			Enabled: false,
		},
	}
}

func GetTexturesKurinItemWelder(item *KurinItem, game *KurinGame) []int {
	if(item.Data.(KurinItemWelderData).Enabled) {
		if int64(math.Floor(float64(game.Ticks) / 8)) % 2 == 0 {
			return []int{0, 2}
		}
		return []int{0, 1}
	}
	
	return []int{0}
}

func GetTextureHandKurinItemWelder(item *KurinItem, game *KurinGame) int {
	if(item.Data.(KurinItemWelderData).Enabled) {
		if int64(math.Floor(float64(game.Ticks) / 8)) % 2 == 0 {
			return 1
		}
		return 0
	}
	
	return 0
}

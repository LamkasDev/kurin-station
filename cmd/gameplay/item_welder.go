package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

type KurinItemWelderData struct {
}

func NewKurinItemWelder(transform *sdlutils.Transform) *KurinItem {
	return &KurinItem{
		Type:      "welder",
		Transform: transform,
		GetTextures: GetTexturesKurinItemWelder,
		Process: func(item *KurinItem, game *KurinGame) {},
		Data: KurinItemWelderData{},
	}
}

func GetTexturesKurinItemWelder(item *KurinItem, game *KurinGame) []int {
	return []int{0, 1}
}
